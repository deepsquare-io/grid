package benchmark

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/secret"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateHPLPhase1JobDefinition(t *testing.T) {
	// Arrange
	tests := []struct {
		title    string
		opts     []Option
		expected *Benchmark
	}{
		{
			title: "1 node, 2 gpus per node",
			opts: []Option{
				WithImage("registry-1.deepsquare.run#library/hpc-benchmarks:23.5"),
				WithClusterSpecs(1, 16, 2, 128460),
			},
			expected: &Benchmark{
				NTasks:        2,
				NTasksPerNode: 2,
				MinNodes:      1,
				MaxNodes:      1,
				GPUsPerNode:   2,
				CPUsPerNode:   16,
				CPUsPerTask:   8,
				Memory:        utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

#SBATCH -N 1-1
#SBATCH --ntasks=2
#SBATCH --ntasks-per-node=2
#SBATCH --mem=0
#SBATCH --mincpus=16
#SBATCH --gpus-per-node=2
#SBATCH --cpus-per-task=8

set -ex

GPU="$(srun --ntasks=1 -N 1-1 --gpus-per-task=1 nvidia-smi --query-gpu=name --format=csv,noheader | head -1)"
export GPU

srun --ntasks=1 -N 1-1 --container-image="registry-1.docker.io#library/python:slim" sh -c '
set -ex
pip3 install --no-cache-dir archspec

export DEBIAN_FRONTEND=noninteractive

apt update -y -qq && apt install -y -qq golang curl

curl -fsSL -k \
  -d "{\"microarch\":\"$(archspec cpu)\",\"os\":\"$(go env GOOS)\",\"arch\":\"$(go env GOARCH)\", \"cpu\":\"$(grep "model name" /proc/cpuinfo | awk -F: '"'"'{print $2}'"'"' | sed '"'"'s/^[ \t]*//'"'"' | head -1)\",\"gpu\":\"${GPU}\"}" \
  -X POST \
  -H "X-Secret: %s" \
  -H "Content-Type: application/json" \
  "https://localhost:3000/benchmark/machine"
'
export NCCL_P2P_DISABLE=1
# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory or TCP as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp
export NCCL_IB_DISABLE=1

GPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
GPU_AFFINITY="${GPU_AFFINITY%%|}"
export GPU_AFFINITY

CPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $7}'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
CPU_AFFINITY="${CPU_AFFINITY%%|}"
export CPU_AFFINITY

set +e

srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --nodes=1-1 \
  --ntasks=2 \
  --ntasks-per-node=2 \
  --gpus-per-task=1 \
  --container-image="registry-1.deepsquare.run#library/hpc-benchmarks:23.5" \
  bash -c 'sed -Ei "s/:1//g" ./hpl.sh
sed -Ei "s/'"'"':'"'"'/'"'"'|'"'"'/g" ./hpl.sh
cat << '"'"'EOF'"'"' > /tmp/test.dat \
  && ./hpl.sh \
  --xhpl-ai \
  --cpu-affinity "$CPU_AFFINITY" \
  --gpu-affinity "$GPU_AFFINITY" \
  --dat "/tmp/test.dat"
HPLinpack benchmark input file
Innovative Computing Laboratory, University of Tennessee
HPL.out      output file name (if any)
6            device out (6=stdout,7=stderr,file)
10 # of problems sizes (N)
95000 96000 97000 98000 100000 101000 102000 103000 105000 106000   Ns
10   # of NBs
64 128 224 256 384 512 640 768 896 1024    NBs
0            PMAP process mapping (0=Row-,1=Column-major)
1            # of process grids (P x Q)
2            Ps
1            Qs
16.0         threshold
1            # of panel fact
2            PFACTs (0=left, 1=Crout, 2=Right)
1            # of recursive stopping criterium
4            NBMINs (>= 1)
1            # of panels in recursion
2            NDIVs
1            # of recursive panel fact.
1            RFACTs (0=left, 1=Crout, 2=Right)
1            # of broadcast
1            BCASTs (0=1rg,1=1rM,2=2rg,3=2rM,4=Lng,5=LnM)
1            # of lookahead depth
1            DEPTHs (>=0)
2            SWAP (0=bin-exch,1=long,2=mix)
64           swapping threshold
1            L1 in (0=transposed,1=no-transposed) form
0            U  in (0=transposed,1=no-transposed) form
1            Equilibration (0=no,1=yes)
8            memory alignment in double (> 0)
EOF'

LOG_FILE="$(scontrol show job $SLURM_JOB_ID | grep "StdOut=" | sed 's/.*StdOut=//g')"

curl -fsSL -k \
  --upload-file \
  "$LOG_FILE" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/hpl/phase1?nodes=1&cpusPerNode=16&gpusPerNode=2&memPerNode=128460"
`,
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
				),
			},
		},
		{
			title: "1 node, 4 gpus per node",
			opts: []Option{
				WithImage("registry-1.deepsquare.run#library/hpc-benchmarks:23.5"),
				WithClusterSpecs(1, 16, 4, 128460),
			},
			expected: &Benchmark{
				NTasks:        4,
				NTasksPerNode: 4,
				MinNodes:      1,
				MaxNodes:      1,
				GPUsPerNode:   4,
				CPUsPerNode:   16,
				CPUsPerTask:   4,
				Memory:        utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

#SBATCH -N 1-1
#SBATCH --ntasks=4
#SBATCH --ntasks-per-node=4
#SBATCH --mem=0
#SBATCH --mincpus=16
#SBATCH --gpus-per-node=4
#SBATCH --cpus-per-task=4

set -ex

GPU="$(srun --ntasks=1 -N 1-1 --gpus-per-task=1 nvidia-smi --query-gpu=name --format=csv,noheader | head -1)"
export GPU

srun --ntasks=1 -N 1-1 --container-image="registry-1.docker.io#library/python:slim" sh -c '
set -ex
pip3 install --no-cache-dir archspec

export DEBIAN_FRONTEND=noninteractive

apt update -y -qq && apt install -y -qq golang curl

curl -fsSL -k \
  -d "{\"microarch\":\"$(archspec cpu)\",\"os\":\"$(go env GOOS)\",\"arch\":\"$(go env GOARCH)\", \"cpu\":\"$(grep "model name" /proc/cpuinfo | awk -F: '"'"'{print $2}'"'"' | sed '"'"'s/^[ \t]*//'"'"' | head -1)\",\"gpu\":\"${GPU}\"}" \
  -X POST \
  -H "X-Secret: %s" \
  -H "Content-Type: application/json" \
  "https://localhost:3000/benchmark/machine"
'
export NCCL_P2P_DISABLE=1
# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory or TCP as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp
export NCCL_IB_DISABLE=1

GPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=4 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
GPU_AFFINITY="${GPU_AFFINITY%%|}"
export GPU_AFFINITY

CPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=4 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $7}'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
CPU_AFFINITY="${CPU_AFFINITY%%|}"
export CPU_AFFINITY

set +e

srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --nodes=1-1 \
  --ntasks=4 \
  --ntasks-per-node=4 \
  --gpus-per-task=1 \
  --container-image="registry-1.deepsquare.run#library/hpc-benchmarks:23.5" \
  bash -c 'sed -Ei "s/:1//g" ./hpl.sh
sed -Ei "s/'"'"':'"'"'/'"'"'|'"'"'/g" ./hpl.sh
cat << '"'"'EOF'"'"' > /tmp/test.dat \
  && ./hpl.sh \
  --xhpl-ai \
  --cpu-affinity "$CPU_AFFINITY" \
  --gpu-affinity "$GPU_AFFINITY" \
  --dat "/tmp/test.dat"
HPLinpack benchmark input file
Innovative Computing Laboratory, University of Tennessee
HPL.out      output file name (if any)
6            device out (6=stdout,7=stderr,file)
10 # of problems sizes (N)
95000 96000 97000 98000 100000 101000 102000 103000 105000 106000   Ns
10   # of NBs
64 128 224 256 384 512 640 768 896 1024    NBs
0            PMAP process mapping (0=Row-,1=Column-major)
1            # of process grids (P x Q)
2            Ps
2            Qs
16.0         threshold
1            # of panel fact
2            PFACTs (0=left, 1=Crout, 2=Right)
1            # of recursive stopping criterium
4            NBMINs (>= 1)
1            # of panels in recursion
2            NDIVs
1            # of recursive panel fact.
1            RFACTs (0=left, 1=Crout, 2=Right)
1            # of broadcast
1            BCASTs (0=1rg,1=1rM,2=2rg,3=2rM,4=Lng,5=LnM)
1            # of lookahead depth
1            DEPTHs (>=0)
2            SWAP (0=bin-exch,1=long,2=mix)
64           swapping threshold
1            L1 in (0=transposed,1=no-transposed) form
0            U  in (0=transposed,1=no-transposed) form
1            Equilibration (0=no,1=yes)
8            memory alignment in double (> 0)
EOF'

LOG_FILE="$(scontrol show job $SLURM_JOB_ID | grep "StdOut=" | sed 's/.*StdOut=//g')"

curl -fsSL -k \
  --upload-file \
  "$LOG_FILE" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/hpl/phase1?nodes=1&cpusPerNode=16&gpusPerNode=4&memPerNode=128460"
`,
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
				),
			},
		},
		{
			title: "2 nodes, 2 gpus per node",
			opts: []Option{
				WithImage("registry-1.deepsquare.run#library/hpc-benchmarks:23.5"),
				WithClusterSpecs(2, 16, 2, 128460),
			},
			expected: &Benchmark{
				NTasks:        4,
				NTasksPerNode: 2,
				MinNodes:      1,
				MaxNodes:      2,
				GPUsPerNode:   2,
				CPUsPerNode:   16,
				CPUsPerTask:   8,
				Memory:        utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

#SBATCH -N 1-2
#SBATCH --ntasks=4
#SBATCH --ntasks-per-node=2
#SBATCH --mem=0
#SBATCH --mincpus=16
#SBATCH --gpus-per-node=2
#SBATCH --cpus-per-task=8

set -ex

GPU="$(srun --ntasks=1 -N 1-1 --gpus-per-task=1 nvidia-smi --query-gpu=name --format=csv,noheader | head -1)"
export GPU

srun --ntasks=1 -N 1-1 --container-image="registry-1.docker.io#library/python:slim" sh -c '
set -ex
pip3 install --no-cache-dir archspec

export DEBIAN_FRONTEND=noninteractive

apt update -y -qq && apt install -y -qq golang curl

curl -fsSL -k \
  -d "{\"microarch\":\"$(archspec cpu)\",\"os\":\"$(go env GOOS)\",\"arch\":\"$(go env GOARCH)\", \"cpu\":\"$(grep "model name" /proc/cpuinfo | awk -F: '"'"'{print $2}'"'"' | sed '"'"'s/^[ \t]*//'"'"' | head -1)\",\"gpu\":\"${GPU}\"}" \
  -X POST \
  -H "X-Secret: %s" \
  -H "Content-Type: application/json" \
  "https://localhost:3000/benchmark/machine"
'
# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory or TCP as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp
export NCCL_IB_DISABLE=1

GPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
GPU_AFFINITY="${GPU_AFFINITY%%|}"
export GPU_AFFINITY

CPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $7}'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
CPU_AFFINITY="${CPU_AFFINITY%%|}"
export CPU_AFFINITY

set +e

srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --nodes=2-2 \
  --ntasks=4 \
  --ntasks-per-node=2 \
  --gpus-per-task=1 \
  --container-image="registry-1.deepsquare.run#library/hpc-benchmarks:23.5" \
  bash -c 'sed -Ei "s/:1//g" ./hpl.sh
sed -Ei "s/'"'"':'"'"'/'"'"'|'"'"'/g" ./hpl.sh
cat << '"'"'EOF'"'"' > /tmp/test.dat \
  && ./hpl.sh \
  --xhpl-ai \
  --cpu-affinity "$CPU_AFFINITY" \
  --gpu-affinity "$GPU_AFFINITY" \
  --dat "/tmp/test.dat"
HPLinpack benchmark input file
Innovative Computing Laboratory, University of Tennessee
HPL.out      output file name (if any)
6            device out (6=stdout,7=stderr,file)
10 # of problems sizes (N)
134000 136000 137000 139000 141000 143000 145000 146000 148000 150000   Ns
10   # of NBs
64 128 224 256 384 512 640 768 896 1024    NBs
0            PMAP process mapping (0=Row-,1=Column-major)
1            # of process grids (P x Q)
2            Ps
2            Qs
16.0         threshold
1            # of panel fact
2            PFACTs (0=left, 1=Crout, 2=Right)
1            # of recursive stopping criterium
4            NBMINs (>= 1)
1            # of panels in recursion
2            NDIVs
1            # of recursive panel fact.
1            RFACTs (0=left, 1=Crout, 2=Right)
1            # of broadcast
1            BCASTs (0=1rg,1=1rM,2=2rg,3=2rM,4=Lng,5=LnM)
1            # of lookahead depth
1            DEPTHs (>=0)
2            SWAP (0=bin-exch,1=long,2=mix)
64           swapping threshold
1            L1 in (0=transposed,1=no-transposed) form
0            U  in (0=transposed,1=no-transposed) form
1            Equilibration (0=no,1=yes)
8            memory alignment in double (> 0)
EOF'

LOG_FILE="$(scontrol show job $SLURM_JOB_ID | grep "StdOut=" | sed 's/.*StdOut=//g')"

curl -fsSL -k \
  --upload-file \
  "$LOG_FILE" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/hpl/phase1?nodes=2&cpusPerNode=16&gpusPerNode=2&memPerNode=128460"
`,
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
				),
			},
		},
		{
			title: "2 nodes, 2 gpus per node with UCX",
			opts: []Option{
				WithImage("registry-1.deepsquare.run#library/hpc-benchmarks:23.5"),
				WithClusterSpecs(2, 16, 2, 128460),
				WithUCX("mlx5_2:1|mlx5_2:1", "rc"),
			},
			expected: &Benchmark{
				NTasks:        4,
				NTasksPerNode: 2,
				MinNodes:      1,
				MaxNodes:      2,
				GPUsPerNode:   2,
				CPUsPerNode:   16,
				CPUsPerTask:   8,
				Memory:        utils.Ptr(uint64(0)),
				Body: fmt.Sprintf(`#!/bin/bash

#SBATCH -N 1-2
#SBATCH --ntasks=4
#SBATCH --ntasks-per-node=2
#SBATCH --mem=0
#SBATCH --mincpus=16
#SBATCH --gpus-per-node=2
#SBATCH --cpus-per-task=8

set -ex

GPU="$(srun --ntasks=1 -N 1-1 --gpus-per-task=1 nvidia-smi --query-gpu=name --format=csv,noheader | head -1)"
export GPU

srun --ntasks=1 -N 1-1 --container-image="registry-1.docker.io#library/python:slim" sh -c '
set -ex
pip3 install --no-cache-dir archspec

export DEBIAN_FRONTEND=noninteractive

apt update -y -qq && apt install -y -qq golang curl

curl -fsSL -k \
  -d "{\"microarch\":\"$(archspec cpu)\",\"os\":\"$(go env GOOS)\",\"arch\":\"$(go env GOARCH)\", \"cpu\":\"$(grep "model name" /proc/cpuinfo | awk -F: '"'"'{print $2}'"'"' | sed '"'"'s/^[ \t]*//'"'"' | head -1)\",\"gpu\":\"${GPU}\"}" \
  -X POST \
  -H "X-Secret: %s" \
  -H "Content-Type: application/json" \
  "https://localhost:3000/benchmark/machine"
'
# Select UCX as MPI P2P communications
export PMIX_MCA_pml=ucx
# Select UCX as Byte-Transport Layer
export PMIX_MCA_btl=^vader,openib,tcp
export OMPI_MCA_pml=ucx
export OMPI_MCA_btl=^vader,openib,tcp

GPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
GPU_AFFINITY="${GPU_AFFINITY%%|}"
export GPU_AFFINITY

CPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $7}'"'"' | tr '"'"'\n'"'"' '"'"'|'"'"'')"
done
CPU_AFFINITY="${CPU_AFFINITY%%|}"
export CPU_AFFINITY

