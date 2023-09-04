package metascheduler

import (
	"context"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"go.uber.org/zap"
)

type Job struct {
	JobID            [32]byte
	Status           JobStatus
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       metaschedulerabi.JobDefinition
	Cost             metaschedulerabi.JobCost
	Time             metaschedulerabi.JobTime
	JobName          [32]byte
	HasCancelRequest bool
	LastError        string
}

func FromStructToJob(s metaschedulerabi.Job) *Job {
	return &Job{
		JobID:            s.JobId,
		Status:           JobStatus(s.Status),
		CustomerAddr:     s.CustomerAddr,
		ProviderAddr:     s.ProviderAddr,
		Definition:       s.Definition,
		Cost:             s.Cost,
		Time:             s.Time,
		JobName:          s.JobName,
		HasCancelRequest: s.HasCancelRequest,
	}
}

type ProviderJobIterator struct {
	Job *Job
	*metaschedulerabi.MetaSchedulerClaimJobEventIterator
	client          MetaScheduler
	providerAddress common.Address
}

func (it *ProviderJobIterator) Next(ctx context.Context) bool {
	for it.MetaSchedulerClaimJobEventIterator.Next() {
		if it.Event.ProviderAddr != it.providerAddress {
			continue
		}

		job, err := it.client.GetJob(ctx, it.Event.JobId)
		if err != nil {
			logger.I.Error("GetJob failed", zap.Error(err))
			return false
		}
		if job.ProviderAddr != it.providerAddress {
			continue
		}

		it.Job = job
		return true
	}

	return false
}

type MetaScheduler interface {
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
		claimNextTopUpJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent,
		claimNextCancellingJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
		claimJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent,
	) (event.Subscription, error)
	// GetProviderAddress fetches the provider public address.
	GetProviderAddress() common.Address
	// GetOldInfo fetches the last registration provider information.
	GetOldInfo(ctx context.Context) (*metaschedulerabi.Provider, error)
	// GetJobStatus fetches the job status.
	GetJobStatus(ctx context.Context, jobID [32]byte) (JobStatus, error)
	SetJobStatus(
		ctx context.Context,
		jobID [32]byte,
		status JobStatus,
		jobDurationMinute uint64,
	) error
	Register(
		ctx context.Context,
		hardware metaschedulerabi.ProviderHardware,
		prices metaschedulerabi.ProviderPrices,
		labels []metaschedulerabi.Label,
	) error
	// GetJob fetches a job.
	GetJob(ctx context.Context, jobID [32]byte) (*Job, error)
	// GetJobs fetches the jobs of the provider.
	GetJobs(ctx context.Context) (*ProviderJobIterator, error)
}
