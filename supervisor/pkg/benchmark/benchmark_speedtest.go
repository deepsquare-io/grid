package benchmark

import (
	"encoding/base64"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"go.uber.org/zap"
)

const DefaultSpeedTestImage = "registry-1.docker.io#gists/speedtest-cli:1.2.0"

func applySpeedTestOptions(opts []BenchmarkOption) *benchmarkOptions {
	o := &benchmarkOptions{
		image:  DefaultSpeedTestImage,
		secret: base64.StdEncoding.EncodeToString(secret.Get()),
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func GenerateSpeedTestBenchmark(
	opts ...BenchmarkOption,
) (*Benchmark, error) {
	o := applySpeedTestOptions(opts)
	benchmark := &Benchmark{
		NTasks: 1,
	}
	sbatchTmpl := template.Must(
		template.New("benchmark").
			Funcs(sprig.TxtFuncMap()).
			ParseFS(templates, "templates/benchmark-speedtest.tmpl"),
	)
	sbatchBuilder := new(strings.Builder)
	if err := sbatchTmpl.ExecuteTemplate(sbatchBuilder, "benchmark", struct {
		Image                   string
		SupervisorPublicAddress string
		Secret                  string
	}{
		Image:                   o.image,
		SupervisorPublicAddress: o.supervisorPublicAddress,
		Secret:                  o.secret,
	}); err != nil {
		logger.I.Error("sbatch templating failed", zap.Error(err))
		return nil, err
	}
	benchmark.Body = sbatchBuilder.String()

	return benchmark, nil
}