set +e

srun \
  --cpu-bind=none \
  --gpu-bind=none \
  --nodes=2-2 \
  --ntasks=4 \
  --ntasks-per-node=2 \
  --gpus-per-task=1 \
  --container-image="registry-1.deepsquare.run#library/hpc-benchmarks:23.5" \
  bash -c 'sed -Ei "s/:1//g" ./hpl.sh
sed -Ei "s/'"'"':'"'"'/'"'"'|'"'"'/g" ./hpl.sh
readarray -t UCX_AFFINITY_MAP <<<"$(tr '"'"'|'"'"' '"'"'\n'"'"'<<<"mlx5_2:1|mlx5_2:1")"
UCX_NET_DEVICES="${UCX_AFFINITY_MAP[${SLURM_NODEID}]}"
export UCX_NET_DEVICES
export UCX_TLS='"'"'rc'"'"'
cat << '"'"'EOF'"'"' > /tmp/test.dat \
  && ./hpl.sh \
  --xhpl-ai \
  --cpu-affinity "$CPU_AFFINITY" \
  --gpu-affinity "$GPU_AFFINITY" \
  --dat "/tmp/test.dat"
HPLinpack benchmark input file
Innovative Computing Laboratory, University of Tennessee
HPL.out      output file name (if any)
6            device out (6=stdout,7=stderr,file)
10 # of problems sizes (N)
134000 136000 137000 139000 141000 143000 145000 146000 148000 150000   Ns
10   # of NBs
64 128 224 256 384 512 640 768 896 1024    NBs
0            PMAP process mapping (0=Row-,1=Column-major)
1            # of process grids (P x Q)
2            Ps
2            Qs
16.0         threshold
1            # of panel fact
2            PFACTs (0=left, 1=Crout, 2=Right)
1            # of recursive stopping criterium
4            NBMINs (>= 1)
1            # of panels in recursion
2            NDIVs
1            # of recursive panel fact.
1            RFACTs (0=left, 1=Crout, 2=Right)
1            # of broadcast
1            BCASTs (0=1rg,1=1rM,2=2rg,3=2rM,4=Lng,5=LnM)
1            # of lookahead depth
1            DEPTHs (>=0)
2            SWAP (0=bin-exch,1=long,2=mix)
64           swapping threshold
1            L1 in (0=transposed,1=no-transposed) form
0            U  in (0=transposed,1=no-transposed) form
1            Equilibration (0=no,1=yes)
8            memory alignment in double (> 0)
EOF'

