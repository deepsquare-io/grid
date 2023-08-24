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
		logger.I.Info("running gc")
		if err := gc.FindAndCancelUnhandledJobs(ctx); err != nil {
			logger.I.Error("gc failed")
		}
		select {
		case <-time.After(15 * time.Minute):
			// Pass
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// FindAndCancelUnhandledJobs matches jobs from the metascheduler to the scheduler.
//
// Jobs that doesn't appear in squeue and is RUNNING on the metascheduler is considered as a zombie.
func (gc *GC) FindAndCancelUnhandledJobs(
	ctx context.Context,
) (err error) {
	it, err := gc.ms.GetJobs(ctx)
	if err != nil {
		logger.I.Error("GetJobs failed", zap.Error(err))
		return err
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
				return err
			}
			if id == 0 {
				logger.I.Warn(
					"found zombie job, putting zombie job to FAILED",
					zap.String("jobID", hexutil.Encode(it.Job.JobID[:])),
				)
				if err := gc.ms.SetJobStatus(ctx, it.Job.JobID, metascheduler.JobStatusFailed, 0); err != nil {
					logger.I.Error(
						"failed to put zombie job in FAILED",
						zap.Error(err),
						zap.Any("job", it.Job),
					)
				}
			}
		}
	}

	return nil
}
