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
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils"
	"go.uber.org/zap"
)

const DefaultSpeedTestImage = "registry-1.docker.io#gists/speedtest-cli:1.2.0"

func applySpeedTestOptions(opts []Option) *options {
	o := &options{
		image:  DefaultSpeedTestImage,
		secret: base64.StdEncoding.EncodeToString(secret.Get()),
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func GenerateSpeedTestBenchmark(
	opts ...Option,
) (*Benchmark, error) {
	o := applySpeedTestOptions(opts)
	benchmark := &Benchmark{
		NTasks:      1,
		CPUsPerTask: 1,
		Memory:      utils.Ptr(uint64(0)),
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
