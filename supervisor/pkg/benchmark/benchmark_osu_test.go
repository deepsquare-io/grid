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

func TestGenerateOSUBenchmark(t *testing.T) {
	// Arrange
	tests := []struct {
		title    string
		opts     []benchmark.BenchmarkOption
		expected *benchmark.Benchmark
	}{
		{
			title: "3 nodes, 2 gpus per node",
			opts: []benchmark.BenchmarkOption{
				benchmark.WithImage("registry-1.deepsquare.run#library/osu-benchmarks:latest"),
				benchmark.WithSupervisorPublicAddress("localhost:3000"),
				benchmark.WithClusterSpecs(3, 16, 2, 100000),
			},
			expected: &benchmark.Benchmark{
				NTasks:        3,
				NTasksPerNode: 1,
				MinNodes:      1,
				MaxNodes:      3,
				CPUsPerTask:   1,
				GPUsPerNode:   1,
				Memory:        utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

#SBATCH -N 1-3
#SBATCH --ntasks=3
#SBATCH --ntasks-per-node=1
#SBATCH --gpus-per-node=1
#SBATCH --cpus-per-task=1
#SBATCH --mem=0

set -x
# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory or TCP as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp
export NCCL_IB_DISABLE=1

# P2P Bidirectional Latency
file1="$(mktemp -t benchmark.XXXXXX)"
dir="$(dirname "$file1")"
srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --gpus-per-task=1 \
  --container-mounts="$dir:$dir:rw" \
  --nodes=2-2 \
  --ntasks=2 \
  --container-image="registry-1.deepsquare.run#library/osu-benchmarks:latest" \
  /osu-micro-benchmarks/mpi/pt2pt/osu_latency | tee "$file1"

# P2P Bidirectional Bandwidth
file2="$(mktemp -t benchmark.XXXXXX)"
dir="$(dirname "$file2")"
srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --gpus-per-task=1 \
  --container-mounts="$dir:$dir:rw" \
  --nodes=2-2 \
  --ntasks=2 \
  --container-image="registry-1.deepsquare.run#library/osu-benchmarks:latest" \
  /osu-micro-benchmarks/mpi/pt2pt/osu_bibw | tee "$file2"

# All to all
file3="$(mktemp -t benchmark.XXXXXX)"
dir="$(dirname "$file3")"
srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --gpus-per-task=1 \
  --container-mounts="$dir:$dir:rw" \
  --nodes=3-3 \
  --ntasks=3 \
  --container-image="registry-1.deepsquare.run#library/osu-benchmarks:latest" \
  /osu-micro-benchmarks/mpi/collective/osu_alltoall | tee "$file3"

curl -fsSL \
  --upload-file \
  "$file1" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/osu/pt2pt-latency"

curl -fsSL \
  --upload-file \
  "$file2" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/osu/pt2pt-bibw"

curl -fsSL \
  --upload-file \
  "$file3" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/osu/alltoall"
`,
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get())),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			b, err := benchmark.GenerateOSUBenchmark(tt.opts...)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Body, b.Body)
			assert.Equal(t, tt.expected, b)
		})
	}
}
