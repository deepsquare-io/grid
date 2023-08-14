package gc

import (
	"context"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
)

type GC struct {
	ms        metascheduler.MetaScheduler
	scheduler scheduler.Scheduler
}

func NewGC(
	ms metascheduler.MetaScheduler,
	scheduler scheduler.Scheduler,
) *GC {
	return &GC{
		ms:        ms,
		scheduler: scheduler,
	}
}

func (gc *GC) Loop(ctx context.Context) error {
	for {
		select {
		case <-time.After(15 * time.Minute):
			logger.I.Info("running gc")
			jobs, err := gc.FindUnhandledJobs(ctx)
			if err != nil {
				logger.I.Error("gc failed")
				continue
			}
			for _, job := range jobs {
				logger.I.Warn(
					"putting zombie job to FAILED",
					zap.String("jobID", hexutil.Encode(job.JobID[:])),
				)
				if err := gc.ms.SetJobStatus(ctx, job.JobID, metascheduler.JobStatusFailed, 0); err != nil {
					logger.I.Error("failed to put zombie job in FAILED", zap.Error(err))
				}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// FindUnhandledJobs matches jobs from the metascheduler to the scheduler.
//
// Jobs that doesn't appear in squeue and is RUNNING on the metascheduler is considered as a zombie.
func (gc *GC) FindUnhandledJobs(ctx context.Context) (jobs []*metascheduler.Job, err error) {
	it, err := gc.ms.GetJobs(ctx)
	if err != nil {
		logger.I.Error("GetJobs failed", zap.Error(err))
		return jobs, err
	}

	for it.Next(ctx) {
		if it.Job.Status == metascheduler.JobStatusRunning {
			id, err := gc.scheduler.FindRunningJobByName(
				ctx,
				hexutil.Encode(it.Job.JobID[:]),
				strings.ToLower(it.Job.CustomerAddr.Hex()),
			)
			if err != nil {
				logger.I.Warn(
					"FindRunningJobByName failed",
					zap.Error(err),
					zap.String("name", hexutil.Encode(it.Job.JobID[:])),
					zap.String("user", strings.ToLower(it.Job.CustomerAddr.Hex())),
				)
				return jobs, err
			}
			if id == 0 {
				logger.I.Warn(
					"found zombie job",
					zap.String("ID", hexutil.Encode(it.Job.JobID[:])),
					zap.Error(err),
				)
				jobs = append(jobs, it.Job)
			}
		}
	}

	return jobs, nil
}
