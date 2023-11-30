// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package benchmark

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
	"text/template"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils"
	"go.uber.org/zap"
)

const DefaultHPLImage = "registry-1.deepsquare.run#library/hpc-benchmarks:23.5"

// Problem sizes depend on the RAM. The greater it is, the more precise it is.
//
// At 80%, it is possible that Linux will swap the memory. A drop of performance will be seen.
var benchmarkMemoryUsePercentage = []float64{
	0.50, // 50% is a sure hit. We do this for healthchecking the HPL health.
	0.60,
	0.75,
	0.80, // This the most unstable test, which may end in OOM.
}

// Block sizes depend on the performance of the processing unit.
//
// It must be tested randomly. Though, it is said that an A100 takes a 1024 in input.
var benchmarkBlockSizes = []int{
	64, 128, 224, 256, 384, 512, 640, 768, 896, 1024,
}

func intToString(a []int, sep string) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, sep)
}

func applyHPLOptions(opts []Option) *options {
	o := &options{
		image:                   DefaultHPLImage,
		nodes:                   1,
		secret:                  base64.StdEncoding.EncodeToString(secret.Get()),
		supervisorPublicAddress: "localhost:3000",
		additionalEnv:           make(map[string]string),
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
	opts ...Option,
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
		NBlockSize:   uint64(len(benchmarkBlockSizes)),
		BlockSize:    intToString(benchmarkBlockSizes, " "),
	}

	return prepareHPLJobDefinition(params, o)
}

func prepareHPLJobDefinition(
	params *hplParams,
	o *options,
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
		template.New("benchmark").
			Funcs(funcMap()).
			ParseFS(templates, "templates/benchmark-hpl.tmpl", "templates/dat.tmpl"),
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
		Env                     map[string]string
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
		Env:                     o.additionalEnv,
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
