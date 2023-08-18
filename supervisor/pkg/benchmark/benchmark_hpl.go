package benchmark

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"go.uber.org/zap"
)

const DefaultHPLImage = "registry-1.deepsquare.run#library/hpc-benchmarks:23.5"

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

func applyHPLOptions(opts []BenchmarkOption) *benchmarkOptions {
	o := &benchmarkOptions{
		image:                   DefaultHPLImage,
		nodes:                   1,
		secret:                  base64.StdEncoding.EncodeToString(secret.Get()),
		supervisorPublicAddress: "localhost:3000",
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type hplParams struct {
	NProblemSize uint64
	ProblemSize  string
	NBlockSize   uint64
	BlockSize    string
	P            uint64
	Q            uint64
}

func GeneratePhase1HPLBenchmark(
	opts ...BenchmarkOption,
) (*Benchmark, error) {
	o := applyHPLOptions(opts)
	o.phase = "phase1"
	p, q, err := calculateProcessGrid(o.gpusPerNode, o.nodes)
	if err != nil {
		return nil, fmt.Errorf("failed to compute p and q: %w", err)
	}
	nProblemSize, problemSize := calculateProblemSize(o.memPerNode, o.nodes)
	params := &hplParams{
		P:            p,
		Q:            q,
		NProblemSize: nProblemSize,
		ProblemSize:  problemSize,
		NBlockSize:   10,
		BlockSize:    "64 128 224 256 384 512 640 768 896 1024",
	}

	return prepareHPLJobDefinition(params, o)
}

func GeneratePhase2HPLBenchmark(
	newP uint64,
	newQ uint64,
	newProblemSize uint64,
	newBlockSize uint64,
	opts ...BenchmarkOption,
) (*Benchmark, error) {
	o := applyHPLOptions(opts)
	o.phase = "phase2"
	params := &hplParams{
		P:            newP,
		Q:            newQ,
		NProblemSize: 1,
		ProblemSize:  strconv.FormatUint(newProblemSize, 10),
		NBlockSize:   1,
		BlockSize:    strconv.FormatUint(newBlockSize, 10),
	}

	return prepareHPLJobDefinition(params, o)
}

func prepareHPLJobDefinition(
	params *hplParams,
	o *benchmarkOptions,
) (*Benchmark, error) {
	benchmark := &Benchmark{
		MinNodes:      1,
		MaxNodes:      o.nodes,
		NTasks:        params.P * params.Q,
		NTasksPerNode: (params.P * params.Q) / o.nodes,
		CPUsPerNode:   o.cpusPerNode,
		GPUsPerNode:   o.gpusPerNode,
		CPUsPerTask:   o.cpusPerNode / ((params.P * params.Q) / o.nodes),
		Memory:        utils.Ptr(uint64(0)),
	}

	sbatchTmpl := template.Must(
		template.New("benchmark").Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap{
			"div": func(a uint64, b uint64) uint64 {
				return a / b
			},
			"mul": func(a uint64, b uint64) uint64 {
				return a * b
			},
		}).ParseFS(templates, "templates/benchmark-hpl.tmpl", "templates/dat.tmpl"),
	)
	sbatchBuilder := new(strings.Builder)
	if err := sbatchTmpl.ExecuteTemplate(sbatchBuilder, "benchmark", struct {
		Image                   string
		BenchmarkParams         hplParams
		Benchmark               Benchmark
		SupervisorPublicAddress string
		Phase                   string
		Secret                  string
		UCX                     bool
		UCXAffinity             string
		UCXTransport            string
		Trace                   bool
		MemPerNode              uint64
	}{
		Image:                   o.image,
		BenchmarkParams:         *params,
		Benchmark:               *benchmark,
		SupervisorPublicAddress: o.supervisorPublicAddress,
		Phase:                   o.phase,
		Secret:                  o.secret,
		UCX:                     o.ucx,
		UCXAffinity:             o.ucxAffinity,
		UCXTransport:            o.ucxTransport,
		Trace:                   o.trace,
		MemPerNode:              o.memPerNode,
	}); err != nil {
		logger.I.Error("sbatch templating failed", zap.Error(err))
		return nil, err
	}

	benchmark.Body = sbatchBuilder.String()

	return benchmark, nil
}

// calculateProcessGrid computes the optimal values of P and Q based on the number of GPUs available per nodes
func calculateProcessGrid(
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

// calculateProblemSize computes the problem size from the ram available
func calculateProblemSize(
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
