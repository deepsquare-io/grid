package job

import (
	"context"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"go.uber.org/zap"
)

const pollingTime = time.Duration(5 * time.Second)
const claimJobMaxTimeout = time.Duration(60 * time.Second)

type JobClaimer interface {
	Claim(ctx context.Context) (*metascheduler.MetaSchedulerClaimNextJobEvent, error)
}

type JobScheduler interface {
	Submit(req *slurm.SubmitJobRequest) (int, error)
}

type JobBatchFetcher interface {
	Fetch(ctx context.Context, hash string) (string, error)
}

type Watcher struct {
	claimer      JobClaimer
	scheduler    JobScheduler
	batchFetcher JobBatchFetcher
}

func New(
	claimer JobClaimer,
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
		func() {
			ctx, cancel := context.WithTimeout(parent, claimJobMaxTimeout)
			defer cancel()

			go func() {
				r, err := w.claimer.Claim(ctx)
				if err != nil {
					logger.I.Error("claimed a job", zap.Any("event", r))
					done <- err
				} else {
					resp <- r
				}
			}()

			select {
			case r := <-resp:
				body, err := w.batchFetcher.Fetch(ctx, r.JobDefinition.BatchLocationHash)
				if err != nil {
					logger.I.Error("slurm fetch job body failed", zap.Error(err))
					return
				}
				job := eth.JobDefinitionMapToSlurm(r.JobDefinition, r.MaxDurationMinute, body)
				req := &slurm.SubmitJobRequest{
					Name:          string(r.JobId[:]),
					User:          r.CustomerAddr.String(),
					JobDefinition: &job,
				}
				slurmJobID, err := w.scheduler.Submit(req)
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
					logger.I.Error("claimJob failed", zap.Error(err))
				}

			case <-ctx.Done():
				logger.I.Warn("claimJob context closed", zap.Error(ctx.Err()))
			}
		}()

		time.Sleep(pollingTime)
	}
}
