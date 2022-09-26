package job

import (
	"context"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/oracle"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"go.uber.org/zap"
)

const pollingTime = time.Duration(10 * time.Second)

type Watcher struct {
	eth    *eth.DataSource
	slurm  *slurm.Service
	oracle *oracle.DataSource
}

func New(
	eth *eth.DataSource,
	slurm *slurm.Service,
	oracle *oracle.DataSource,
) *Watcher {
	if eth == nil {
		logger.I.Panic("eth is nil")
	}
	if slurm == nil {
		logger.I.Panic("slurm is nil")
	}
	if oracle == nil {
		logger.I.Panic("oracle is nil")
	}
	return &Watcher{
		eth:    eth,
		slurm:  slurm,
		oracle: oracle,
	}
}

// Watch submits a job when the metascheduler schedule a job.
func (w *Watcher) Watch(ctx context.Context) error {
	resp := make(chan *metascheduler.MetaSchedulerClaimNextJobEvent)
	done := make(chan error)
	for {
		func() {
			ctx, cancel := context.WithTimeout(ctx, time.Duration(60*time.Second))
			defer cancel()

			go func() {
				r, err := w.eth.ClaimJob(ctx)
				resp <- r
				done <- err
			}()

			select {
			case r := <-resp:
				// TODO: do not use mock
				// w.oracle.FetchJobBatch(ctx, r.JobDefinition.BatchLocationHash)
				body := `#!/bin/sh

				srun hostname
				srun sleep infinity
				`
				job := eth.JobDefinitionMapToSlurm(r.JobDefinition, r.MaxDurationMinute, body)
				req := &slurm.SubmitJobRequest{
					Name:          string(r.JobId[:]),
					User:          r.CustomerAddr.String(),
					JobDefinition: job,
				}
				slurmJobID, err := w.slurm.SubmitJob(req)
				if err != nil {
					logger.I.Error("slurm submit job failed", zap.Error(err))
				} else {
					logger.I.Info(
						"submitted a job successfully",
						zap.Int("JobID", slurmJobID),
						zap.Any("Req", req),
					)
				}
			case err := <-done:
				if err != nil {
					logger.I.Error("claimJob failed", zap.Error(err))
				}

			case <-ctx.Done():
				logger.I.Error("claimJob timed out", zap.Error(ctx.Err()))
			}
		}()

		time.Sleep(pollingTime)
	}
}
