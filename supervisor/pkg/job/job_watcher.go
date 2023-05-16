package job

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils/try"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// Watch incoming jobs and handles it.
func (w *Watcher) Watch(parent context.Context) error {
	// Ticker for claim
	queryTicker := time.NewTicker(w.pollingTime)
	defer queryTicker.Stop()

	claimNextJobEvents := make(chan *metascheduler.MetaSchedulerClaimJobEvent, 100)
	defer close(claimNextJobEvents)
	claimNextCancellingJobEvents := make(
		chan *metascheduler.MetaSchedulerClaimNextCancellingJobEvent,
		100,
	)
	defer close(claimNextCancellingJobEvents)
	claimNextTopUpJobEvents := make(chan *metascheduler.MetaSchedulerClaimNextTopUpJobEvent, 100)
	defer close(claimNextTopUpJobEvents)
	sub, err := w.metaQueue.WatchEvents(
		parent,
		claimNextTopUpJobEvents,
		claimNextCancellingJobEvents,
		claimNextJobEvents,
	)
	if err != nil {
		return err
	}
	logger.I.Debug("Watching events...")

	for {
		select {
		case <-parent.Done():
			return parent.Err()

		// ClaimNextJobEvents handling
		case event := <-claimNextJobEvents:
			if !strings.EqualFold(
				event.ProviderAddr.Hex(),
				w.metaQueue.GetProviderAddress().Hex(),
			) {
				continue
			}
			logger.I.Debug("Received claimNextJobEvent...", zap.Any("event", event))
			go w.handleClaimNextJob(parent, event)

		// ClaimNextCancellingJobEvents handling
		case event := <-claimNextCancellingJobEvents:
			if !strings.EqualFold(
				event.ProviderAddr.Hex(),
				w.metaQueue.GetProviderAddress().Hex(),
			) {
				continue
			}
			logger.I.Debug("Received claimNextCancellingJobEvent...", zap.Any("event", event))
			go w.handleClaimNextCancellingJobEvent(parent, event)

		// ClaimNextTopUpJobEvents handling
		case event := <-claimNextTopUpJobEvents:
			if !strings.EqualFold(
				event.ProviderAddr.Hex(),
				w.metaQueue.GetProviderAddress().Hex(),
			) {
				continue
			}
			logger.I.Debug("Received claimNextTopUpJobEvent...", zap.Any("event", event))
			go w.handleClaimNextTopUpEvent(parent, event)

		// Claim indefinitely
		case <-queryTicker.C:
			go w.handleClaimIndefinitely(parent)

		case err := <-sub.Err():
			return err
		}
	}
}

func (w *Watcher) handleClaimIndefinitely(parent context.Context) {
	done := make(chan error)
	ctx, cancel := context.WithTimeout(parent, claimJobMaxTimeout)
	defer cancel()

	go func(ctx context.Context) {
		// One shot
		defer close(done)

		// Slurm healthcheck first
		if err := w.scheduler.HealthCheck(ctx); err != nil {
			done <- err
			return
		}

		if err := w.metaQueue.Claim(ctx); err != nil {
			logger.I.Info("failed to claim a job", zap.Error(err))
			done <- err
			return
		}

		if err := w.metaQueue.ClaimCancelling(ctx); err != nil {
			logger.I.Info("failed to claim cancelling job", zap.Error(err))
			done <- err
			return
		}

		if err := w.metaQueue.ClaimTopUp(ctx); err != nil {
			logger.I.Info("failed to claim topup job", zap.Error(err))
			done <- err
			return
		}
	}(ctx)

	// Await for the claim response
	select {
	case err := <-done:
		if err != nil {
			logger.I.Error("ClaimIndefinitely failed", zap.Error(err))
		}

	case <-ctx.Done():
		logger.I.Warn("ClaimIndefinitely context closed", zap.Error(ctx.Err()))
	}
}

