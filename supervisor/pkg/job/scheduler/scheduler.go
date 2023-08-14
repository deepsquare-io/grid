package scheduler

import (
	"context"
)

type JobDefinition struct {
	// TimeLimit is a time allocation which at the end kills the running job.
	//
	// TimeLimit is in minutes.
	TimeLimit uint64
	// NTasks indicates the number parallel tasks.
	NTasks uint64
	// NTasksPerNode indicates the number parallel tasks per node.
	//
	// If NTasks is not 0, NTasksPerNode is a maximum. Otherwise, it is a scrictly the number of tasks per node.
	NTasksPerNode uint64
	// MinNodes indicates the minimum number of allocated node.
	MinNodes uint64
	// MaxNodes indicates the maximum number of allocated node.
	// MinNodes is required.
	MaxNodes uint64
	// GPUsPerTask indicates the number of requested GPU.
	GPUsPerTask *uint64
	// GPUsPerNode indicates the number of requested GPUs per node.
	GPUsPerNode uint64
	// CPUs indicates the number of requested CPU.
	CPUsPerTask uint64
	// CPUs indicates the minimum number of CPU per node.
	CPUsPerNode uint64
	// MemoryPerCpu indicates the number of requested MB of memory.
	MemoryPerCPU uint64
	// Memory indicates the number of requested MB of memory.
	Memory *uint64
	// Body of the job, in a sbatch script.
	Body string
	// Wait for the job to end. The exit code of the sbatch will be the exit code
	// of the job.
	Wait bool
}

type SubmitRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	// Prefix is appended to the log and comment.
	Prefix string
	// JobDefinition specifies the job allocations
	*JobDefinition
}

type findSpecOptions struct {
	onlyResponding bool
}
type FindSpecOption func(*findSpecOptions)

func newFindSpecOptions(opts ...FindSpecOption) *findSpecOptions {
	o := &findSpecOptions{
		onlyResponding: false,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithOnlyResponding() FindSpecOption {
	return func(fso *findSpecOptions) {
		fso.onlyResponding = true
	}
}

type Scheduler interface {
	// HealthCheck verifies if the scheduler accepts jobs.
	HealthCheck(ctx context.Context) error
	// Submit a job to the scheduler.
	Submit(ctx context.Context, req *SubmitRequest) (string, error)
	// CancelJob kills a job.
	CancelJob(ctx context.Context, name string, user string) error
	// TopUp increases the time limit in minutes of a job.
	TopUp(ctx context.Context, name string, additionalTime uint64) error
	// Find the memory (MB) per node
	FindMemPerNode(ctx context.Context, opts ...FindSpecOption) ([]uint64, error)
	// Find the GPU per node.
	FindGPUsPerNode(ctx context.Context, opts ...FindSpecOption) ([]uint64, error)
	// Find the CPU per node.
	FindCPUsPerNode(ctx context.Context, opts ...FindSpecOption) ([]uint64, error)
	// Find the total number of memory (MB) available
	FindTotalMem(ctx context.Context) (uint64, error)
	// Find the total number of GPUs available.
	FindTotalGPUs(ctx context.Context) (uint64, error)
	// Find the total number of CPUs available.
	FindTotalCPUs(ctx context.Context) (uint64, error)
	// Find the total number of nodes available.
	FindTotalNodes(ctx context.Context, opts ...FindSpecOption) (uint64, error)
	// FindRunningJobByName find a running job using squeue.
	//
	// Returns 0 if not found.
	FindRunningJobByName(
		ctx context.Context,
		name string,
		user string,
	) (int, error)
}

type Executor interface {
	ExecAs(ctx context.Context, user string, cmd string) (string, error)
}
