package benchmark_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	JobName = "HPL-Benchmark"
	admin   = "root"
)

type BenchmarkLauncherTestSuite struct {
	suite.Suite
	scheduler     *mocks.Scheduler
	secretManager *mocks.SecretManager
	impl          benchmark.Launcher
}

func (suite *BenchmarkLauncherTestSuite) SetupSubTest() {
	suite.scheduler = mocks.NewScheduler(suite.T())
	suite.secretManager = mocks.NewSecretManager(suite.T())

	suite.impl = benchmark.NewLauncher(
		"/etc/hpl-benchmark/hpc-benchmarks:hpl.sqsh",
		"supervisor.example.com:3000",
		suite.scheduler,
		benchmark.WithSecretManager(suite.secretManager),
	)
}

func (suite *BenchmarkLauncherTestSuite) TestRunPhase1() {
	// Arrange
	tests := []struct {
		title                 string
		gpusPerNode           uint64
		cpusPerNode           uint64
		memPerNode            uint64
		nodes                 uint64
		expectedSubmitRequest scheduler.SubmitRequest
	}{
		{
			title:       "1 node, 2 gpus per node",
			cpusPerNode: 16,
			gpusPerNode: 2,
			memPerNode:  128460,
			nodes:       1,
			expectedSubmitRequest: scheduler.SubmitRequest{
				Name:   JobName,
				User:   admin,
				Prefix: "benchmark",
				JobDefinition: &scheduler.JobDefinition{
					NTasks:        4,
					NTasksPerNode: 4,
					MinNodes:      1,
					MaxNodes:      1,
					GPUsPerNode:   2,
					CPUsPerNode:   16,
					Memory:        utils.Ptr(uint64(0)),
					Wait:          true,
					TimeLimit:     60,
					Body: `#!/bin/bash

#SBATCH -N 1-1
#SBATCH --ntasks=4
#SBATCH --ntasks-per-node=4
#SBATCH --mem=0
#SBATCH --mincpus=16
#SBATCH --gpus-per-node=2
#SBATCH --cpus-per-task=4

# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp

GPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
  # Add affinity for GPU sharing
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
done
GPU_AFFINITY="${GPU_AFFINITY%:}"
export GPU_AFFINITY

CPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $7}'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
  # Add affinity for GPU sharing
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
done
CPU_AFFINITY="${CPU_AFFINITY%:}"
export CPU_AFFINITY

mkdir -p /tmp/benchmark-result

srun --mpi=pmix_v4 \
  --cpu-bind=none \
  --gpu-bind=none \
  --nodes=1-1 \
  --ntasks=4 \
  --ntasks-per-node=4 \
  --container-mounts="/tmp/benchmark-result:/out:rw" \
  --container-image="/etc/hpl-benchmark/hpc-benchmarks:hpl.sqsh" \
  sh -c 'cat << '"'"'EOF'"'"' > /tmp/test.dat \
  && sed -Ei "s/:1//g" ./hpl.sh \
  && ./hpl.sh --xhpl-ai --cpu-affinity $CPU_AFFINITY --cpu-cores-per-rank 4 --gpu-affinity $GPU_AFFINITY --dat "/tmp/test.dat"
HPLinpack benchmark input file
Innovative Computing Laboratory, University of Tennessee
/out/HPL.out      output file name (if any)
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

curl -sS \
  --upload-file \
  "$LOG_FILE" \
  -H "X-Secret: U0VDUkVU" \
  "https://supervisor.example.com:3000/benchmark/phase1?nodes=1"
`,
				},
			},
		},
		{
			title:       "1 node, 4 gpus per node",
			cpusPerNode: 16,
			gpusPerNode: 4,
			memPerNode:  128460,
			nodes:       1,
			expectedSubmitRequest: scheduler.SubmitRequest{
				Name:   JobName,
				User:   admin,
				Prefix: "benchmark",
				JobDefinition: &scheduler.JobDefinition{
					NTasks:        4,
					NTasksPerNode: 4,
					MinNodes:      1,
					MaxNodes:      1,
					GPUsPerNode:   4,
					CPUsPerNode:   16,
					TimeLimit:     60,
					Wait:          true,
					Memory:        utils.Ptr(uint64(0)),
					Body: `#!/bin/bash

#SBATCH -N 1-1
#SBATCH --ntasks=4
#SBATCH --ntasks-per-node=4
#SBATCH --mem=0
#SBATCH --mincpus=16
#SBATCH --gpus-per-node=4
#SBATCH --cpus-per-task=4

# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp

GPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=4 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
done
GPU_AFFINITY="${GPU_AFFINITY%:}"
export GPU_AFFINITY

CPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=4 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $7}'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
done
CPU_AFFINITY="${CPU_AFFINITY%:}"
export CPU_AFFINITY

mkdir -p /tmp/benchmark-result

srun --mpi=pmix_v4 \
  --cpu-bind=none \
  --gpu-bind=none \
  --nodes=1-1 \
  --ntasks=4 \
  --ntasks-per-node=4 \
  --gpus-per-task=1 \
  --container-mounts="/tmp/benchmark-result:/out:rw" \
  --container-image="/etc/hpl-benchmark/hpc-benchmarks:hpl.sqsh" \
  sh -c 'cat << '"'"'EOF'"'"' > /tmp/test.dat \
  && sed -Ei "s/:1//g" ./hpl.sh \
  && ./hpl.sh --xhpl-ai --cpu-affinity $CPU_AFFINITY --cpu-cores-per-rank 4 --gpu-affinity $GPU_AFFINITY --dat "/tmp/test.dat"
HPLinpack benchmark input file
Innovative Computing Laboratory, University of Tennessee
/out/HPL.out      output file name (if any)
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

curl -sS \
  --upload-file \
  "$LOG_FILE" \
  -H "X-Secret: U0VDUkVU" \
  "https://supervisor.example.com:3000/benchmark/phase1?nodes=1"
`,
				},
			},
		},
		{
			title:       "2 nodes, 2 gpus per node",
			cpusPerNode: 16,
			gpusPerNode: 2,
			memPerNode:  128460,
			nodes:       2,
			expectedSubmitRequest: scheduler.SubmitRequest{
				Name:   JobName,
				User:   admin,
				Prefix: "benchmark",
				JobDefinition: &scheduler.JobDefinition{
					NTasks:        4,
					NTasksPerNode: 2,
					MinNodes:      1,
					MaxNodes:      2,
					GPUsPerNode:   2,
					CPUsPerNode:   16,
					TimeLimit:     60,
					Wait:          true,
					Memory:        utils.Ptr(uint64(0)),
					Body: `#!/bin/bash

#SBATCH -N 1-2
#SBATCH --ntasks=4
#SBATCH --ntasks-per-node=2
#SBATCH --mem=0
#SBATCH --mincpus=16
#SBATCH --gpus-per-node=2
#SBATCH --cpus-per-task=8

# Select obi-wan as MPI P2P communications
export PMIX_MCA_pml=ob1
# Select shared-memory as Byte-Transport Layer
export PMIX_MCA_btl=vader,self,tcp
export OMPI_MCA_pml=ob1
export OMPI_MCA_btl=vader,self,tcp

GPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  GPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $1}'"'"' | sed '"'"'s/GPU//'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
done
GPU_AFFINITY="${GPU_AFFINITY%:}"
export GPU_AFFINITY

CPU_AFFINITY=""
for node in $(scontrol show hostnames "$SLURM_NODELIST"); do
  CPU_AFFINITY+="$(srun --ntasks=1 -N 1-1 --gpus-per-task=2 --gpu-bind=none --cpu-bind=none -w "$node" sh -c 'nvidia-smi topo -m | grep -E '"'"'^GPU[0-9]+'"'"' | awk '"'"'{print $7}'"'"' | tr '"'"'\n'"'"' '"'"':'"'"'')"
done
CPU_AFFINITY="${CPU_AFFINITY%:}"
export CPU_AFFINITY

mkdir -p /tmp/benchmark-result

srun --mpi=pmix_v4 \
  --cpu-bind=none \
  --gpu-bind=none \
  --nodes=2-2 \
  --ntasks=4 \
  --ntasks-per-node=2 \
  --gpus-per-task=1 \
  --container-mounts="/tmp/benchmark-result:/out:rw" \
  --container-image="/etc/hpl-benchmark/hpc-benchmarks:hpl.sqsh" \
  sh -c 'cat << '"'"'EOF'"'"' > /tmp/test.dat \
  && sed -Ei "s/:1//g" ./hpl.sh \
  && ./hpl.sh --xhpl-ai --cpu-affinity $CPU_AFFINITY --cpu-cores-per-rank 8 --gpu-affinity $GPU_AFFINITY --dat "/tmp/test.dat"
HPLinpack benchmark input file
Innovative Computing Laboratory, University of Tennessee
/out/HPL.out      output file name (if any)
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

curl -sS \
  --upload-file \
  "$LOG_FILE" \
  -H "X-Secret: U0VDUkVU" \
  "https://supervisor.example.com:3000/benchmark/phase1?nodes=2"
`,
				},
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.title, func() {
			suite.secretManager.EXPECT().Get().Return([]byte("SECRET"))
			suite.scheduler.EXPECT().
				FindCPUsPerNode(mock.Anything).
				Return([]uint64{tt.cpusPerNode}, nil)
			suite.scheduler.EXPECT().
				FindGPUsPerNode(mock.Anything).
				Return([]uint64{tt.gpusPerNode}, nil)
			suite.scheduler.EXPECT().
				FindMemPerNode(mock.Anything).
				Return([]uint64{tt.memPerNode}, nil)
			suite.scheduler.EXPECT().
				Submit(mock.Anything, &tt.expectedSubmitRequest).
				Return("success", nil)

			// Act
			err := suite.impl.RunPhase1(context.Background(), tt.nodes)

			// Assert
			suite.NoError(err)
		})
	}
}

func (suite *BenchmarkLauncherTestSuite) TestCalculateProcessGrid() {
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
			expectedQ:   2,
		},
		{
			nodes:       3,
			gpusPerNode: 5,
			expectedP:   3,
			expectedQ:   5,
		},
		{
			nodes:       1,
			gpusPerNode: 15,
			expectedP:   3,
			expectedQ:   5,
		},
	}

	for _, tt := range tests {
		suite.Run(fmt.Sprintf("%d = %d * %d", tt.gpusPerNode, tt.expectedP, tt.expectedQ), func() {
			P, Q, err := benchmark.CalculateProcessGrid(tt.gpusPerNode, tt.nodes)

			// Assert
			suite.NoError(err)
			suite.Equal(tt.expectedP, P)
			suite.Equal(tt.expectedQ, Q)
		})
	}
}

func (suite *BenchmarkLauncherTestSuite) TestCalculateProblemSize() {
	// Arrange
	expectedProblemSize := "95000 96000 97000 98000 100000 101000 102000 103000 105000 106000 "
	memPerNode := uint64(128460)

	// Act
	_, problemSizeStr := benchmark.CalculateProblemSize(memPerNode, 1)

	// Assert
	suite.Equal(expectedProblemSize, problemSizeStr)
}

func TestBenchmarkLauncherTestSuite(t *testing.T) {
	suite.Run(t, &BenchmarkLauncherTestSuite{})
}
