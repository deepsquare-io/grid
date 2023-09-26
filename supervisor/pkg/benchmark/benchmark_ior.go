package benchmark

import (
	"encoding/base64"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils"
	"go.uber.org/zap"
)

const DefaultIORImage = "registry-1.deepsquare.run#library/ior-benchmarks:latest"

func applyIOROptions(opts []Option) *options {
	o := &options{
		image:                   DefaultIORImage,
		nodes:                   1,
		secret:                  base64.StdEncoding.EncodeToString(secret.Get()),
		supervisorPublicAddress: "localhost:3000",
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func GenerateIORBenchmark(
	opts ...Option,
) (*Benchmark, error) {
	o := applyIOROptions(opts)
	benchmark := &Benchmark{
		MinNodes:      1,
		MaxNodes:      o.nodes,
		NTasks:        o.cpusPerNode * o.nodes,
		NTasksPerNode: o.cpusPerNode,
		CPUsPerNode:   o.cpusPerNode,
		CPUsPerTask:   1,
		Memory:        utils.Ptr(uint64(0)),
	}
	sbatchTmpl := template.Must(
		template.New("benchmark").
			Funcs(sprig.TxtFuncMap()).
			ParseFS(templates, "templates/benchmark-ior.tmpl"),
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
