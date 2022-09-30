package job

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"go.uber.org/zap"
)

const pollingTime = time.Duration(5 * time.Second)
const claimJobMaxTimeout = time.Duration(60 * time.Second)

type JobMetaQueue interface {
	Claim(ctx context.Context) (*metascheduler.MetaSchedulerClaimNextJobEvent, error)
}

type JobScheduler interface {
	HealthCheck(ctx context.Context) error
	Submit(ctx context.Context, req *slurm.SubmitJobRequest) (int, error)
}

type JobBatchFetcher interface {
	Fetch(ctx context.Context, hash string) (string, error)
}

type Watcher struct {
	claimer      JobMetaQueue
	scheduler    JobScheduler
	batchFetcher JobBatchFetcher
}

func New(
	claimer JobMetaQueue,
	scheduler JobScheduler,
	batchFetcher JobBatchFetcher,
) *Watcher {
	if claimer == nil {
		logger.I.Panic("claimer is nil")
	}
	if scheduler == nil {
		logger.I.Panic("scheduler is nil")
	}
	if batchFetcher == nil {
		logger.I.Panic("batchFetcher is nil")
	}
	return &Watcher{
		claimer:      claimer,
		scheduler:    scheduler,
		batchFetcher: batchFetcher,
	}
}

// Watch submits a job when the metascheduler schedule a job.
func (w *Watcher) Watch(parent context.Context) error {
	resp := make(chan *metascheduler.MetaSchedulerClaimNextJobEvent)
	done := make(chan error)
	for {
		func(parent context.Context) {
			ctx, cancel := context.WithTimeout(parent, claimJobMaxTimeout)
			defer cancel()

			go func(ctx context.Context) {
				// Slurm healthcheck first
				err := w.scheduler.HealthCheck(ctx)
				if err != nil {
					done <- err
					return
				}

				r, err := w.claimer.Claim(ctx)
				if err != nil {
					done <- err
				} else {
					logger.I.Info("claimed a job", zap.Any("event", r))
					resp <- r
				}
			}(ctx)

			select {
			case r := <-resp:
				body, err := w.batchFetcher.Fetch(ctx, r.JobDefinition.BatchLocationHash)
				if err != nil {
					logger.I.Error("slurm fetch job body failed", zap.Error(err))
					return
				}
				job := eth.JobDefinitionMapToSlurm(r.JobDefinition, r.MaxDurationMinute, body)
				req := &slurm.SubmitJobRequest{
					Name:          hex.EncodeToString(r.JobId[:]),
					User:          r.CustomerAddr.String(),
					JobDefinition: &job,
				}
				slurmJobID, err := w.scheduler.Submit(ctx, req)
				if err != nil {
					logger.I.Error("slurm submit job failed", zap.Error(err))
					return
				}
				logger.I.Info(
					"submitted a job successfully",
					zap.Int("JobID", slurmJobID),
					zap.Any("Req", req),
				)
			case err := <-done:
				if err != nil {
					logger.I.Error("watcher failed", zap.Error(err))
				}

			case <-ctx.Done():
				logger.I.Warn("watcher context closed", zap.Error(ctx.Err()))
			}
		}(parent)

		time.Sleep(pollingTime)
	}
}
