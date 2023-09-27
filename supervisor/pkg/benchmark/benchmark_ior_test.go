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

func TestGenerateIORBenchmark(t *testing.T) {
	// Arrange
	tests := []struct {
		title    string
		opts     []benchmark.Option
		expected *benchmark.Benchmark
	}{
		{
			title: "3 nodes",
			opts: []benchmark.Option{
				benchmark.WithImage("registry-1.deepsquare.run#library/ior-benchmarks:latest"),
				benchmark.WithSupervisorPublicAddress("localhost:3000"),
				benchmark.WithClusterSpecs(3, 16, 2, 100000),
			},
			expected: &benchmark.Benchmark{
				NTasks:        48,
				NTasksPerNode: 16,
				MinNodes:      1,
				MaxNodes:      3,
				CPUsPerTask:   1,
				CPUsPerNode:   16,
				Memory:        utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

#SBATCH -N 1-3
#SBATCH --ntasks=48
#SBATCH --ntasks-per-node=16
#SBATCH --mincpus=16
#SBATCH --cpus-per-task=1
#SBATCH --mem=0

set -ex
# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory or TCP as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp

export SCRATCH_VOLUME=/opt/cache/shared
export SHARED_TMP_VOLUME=/opt/cache/persistent
export SHARED_WORLD_TMP_VOLUME=/opt/cache/world-tmp
export DISK_TMP_VOLUME=/opt/cache/disk/tmp
export DISK_WORLD_TMP_VOLUME=/opt/cache/disk/world-tmp

# Block size is the size of all IO operation for 1 task. 10 Go per node.
export BLOCK_SIZE="$(( 10737418240 / 16))"
# Transfer size is the size a single IO operation.
export TRANSFER_SIZE=2M

umask 077

RESULT_DIR="${SHARED_WORLD_TMP_VOLUME}$(mktemp --directory -t benchmark.XXXXXX -u)"
export RESULT_DIR

srun -N 3-3 \
  --ntasks=3 \
  --ntasks-per-node=1 \
  mkdir -p "$RESULT_DIR"

# -w writeFile
# -r readFile
# -o testFile (outputFile)
# -b blockSize
# -a API for IO ([POSIX|AIO|DUMMY|MPIIO|MMAP])
# -i Iterations
# -F file-per-process
# -z random io
# -t transferSize
# -C reorderTasks (changes task ordering for readback (useful to avoid client cache))
# -e perform fsync upon POSIX write close, make sure reads are only started are all writes are done.
# -g use barriers between open, write/read, and close
srun \
  -N 1-1 \
  --ntasks="16" \
  --container-image="registry-1.deepsquare.run#library/ior-benchmarks:latest" \
  --container-mounts="$DISK_TMP_VOLUME:$DISK_TMP_VOLUME:rw,$RESULT_DIR:$RESULT_DIR:rw" \
  bash -c '
/usr/local/bin/ior \
  -w \
  -r \
  -o "$DISK_TMP_VOLUME/test.ior" \
  -b "$BLOCK_SIZE" \
  -a POSIX \
  -i 5 \
  -F \
  -t "$TRANSFER_SIZE" \
  -C \
  -e \
  -g \
  -v \
  -O summaryFormat=CSV \
  -O summaryFile="$RESULT_DIR/disk_tmp.csv"'

cat "$RESULT_DIR/disk_tmp.csv"

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/disk_tmp.csv" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/ior/disk-tmp"

srun \
  -N 1-1 \
  --ntasks="16" \
  --container-image="registry-1.deepsquare.run#library/ior-benchmarks:latest" \
  --container-mounts="$DISK_WORLD_TMP_VOLUME:$DISK_WORLD_TMP_VOLUME:rw,$RESULT_DIR:$RESULT_DIR:rw" \
  bash -c '
/usr/local/bin/ior \
  -w \
  -r \
  -o "$DISK_WORLD_TMP_VOLUME/test.ior" \
  -b "$BLOCK_SIZE" \
  -a POSIX \
  -i 5 \
  -F \
  -t "$TRANSFER_SIZE" \
  -C \
  -e \
  -g \
  -v \
  -O summaryFormat=CSV \
  -O summaryFile="$RESULT_DIR/disk_world_tmp.csv"'

cat "$RESULT_DIR/disk_world_tmp.csv"

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/disk_world_tmp.csv" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/ior/disk-world-tmp"

srun \
  -N 3-3 \
  --distribution cyclic \
  --ntasks="48" \
  --container-image="registry-1.deepsquare.run#library/ior-benchmarks:latest" \
  --container-mounts="$SCRATCH_VOLUME:$SCRATCH_VOLUME:rw,$RESULT_DIR:$RESULT_DIR:rw" \
  bash -c '
/usr/local/bin/ior \
  -w \
  -r \
  -o "$SCRATCH_VOLUME/test.ior" \
  -b "$BLOCK_SIZE" \
  -a MPIIO \
  -i 5 \
  -t "$TRANSFER_SIZE" \
  -C \
  --collective \
  -g \
  -v \
  -O summaryFormat=CSV \
  -O summaryFile="$RESULT_DIR/scratch.csv"'

cat "$RESULT_DIR/scratch.csv"

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/scratch.csv" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/ior/scratch"

srun \
  -N 3-3 \
  --distribution cyclic \
  --ntasks="48" \
  --container-image="registry-1.deepsquare.run#library/ior-benchmarks:latest" \
  --container-mounts="$SHARED_TMP_VOLUME:$SHARED_TMP_VOLUME:rw,$RESULT_DIR:$RESULT_DIR:rw" \
  bash -c '
/usr/local/bin/ior \
  -w \
  -r \
  -o "$SHARED_TMP_VOLUME/test.ior" \
  -b "$BLOCK_SIZE" \
  -a MPIIO \
  -i 5 \
  -t "$TRANSFER_SIZE" \
  -C \
  --collective \
  -g \
  -v \
  -O summaryFormat=CSV \
  -O summaryFile="$RESULT_DIR/shared_tmp.csv"'

cat "$RESULT_DIR/shared_tmp.csv"

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/shared_tmp.csv" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/ior/shared-tmp"

srun \
  -N 3-3 \
  --distribution cyclic \
  --ntasks="48" \
  --container-image="registry-1.deepsquare.run#library/ior-benchmarks:latest" \
  --container-mounts="$SHARED_WORLD_TMP_VOLUME:$SHARED_WORLD_TMP_VOLUME:rw,$RESULT_DIR:$RESULT_DIR:rw" \
  bash -c '
/usr/local/bin/ior \
  -w \
  -r \
  -o "$SHARED_WORLD_TMP_VOLUME/test.ior" \
  -b "$BLOCK_SIZE" \
  -a MPIIO \
  -i 5 \
  -t "$TRANSFER_SIZE" \
  -C \
  --collective \
  -g \
  -v \
  -O summaryFormat=CSV \
  -O summaryFile="$RESULT_DIR/shared_world_tmp.csv"'

cat "$RESULT_DIR/shared_world_tmp.csv"

curl -fsSL -k \
  --upload-file \
  "$RESULT_DIR/shared_world_tmp.csv" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/ior/shared-world-tmp"`,
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			b, err := benchmark.GenerateIORBenchmark(tt.opts...)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Body, b.Body)
			assert.Equal(t, tt.expected, b)
		})
	}
}
