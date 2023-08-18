package benchmark_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSpeedTestBenchmark(t *testing.T) {
	// Arrange
	tests := []struct {
		title    string
		opts     []benchmark.BenchmarkOption
		expected *benchmark.Benchmark
	}{
		{
			title: "3 nodes, 2 gpus per node",
			opts: []benchmark.BenchmarkOption{
				benchmark.WithImage("registry-1.docker.io#gists/speedtest-cli:1.2.0"),
				benchmark.WithSupervisorPublicAddress("localhost:3000"),
			},
			expected: &benchmark.Benchmark{
				NTasks:      1,
				CPUsPerTask: 1,
				Memory:      utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

set -e

file="$(mktemp -t benchmark.XXXXXX)"
dir="$(dirname "$file")"

srun --container-mounts="$dir:$dir:rw" \
  --container-image="registry-1.docker.io#gists/speedtest-cli:1.2.0" \
  /usr/local/bin/speedtest --accept-license --accept-gdpr -f json-pretty > "$file"

curl -fsSL \
  --upload-file \
  "$file" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/speedtest"
`, base64.StdEncoding.EncodeToString(secret.Get())),
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