func (w *Watcher) handleClaimNextJob(
	ctx context.Context,
	event *metascheduler.MetaSchedulerClaimJobEvent,
) {
	if event == nil {
		logger.I.Error(
			"job is nil, we didn't find a job",
		)
		return
	}
	// Reject the job if the time limit is incorrect
	if event.MaxDurationMinute <= 0 {
		logger.I.Error(
			"set job failed because the time limit is invalid",
			zap.Any("claim_resp", event),
		)
		if err := w.metaQueue.SetJobStatus(ctx, event.JobId, eth.JobStatusFailed, 0); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return
	}

	// Fetch the job script
	body, err := w.batchFetcher.Fetch(ctx, event.JobDefinition.BatchLocationHash)
	if err != nil {
		logger.I.Error("slurm fetch job body failed, setting job to failed", zap.Error(err))
		if err := w.metaQueue.SetJobStatus(ctx, event.JobId, eth.JobStatusFailed, 0); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return
	}

	definition := DefinitionFromMetascheduler(event.JobDefinition, event.MaxDurationMinute, body)
	req := &SubmitRequest{
		Name:       hexutil.Encode(event.JobId[:]),
		User:       strings.ToLower(event.CustomerAddr.Hex()),
		Definition: &definition,
	}

	// Submit the job script
	resp, err := w.scheduler.Submit(ctx, req)
	if err != nil {
		logger.I.Error("slurm submit job failed", zap.Error(err))
		if err := w.metaQueue.RefuseJob(ctx, event.JobId); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return
	}
	logger.I.Info(
		"submitted a job successfully",
		zap.String("Response", resp),
		zap.Any("Req", req),
	)
}

func (w *Watcher) handleClaimNextCancellingJobEvent(
	ctx context.Context,
	event *metascheduler.MetaSchedulerClaimNextCancellingJobEvent,
) {
	status, err := w.metaQueue.GetJobStatus(ctx, event.JobId)
	if err != nil {
		logger.I.Error(
			"GetJobStatus failed, abort handleClaimNextCancellingJobEvent",
			zap.Error(err),
		)
		return
	}
	if err := try.Do(
		10, 5*time.Second,
		func(_ int) error {
			if err := try.Do(
				5, 5*time.Second,
				func(_ int) error {
					return w.scheduler.CancelJob(ctx, &CancelRequest{
						Name: hexutil.Encode(event.JobId[:]),
						User: strings.ToLower(event.CustomerAddr.Hex()),
					})
				}); err != nil {
				logger.I.Error("CancelJob failed, abort handleClaimNextCancellingJobEvent", zap.Error(err))
				return err
			}

			time.Sleep(5 * time.Second)

			status, err = w.metaQueue.GetJobStatus(ctx, event.JobId)
			if err != nil {
				logger.I.Error("GetJobStatus failed, considering as Cancelled", zap.Error(err))
				status = eth.JobStatusCancelled
			}
			if status != eth.JobStatusCancelled {
				return errors.New("failed to cancel job")
			}
			return nil
		}); err != nil {
		logger.I.Error("failed to cancel, considering as CANCELLED", zap.Error(err))
		if err := w.metaQueue.SetJobStatus(ctx, event.JobId, eth.JobStatusCancelled, 0); err != nil {
			logger.I.Error("even considering as CANCELLED, it failed", zap.Error(err))
			return
		}
	}
}

func (w *Watcher) handleClaimNextTopUpEvent(
	ctx context.Context,
	event *metascheduler.MetaSchedulerClaimNextTopUpJobEvent,
) {
	if err := try.Do(5, 5*time.Second, func(_ int) error {
		return w.scheduler.TopUp(ctx, &TopUpRequest{
			Name:           hexutil.Encode(event.JobId[:]),
			AdditionalTime: event.MaxDurationMinute,
		})
	}); err != nil {
		logger.I.Error("failed to topup", zap.Error(err))
	}
}
