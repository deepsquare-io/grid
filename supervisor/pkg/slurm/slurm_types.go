package slurm

import "context"

type JobDefinition struct {
	// TimeLimit is a time allocation which at the end kills the running job.
	//
	// TimeLimit is in minutes.
	TimeLimit uint64
	// NTasks indicates the number
	NTasks uint64
	// GPUsPerNode indicates the number of requested GPU.
	GPUsPerNode uint64
	// CPUs indicates the number of requested CPU.
	CPUsPerTask uint64
	// MemoryPerNode indicates the number of requested MB of memory.
	MemoryPerNode uint64
	// Body of the job, in a sbatch script.
	Body string
}

type Executor interface {
	ExecAs(ctx context.Context, user string, cmd string) (string, error)
}
