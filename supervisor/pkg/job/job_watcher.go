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

const claimJobMaxTimeout = time.Duration(60 * time.Second)

type Watcher struct {
	metaQueue    JobMetaQueue
	scheduler    JobScheduler
	batchFetcher JobBatchFetcher
	pollingTime  time.Duration
}

func New(
	claimer JobMetaQueue,
	scheduler JobScheduler,
	batchFetcher JobBatchFetcher,
	pollingTime time.Duration,
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
		metaQueue:    claimer,
		scheduler:    scheduler,
		batchFetcher: batchFetcher,
		pollingTime:  pollingTime,
	}
}

// Watch submits a job when the metascheduler schedule a job.
func (w *Watcher) Watch(parent context.Context) error {
	queryTicker := time.NewTicker(w.pollingTime)
	defer queryTicker.Stop()

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

				r, err := w.metaQueue.Claim(ctx)
				if err != nil {
					logger.I.Info("failed to claim a job", zap.Error(err))
					done <- err
				} else {
					logger.I.Info("claimed a job", zap.Any("event", r))
					resp <- r
				}
			}(ctx)

			// Await for the claim response
			select {
			case r := <-resp:
				// Reject the job if the time limit is incorrect
				if r.MaxDurationMinute <= 0 {
					logger.I.Error(
						"refuse job because the time limit is invalid",
						zap.Any("claim_resp", r),
					)
					if err := w.metaQueue.RefuseJob(ctx, r.JobId); err != nil {
						logger.I.Error("failed to refuse a job", zap.Error(err))
					}
					return
				}

				// Fetch the job script
				body, err := w.batchFetcher.Fetch(ctx, r.JobDefinition.BatchLocationHash)
				if err != nil {
					logger.I.Error("slurm fetch job body failed", zap.Error(err))
					if err := w.metaQueue.RefuseJob(ctx, r.JobId); err != nil {
						logger.I.Error("failed to refuse a job", zap.Error(err))
					}
					return
				}

				job := eth.JobDefinitionMapToSlurm(r.JobDefinition, r.MaxDurationMinute, body)
				req := &slurm.SubmitJobRequest{
					Name:          hex.EncodeToString(r.JobId[:]),
					User:          r.CustomerAddr.Hex(),
					JobDefinition: &job,
				}

				// Submit the job script
				slurmJobID, err := w.scheduler.Submit(ctx, req)
				if err != nil {
					logger.I.Error("slurm submit job failed", zap.Error(err))
					if err := w.metaQueue.RefuseJob(ctx, r.JobId); err != nil {
						logger.I.Error("failed to refuse a job", zap.Error(err))
					}
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

		// Await for ticking
		select {
		case <-parent.Done():
			return parent.Err()
		case <-queryTicker.C:
		}
	} // for loop
}
