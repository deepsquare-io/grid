package job

import (
	"context"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

type JobMetaQueue interface {
	// Claim a job for scheduling.
	Claim(ctx context.Context) error
	// Claim cancelling calls.
	ClaimCancelling(ctx context.Context) error
	// Claim top up calls.
	ClaimTopUp(ctx context.Context) error
	// Refuse a job for metascheduling.
	RefuseJob(ctx context.Context, jobID [32]byte) error
	// WatchEvents observes the incoming ClaimNextTopUpJobEvent, ClaimNextCancellingJobEvent and ClaimJobEvent.
	WatchEvents(
		ctx context.Context,
		claimNextTopUpJobEvents chan<- *metascheduler.MetaSchedulerClaimNextTopUpJobEvent,
		claimNextCancellingJobEvents chan<- *metascheduler.MetaSchedulerClaimNextCancellingJobEvent,
		claimJobEvents chan<- *metascheduler.MetaSchedulerClaimJobEvent,
	) (event.Subscription, error)
	// GetProviderAddress fetches the provider public address
	GetProviderAddress() common.Address
	// GetJobStatus fetches the job status.
	GetJobStatus(ctx context.Context, jobID [32]byte) (eth.JobStatus, error)
	SetJobStatus(
		ctx context.Context,
		jobID [32]byte,
		status eth.JobStatus,
		jobDurationMinute uint64,
	) error
}

type JobScheduler interface {
	// HealthCheck verifies if the scheduler accepts jobs.
	HealthCheck(ctx context.Context) error
	// Submit a job to the scheduler.
	Submit(ctx context.Context, req *SubmitRequest) (string, error)
	// CancelJob kills a job.
	CancelJob(ctx context.Context, req *CancelRequest) error
	// TopUp increases the time limit of a job.
	TopUp(ctx context.Context, req *TopUpRequest) error
}

type Definition struct {
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
	*Definition
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

type JobBatchFetcher interface {
	// Fetch a job batch content.
	Fetch(ctx context.Context, hash string) (string, error)
}

func DefinitionFromMetascheduler(j metascheduler.JobDefinition, t uint64, body string) Definition {
	return Definition{
		NTasks:       j.Ntasks,
		GPUsPerTask:  j.GpuPerTask,
		CPUsPerTask:  j.CpuPerTask,
		TimeLimit:    t,
		MemoryPerCPU: j.MemPerCpu,
		Body:         body,
	}
}
