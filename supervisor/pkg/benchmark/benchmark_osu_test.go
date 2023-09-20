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
		opts     []benchmark.Option
		expected *benchmark.Benchmark
	}{
		{
			title: "3 nodes, 2 gpus per node",
			opts: []benchmark.Option{
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

set -ex
# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory or TCP as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp
export NCCL_IB_DISABLE=1

export SHARED_WORLD_TMP_VOLUME=/opt/cache/world-tmp
umask 077

RESULT_DIR="${SHARED_WORLD_TMP_VOLUME}$(mktemp --directory -t benchmark.XXXXXX -u)"
export RESULT_DIR

srun -N 3-3 \
  --ntasks=3 \
  --ntasks-per-node=1 \
  mkdir -p "$RESULT_DIR"

# P2P Bidirectional Latency
srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --gpus-per-task=1 \
  --container-mounts="$RESULT_DIR:$RESULT_DIR:rw" \
  --nodes=2-2 \
  --ntasks=2 \
  --container-image="registry-1.deepsquare.run#library/osu-benchmarks:latest" \
  bash -c '
/osu-micro-benchmarks/mpi/pt2pt/osu_latency | tee "$RESULT_DIR/pt2pt-latency"'

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/pt2pt-latency" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/osu/pt2pt-latency"

# P2P Bidirectional Bandwidth
srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --gpus-per-task=1 \
  --container-mounts="$RESULT_DIR:$RESULT_DIR:rw" \
  --nodes=2-2 \
  --ntasks=2 \
  --container-image="registry-1.deepsquare.run#library/osu-benchmarks:latest" \
  bash -c '
/osu-micro-benchmarks/mpi/pt2pt/osu_bibw | tee "$RESULT_DIR/pt2pt-bibw"'

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/pt2pt-bibw" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/osu/pt2pt-bibw"

# All to all
srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --gpus-per-task=1 \
  --container-mounts="$RESULT_DIR:$RESULT_DIR:rw" \
  --nodes=3-3 \
  --ntasks=3 \
  --container-image="registry-1.deepsquare.run#library/osu-benchmarks:latest" \
  bash -c '
/osu-micro-benchmarks/mpi/collective/osu_alltoall | tee "$RESULT_DIR/alltoall"'

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/alltoall" \
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
