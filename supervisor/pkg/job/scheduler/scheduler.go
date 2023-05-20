package scheduler

import (
	"context"
)

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

type CancelRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
}

type SubmitRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation.
	User string
	// JobDefinition specifies the job allocations
	*JobDefinition
}

type TopUpRequest struct {
	// Name of the job
	Name string
	// AdditionalTime is the number of minutes to be added
	AdditionalTime uint64
}

type FindRunningJobByNameRequest struct {
	// Name of the job
	Name string
	// User is a UNIX User used for impersonation. This user should be SLURM admin.
	User string
}

type Scheduler interface {
	// HealthCheck verifies if the scheduler accepts jobs.
	HealthCheck(ctx context.Context) error
	// Submit a job to the scheduler.
	Submit(ctx context.Context, req *SubmitRequest) (string, error)
	// CancelJob kills a job.
	CancelJob(ctx context.Context, req *CancelRequest) error
	// TopUp increases the time limit of a job.
	TopUp(ctx context.Context, req *TopUpRequest) error
}

type Executor interface {
	ExecAs(ctx context.Context, user string, cmd string) (string, error)
}
