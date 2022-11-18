package job

import (
	"context"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/ethereum/go-ethereum/event"
)

type JobMetaQueue interface {
	// Claim a job for scheduling.
	Claim(ctx context.Context) error
	// Refuse a job for metascheduling.
	RefuseJob(ctx context.Context, jobID [32]byte) error
	// WatchClaimNextJobEvent observes the incoming ClaimNextJobEvents.
	WatchClaimNextJobEvent(
		ctx context.Context,
		sink chan<- *metascheduler.MetaSchedulerClaimNextJobEvent,
	) (event.Subscription, error)
}

type JobScheduler interface {
	// HealthCheck verifies if the scheduler accepts jobs.
	HealthCheck(ctx context.Context) error
	// Submit a job to the scheduler.
	Submit(ctx context.Context, req *slurm.SubmitJobRequest) (string, error)
}

type JobBatchFetcher interface {
	// Fetch a job batch content.
	Fetch(ctx context.Context, hash string) (string, error)
}
