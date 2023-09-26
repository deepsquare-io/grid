package gc

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils/try"
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
		switch it.Job.Status {
		case metascheduler.JobStatusRunning:
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
					"found zombie job, putting zombie job to PANIC",
					zap.String("jobID", hexutil.Encode(it.Job.JobID[:])),
				)
				if err := gc.ms.SetJobStatus(
					ctx,
					it.Job.JobID,
					metascheduler.JobStatusPanicked,
					0,
					metascheduler.SetJobStatusWithError(errors.New("provider lost the job")),
				); err != nil {
					logger.I.Error(
						"failed to put zombie job in PANIC",
						zap.Error(err),
						zap.Any("job", it.Job),
					)
				}
			}
		case metascheduler.JobStatusScheduled:
			if err := try.Do(10, 5*time.Second, func(try int) error {
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
						"possible zombie job",
						zap.String("jobID", hexutil.Encode(it.Job.JobID[:])),
					)
					return fmt.Errorf("zombie job: %s", hexutil.Encode(it.Job.JobID[:]))
				}
				return nil
			}); err != nil {
				logger.I.Warn(
					"putting zombie job in PANIC",
					zap.Error(err),
					zap.Any("job", it.Job),
				)
				if err := gc.ms.SetJobStatus(
					ctx,
					it.Job.JobID,
					metascheduler.JobStatusPanicked,
					0,
					metascheduler.SetJobStatusWithError(errors.New("provider lost the job")),
				); err != nil {
					logger.I.Error(
						"failed to put zombie job in PANIC",
						zap.Error(err),
						zap.Any("job", it.Job),
					)
				}
			}
		}
	}

	return nil
}
