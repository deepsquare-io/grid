package benchmark

import (
	"context"
	"embed"
	"encoding/base64"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils/hash"
	"go.uber.org/zap"
)

const (
	GBtoMB        = 1000
	jobNameFormat = "benchmark-%s"
)

//go:embed templates/*.tmpl
var templates embed.FS

var benchmarkMemoryUsePercentage = []float64{
	0.75,
	0.76,
	0.77,
	0.78,
	0.79,
	0.80,
	0.81,
	0.82,
	0.83,
	0.84,
}

type Launcher interface {
	RunPhase1(ctx context.Context) error
	RunPhase2(
		ctx context.Context,
		newP uint64,
		newQ uint64,
		newProblemSize uint64,
		newBlockSize uint64,
	) error
	// Verify is used to check if a job came from the same launcher.
	Verify(data []byte) bool
	// Cancel cancels all running benchmark
	Cancel(ctx context.Context) error
	// Get the generated job name for benchmarks.
	GetJobName() string
}

type launcher struct {
	secretManager           secret.Manager
	Image                   string
	supervisorPublicAddress string
	user                    string
	Scheduler               scheduler.Scheduler
	nodes                   uint64
	cpusPerNode             []uint64
	memPerNode              []uint64
	gpusPerNode             []uint64
	timeLimit               time.Duration
	ucx                     bool
	ucxAffinity             string
	ucxTransport            string
	wait                    bool
	// trace enables benchmark trace logging.
	trace bool
}

type BenchmarkParams struct {
	NProblemSize uint64
	ProblemSize  string
	NBlockSize   uint64
	BlockSize    string
	P            uint64
	Q            uint64
}

type LauncherOption func(*launcher)

func WithSecretManager(secretManager secret.Manager) LauncherOption {
	return func(l *launcher) {
		l.secretManager = secretManager
	}
}

func WithUCX(affinity string, transport string) LauncherOption {
	return func(l *launcher) {
		l.ucx = true
		l.ucxTransport = transport
		l.ucxAffinity = affinity
	}
}

func WithNoWait() LauncherOption {
	return func(l *launcher) {
		l.wait = false
	}
}

func WithTrace() LauncherOption {
	return func(l *launcher) {
		l.trace = true
	}
}

func NewLauncher(
	image string,
	user string,
	supervisorPublicAddress string,
	scheduler scheduler.Scheduler,
	nodes uint64,
	cpusPerNode []uint64,
	memPerNode []uint64,
	gpusPerNode []uint64,
	timeLimit time.Duration,
	opts ...LauncherOption,
) Launcher {
	l := &launcher{
		Image:                   image,
		Scheduler:               scheduler,
		user:                    user,
		supervisorPublicAddress: supervisorPublicAddress,
		nodes:                   nodes,
		cpusPerNode:             cpusPerNode,
		gpusPerNode:             gpusPerNode,
		memPerNode:              memPerNode,
		timeLimit:               timeLimit,
		secretManager:           secret.NewManager(),
		wait:                    true,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *launcher) GetJobName() string {
	hash := hash.GenerateAlphanumeric(l.supervisorPublicAddress)
	return fmt.Sprintf(jobNameFormat, hash)
}

func (l *launcher) runBenchmark(
	ctx context.Context,
	params *BenchmarkParams,
	phase string,
) error {
	log := logger.I.With(zap.String("phase", phase))
	jobDefinition, err := l.createJobDefinition(
		ctx,
		params.P,
		params.Q,
		l.nodes,
		slices.Min(l.gpusPerNode),
		slices.Min(l.cpusPerNode),
	)
	if err != nil {
		log.Error("failed to create job definition", zap.Error(err))
		return err
	}

	// Compute CPUsPerTask to be passed as a srun params (and not as a jobDefinition)
	jobDefinition.CPUsPerTask = jobDefinition.CPUsPerNode / jobDefinition.NTasksPerNode

	body, err := l.generateBody(ctx, params, jobDefinition, phase)
	if err != nil {
		return err
	}
	jobDefinition.Body = body

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	errC := make(chan error, 1)
	go func() {
		_, err := l.Scheduler.Submit(ctx, &scheduler.SubmitRequest{
			Name:          l.GetJobName(),
			User:          l.user,
			Prefix:        l.GetJobName(),
			JobDefinition: jobDefinition,
		})
		errC <- err
	}()

	log.Info("benchmark started")

	for {
		select {
		case err := <-errC:
			if err != nil {
				log.Error("benchmark failed", zap.Error(err))
			} else {
				log.Info("benchmark succeeded")
			}
			return err
		case <-ticker.C:
			log.Info("benchmark is still running")
		}
	}
}

func (l *launcher) RunPhase1(ctx context.Context) error {
	p, q, err := CalculateProcessGrid(slices.Min(l.gpusPerNode), l.nodes)
	if err != nil {
		logger.I.Error("failed to compute p and q", zap.Error(err))
		return err
	}
	nProblemSize, problemSize := CalculateProblemSize(slices.Min(l.memPerNode), l.nodes)
	params := &BenchmarkParams{
		P:            p,
		Q:            q,
		NProblemSize: nProblemSize,
		ProblemSize:  problemSize,
		NBlockSize:   10,
		BlockSize:    "64 128 224 256 384 512 640 768 896 1024",
	}

	return l.runBenchmark(ctx, params, "phase1")
}

func (l *launcher) RunPhase2(
	ctx context.Context,
	newP uint64,
	newQ uint64,
	newProblemSize uint64,
	newBlockSize uint64,
) error {
	params := &BenchmarkParams{
		P:            newP,
		Q:            newQ,
		NProblemSize: 1,
		ProblemSize:  strconv.FormatUint(newProblemSize, 10),
		NBlockSize:   1,
		BlockSize:    strconv.FormatUint(newBlockSize, 10),
	}

	return l.runBenchmark(ctx, params, "phase2")
}

func (l *launcher) Verify(data []byte) bool {
	return l.secretManager.Validate(data)
}

func (l *launcher) generateBody(
	ctx context.Context,
	params *BenchmarkParams,
	jobDefinition *scheduler.JobDefinition,
	phase string,
) (sbatchBody string, err error) {
	sbatchTmpl := template.Must(
		template.New("benchmark").Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap{
			"div": func(a uint64, b uint64) uint64 {
				return a / b
			},
			"mul": func(a uint64, b uint64) uint64 {
				return a * b
			},
		}).ParseFS(templates, "templates/benchmark.tmpl", "templates/dat.tmpl"),
	)
	sbatchBuilder := new(strings.Builder)
	if err := sbatchTmpl.ExecuteTemplate(sbatchBuilder, "benchmark", struct {
		Image                   string
		BenchmarkParams         BenchmarkParams
		JobDefinition           scheduler.JobDefinition
		SupervisorPublicAddress string
		Phase                   string
		Secret                  string
		UCX                     bool
		UCXAffinity             string
		UCXTransport            string
		Trace                   bool
	}{
		Image:                   l.Image,
		BenchmarkParams:         *params,
		JobDefinition:           *jobDefinition,
		SupervisorPublicAddress: l.supervisorPublicAddress,
		Phase:                   phase,
		Secret:                  base64.StdEncoding.EncodeToString(l.secretManager.Get()),
		UCX:                     l.ucx,
		UCXAffinity:             l.ucxAffinity,
		UCXTransport:            l.ucxTransport,
		Trace:                   l.trace,
	}); err != nil {
		logger.I.Error("sbatch templating failed", zap.Error(err))
		return "", err
	}

	return sbatchBuilder.String(), nil
}