LOG_FILE="$(scontrol show job $SLURM_JOB_ID | grep "StdOut=" | sed 's/.*StdOut=//g')"

curl -fsSL -k \
  --upload-file \
  "$LOG_FILE" \
  -H "X-Secret: %s" \
  "https://localhost:3000/benchmark/hpl/phase1?nodes=2&cpusPerNode=16&gpusPerNode=2&memPerNode=128460"
`,
					base64.StdEncoding.EncodeToString(secret.Get()),
					base64.StdEncoding.EncodeToString(secret.Get()),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			jobDefinition, err := GeneratePhase1HPLBenchmark(tt.opts...)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Body, jobDefinition.Body)
			assert.Equal(t, tt.expected, jobDefinition)
		})
	}
}

func TestCalculateProcessGrid(t *testing.T) {
	// Arrange
	tests := []struct {
		gpusPerNode uint64
		nodes       uint64
		expectedP   uint64
		expectedQ   uint64
	}{
		{
			gpusPerNode: 2,
			nodes:       1,
			expectedP:   2,
			expectedQ:   1,
		},
		{
			nodes:       3,
			gpusPerNode: 5,
			expectedP:   5,
			expectedQ:   3,
		},
		{
			nodes:       1,
			gpusPerNode: 15,
			expectedP:   5,
			expectedQ:   3,
		},
	}

	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("%d = %d * %d", tt.gpusPerNode, tt.expectedP, tt.expectedQ),
			func(t *testing.T) {
				P, Q, err := calculateProcessGrid(tt.gpusPerNode, tt.nodes)

				// Assert
				require.NoError(t, err)
				require.Equal(t, tt.expectedP, P)
				require.Equal(t, tt.expectedQ, Q)
			},
		)
	}
}

func TestCalculateProblemSize(t *testing.T) {
	// Arrange
	expectedProblemSize := "95000 96000 97000 98000 100000 101000 102000 103000 105000 106000 "
	memPerNode := uint64(128460)

	// Act
	_, problemSizeStr := calculateProblemSize(memPerNode, 1)

	// Assert
	require.Equal(t, expectedProblemSize, problemSizeStr)
}
