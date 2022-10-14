package job

import (
	"context"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
)

type JobMetaQueue interface {
	// Claim a job for scheduling.
	Claim(ctx context.Context) (*metascheduler.MetaSchedulerClaimNextJobEvent, error)
	// Refuse a job for metascheduling.
	RefuseJob(ctx context.Context, jobID [32]byte) error
}

type JobScheduler interface {
	// HealthCheck verifies if the scheduler accepts jobs.
	HealthCheck(ctx context.Context) error
	// Submit a job to the scheduler.
	Submit(ctx context.Context, req *slurm.SubmitJobRequest) (int, error)
}

type JobBatchFetcher interface {
	// Fetch a job batch content.
	Fetch(ctx context.Context, hash string) (string, error)
}
