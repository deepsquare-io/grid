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

// Watch job submit a job when the metascheduler schedule a job.
func Watch(ctx context.Context, s *eth.DataSource, j *slurm.Service) error {
	resp := make(chan *metascheduler.MetaSchedulerClaimNextJobEvent)
	done := make(chan error)
	for {
		func() {
			ctx, cancel := context.WithTimeout(ctx, time.Duration(60*time.Second))
			defer cancel()

			go func() {
				r, err := s.ClaimJob(ctx)
				resp <- r
				done <- err
			}()

			select {
			case r := <-resp:
				// TODO: fetch sbatch here
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
				slurmJobID, err := j.SubmitJob(req)
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

		// TODO: extract variable
		time.Sleep(time.Duration(10 * time.Second))
	}
}
