package metascheduler

import (
	"context"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

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
	// GetProviderAddress fetches the provider public address
	GetProviderAddress() common.Address
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
		nodes uint64,
		cpus uint64,
		gpus uint64,
		mem uint64,
		gflops float64,
	) error
}
