package benchmark

import (
	"context"
	"embed"
	"encoding/base64"
	"errors"
	"math"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils/array"
	"go.uber.org/zap"
)

const (
	User        = "root"
	GBtoMB      = 1000
	JobName     = "HPL-Benchmark"
	DatFilePath = "hpl.dat"
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
	RunPhase1(ctx context.Context, nodes uint64) error
	RunPhase2(
		ctx context.Context,
		newP uint64,
		newQ uint64,
		newProblemSize uint64,
		newBlockSize uint64,
		nodes uint64,
	) error
	// Verify is used to check if a job came from the same launcher.
	Verify(data []byte) bool
}

type launcher struct {
	secretManager           secret.Manager
	Image                   string
	supervisorPublicAddress string
	Scheduler               scheduler.Scheduler
}

type BenchmarkParams struct {
	NProblemSize uint64
	ProblemSize  string
	NBlockSize   uint64
	BlockSize    string
	P            uint64
	Q            uint64
	CPUsPerTask  uint64
}

type LauncherOption func(*launcher)

func WithSecretManager(secretManager secret.Manager) LauncherOption {
	return func(l *launcher) {
		l.secretManager = secretManager
	}
}

func NewLauncher(
	image string,
	supervisorPublicAddress string,
	scheduler scheduler.Scheduler,
	opts ...LauncherOption,
) Launcher {
	l := &launcher{
		Image:                   image,
		Scheduler:               scheduler,
		supervisorPublicAddress: supervisorPublicAddress,
		secretManager:           secret.NewManager(),
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *launcher) runBenchmark(
	ctx context.Context,
	params *BenchmarkParams,
	phase string,
	nodes uint64,
	gpusPerNode uint64,
) error {
	cpusPerNodes, err := l.Scheduler.FindCPUsPerNode(ctx)
	if err != nil {
		logger.I.Error("failed to find cpus per node", zap.Error(err))
		return err
	}
	if len(cpusPerNodes) == 0 {
		return errors.New("no cpus per node")
	}
	jobDefinition, err := l.createJobDefinition(
		ctx,
		params.P,
		params.Q,
		nodes,
		gpusPerNode,
		array.Min(cpusPerNodes),
	)
	if err != nil {
		logger.I.Error("failed to create job definition", zap.Error(err))
		return err
	}

	// Compute CPUsPerTask to be passed as a srun params (and not as a jobDefinition)
	cpusPerTasks := jobDefinition.CPUsPerNode / jobDefinition.NTasksPerNode
	params.CPUsPerTask = cpusPerTasks

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
			Name:          JobName,
			User:          User,
			Prefix:        "benchmark",
			JobDefinition: jobDefinition,
		})
		errC <- err
	}()

	logger.I.Info("benchmark started")

	for {
		select {
		case err := <-errC:
			if err != nil {
				logger.I.Error("benchmark failed", zap.Error(err))
			} else {
				logger.I.Info("benchmark succeeded")
			}
			return err
		case <-ticker.C:
			logger.I.Info("benchmark is still running")
		}
	}
}

func (l *launcher) RunPhase1(ctx context.Context, nodes uint64) error {
	memPerNodes, err := l.Scheduler.FindMemPerNode(ctx)
	if err != nil {
		logger.I.Error("failed to find mem per node", zap.Error(err))
		return err
	}
	if len(memPerNodes) == 0 {
		return errors.New("mem per node is empty")
	}
	memPerNode := array.Min(memPerNodes)
	gpusPerNodes, err := l.Scheduler.FindGPUsPerNode(ctx)
	if err != nil {
		logger.I.Error("failed to find gpus per node", zap.Error(err))
		return err
	}
	if len(gpusPerNodes) == 0 {
		return errors.New("gpus per node is empty")
	}
	gpuPerNode := array.Min(gpusPerNodes)
	p, q, err := CalculateProcessGrid(gpuPerNode, nodes)
	if err != nil {
		logger.I.Error("failed to compute p and q", zap.Error(err))
		return err
	}
	nProblemSize, problemSize := CalculateProblemSize(memPerNode, nodes)
	params := &BenchmarkParams{
		P:            p,
		Q:            q,
		NProblemSize: nProblemSize,
		ProblemSize:  problemSize,
		NBlockSize:   10,
		BlockSize:    "64 128 224 256 384 512 640 768 896 1024",
	}

	return l.runBenchmark(ctx, params, "phase1", nodes, gpuPerNode)
}

func (l *launcher) RunPhase2(
	ctx context.Context,
	newP uint64,
	newQ uint64,
	newProblemSize uint64,
	newBlockSize uint64,
	nodes uint64,
) error {
	gpusPerNodes, err := l.Scheduler.FindGPUsPerNode(ctx)
	if err != nil {
		logger.I.Error("failed to find gpus per node", zap.Error(err))
		return err
	}
	if len(gpusPerNodes) == 0 {
		return errors.New("gpus per node is empty")
	}
	gpuPerNode := array.Min(gpusPerNodes)
	params := &BenchmarkParams{
		P:            newP,
		Q:            newQ,
		NProblemSize: 1,
		ProblemSize:  strconv.FormatUint(newProblemSize, 10),
		NBlockSize:   1,
		BlockSize:    strconv.FormatUint(newBlockSize, 10),
	}

	return l.runBenchmark(ctx, params, "phase2", nodes, gpuPerNode)
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
	}{
		Image:                   l.Image,
		BenchmarkParams:         *params,
		JobDefinition:           *jobDefinition,
		SupervisorPublicAddress: l.supervisorPublicAddress,
		Phase:                   phase,
		Secret:                  base64.StdEncoding.EncodeToString(l.secretManager.Get()),
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
	jobDefinition.TimeLimit = 60
	jobDefinition.Wait = true

	return jobDefinition, nil
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
		if totalGPUs%i == 0 && i != 1 {
			return i, totalGPUs / i, nil
		}
	}
	return 2, totalGPUs, nil // If no other valid P is found, default to 2
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