func (l *launcher) createJobDefinition(
	ctx context.Context,
	p uint64,
	q uint64,
	nodes uint64,
	gpusPerNode uint64,
	cpusPerNode uint64,
) (*scheduler.JobDefinition, error) {
	jobDefinition := &scheduler.JobDefinition{}
	jobDefinition.MinNodes = 1
	jobDefinition.MaxNodes = nodes
	ntasks := p * q
	jobDefinition.NTasks = ntasks
	ntasksPerNode := p * q / nodes
	jobDefinition.NTasksPerNode = ntasksPerNode
	jobDefinition.CPUsPerNode = cpusPerNode
	jobDefinition.GPUsPerNode = gpusPerNode
	jobDefinition.Memory = utils.Ptr(uint64(0))
	jobDefinition.TimeLimit = uint64(l.timeLimit.Minutes())
	jobDefinition.Wait = l.wait

	return jobDefinition, nil
}

func (l *launcher) Cancel(ctx context.Context) error {
	return l.Scheduler.CancelJob(ctx, l.GetJobName(), l.user)
}

// CalculateProcessGrid computes the optimal values of P and Q based on the number of GPUs available per nodes
func CalculateProcessGrid(
	gpusPerNode uint64,
	nodes uint64,
) (P uint64, Q uint64, err error) {
	totalGPUs := gpusPerNode * nodes
	if totalGPUs == 1 {
		return 1, 1, nil
	}
	sqrtTotalGPUS := uint64(math.Sqrt(float64(totalGPUs)))

	for i := sqrtTotalGPUS; i > 0; i-- {
		if totalGPUs%i == 0 {
			return totalGPUs / i, i, nil
		}
	}
	return totalGPUs, 1, nil // If no other valid P is found, default to 2
}

// CalculateProblemSize computes the problem size from the ram available
func CalculateProblemSize(
	memPerNode uint64,
	nodes uint64,
) (nProblemSize uint64, problemSize string) {
	nProblemSize = uint64(len(benchmarkMemoryUsePercentage))
	for _, values := range benchmarkMemoryUsePercentage {
		problemSizeInt := int(
			math.Sqrt(float64(memPerNode*nodes)/8)*values,
		) * GBtoMB

		problemSize += strconv.Itoa(problemSizeInt) + " "
	}

	return nProblemSize, problemSize
}
