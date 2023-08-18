package benchmark

import (
	"encoding/base64"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"go.uber.org/zap"
)

const DefaultOSUImage = "registry-1.deepsquare.run#library/osu-benchmarks:latest"

func applyOSUOptions(opts []BenchmarkOption) *benchmarkOptions {
	o := &benchmarkOptions{
		image:                   DefaultOSUImage,
		nodes:                   1,
		secret:                  base64.StdEncoding.EncodeToString(secret.Get()),
		supervisorPublicAddress: "localhost:3000",
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func GenerateOSUBenchmark(
	opts ...BenchmarkOption,
) (*Benchmark, error) {
	o := applyOSUOptions(opts)
	benchmark := &Benchmark{
		MinNodes:      1,
		MaxNodes:      o.nodes,
		NTasks:        o.nodes,
		NTasksPerNode: 1,
		GPUsPerNode:   1,
		CPUsPerTask:   1,
		Memory:        utils.Ptr(uint64(0)),
	}
	sbatchTmpl := template.Must(
		template.New("benchmark").
			Funcs(sprig.TxtFuncMap()).
			ParseFS(templates, "templates/benchmark-osu.tmpl"),
	)
	sbatchBuilder := new(strings.Builder)
	if err := sbatchTmpl.ExecuteTemplate(sbatchBuilder, "benchmark", struct {
		Benchmark
		Image                   string
		SupervisorPublicAddress string
		Secret                  string
		UCX                     bool
		UCXAffinity             string
		UCXTransport            string
		Trace                   bool
	}{
		Benchmark:               *benchmark,
		Image:                   o.image,
		SupervisorPublicAddress: o.supervisorPublicAddress,
		Secret:                  o.secret,
		UCX:                     o.ucx,
		UCXAffinity:             o.ucxAffinity,
		UCXTransport:            o.ucxTransport,
		Trace:                   o.trace,
	}); err != nil {
		logger.I.Error("sbatch templating failed", zap.Error(err))
		return nil, err
	}
	benchmark.Body = sbatchBuilder.String()

	return benchmark, nil
}
