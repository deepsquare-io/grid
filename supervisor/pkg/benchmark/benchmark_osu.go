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

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils"
	"go.uber.org/zap"
)

const DefaultOSUImage = "registry-1.deepsquare.run#library/osu-benchmarks:latest"

func applyOSUOptions(opts []Option) *options {
	o := &options{
		image:                   DefaultOSUImage,
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

func GenerateOSUBenchmark(
	opts ...Option,
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
			Funcs(funcMap()).
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
		Env                     map[string]string
	}{
		Benchmark:               *benchmark,
		Image:                   o.image,
		SupervisorPublicAddress: o.supervisorPublicAddress,
		Secret:                  o.secret,
		UCX:                     o.ucx,
		UCXAffinity:             o.ucxAffinity,
		UCXTransport:            o.ucxTransport,
		Trace:                   o.trace,
		Env:                     o.additionalEnv,
	}); err != nil {
		logger.I.Error("sbatch templating failed", zap.Error(err))
		return nil, err
	}
	benchmark.Body = sbatchBuilder.String()

	return benchmark, nil
}
