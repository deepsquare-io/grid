// Copyright (C) 2023 DeepSquare Asociation
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

package benchmark_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSpeedTestBenchmark(t *testing.T) {
	// Arrange
	tests := []struct {
		title    string
		opts     []benchmark.Option
		expected *benchmark.Benchmark
	}{
		{
			title: "3 nodes, 2 gpus per node",
			opts: []benchmark.Option{
				benchmark.WithImage("registry-1.docker.io#gists/speedtest-cli:1.2.0"),
				benchmark.WithSupervisorPublicAddress("localhost:3000"),
			},
			expected: &benchmark.Benchmark{
				NTasks:      1,
				CPUsPerTask: 1,
				Memory:      utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

set -ex

file="$(mktemp -t benchmark.XXXXXX)"
dir="$(dirname "$file")"

srun --container-mounts="$dir:$dir:rw" \
  --container-image="registry-1.docker.io#gists/speedtest-cli:1.2.0" \
  /usr/local/bin/speedtest --accept-license --accept-gdpr -f json-pretty > "$file"

curl -fsSL -k \
  --upload-file \
  "$file" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/speedtest"
`,
					base64.StdEncoding.EncodeToString(secret.Get()),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			b, err := benchmark.GenerateSpeedTestBenchmark(tt.opts...)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Body, b.Body)
			assert.Equal(t, tt.expected, b)
		})
	}
}
