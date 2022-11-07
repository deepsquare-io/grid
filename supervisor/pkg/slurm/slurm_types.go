package slurm

import "context"

type JobDefinition struct {
	// TimeLimit is a time allocation which at the end kills the running job.
	//
	// TimeLimit is in minutes.
	TimeLimit uint64
	// NTasks indicates the number
	NTasks uint64
	// GPUsPerTask indicates the number of requested GPU.
	GPUsPerTask uint64
	// CPUs indicates the number of requested CPU.
	CPUsPerTask uint64
	// MemoryPerCpu indicates the number of requested MB of memory.
	MemoryPerCPU uint64
	// Body of the job, in a sbatch script.
	Body string
}

type Executor interface {
	ExecAs(ctx context.Context, user string, cmd string) (string, error)
}
