package job

import (
	"context"
	"encoding/hex"
	"strings"
	"sync"
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

// Watch incoming jobs and handles it.
func (w *Watcher) Watch(parent context.Context) error {
	var wg sync.WaitGroup
	errChan := make(chan error)

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		err := w.WatchClaimNextJob(ctx)
		if err != nil {
			logger.I.Error("WatchClaimNextJob failed", zap.Error(err))
			errChan <- err
		}
		wg.Done()
	}(parent, &wg)

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		err := w.ClaimNextJobIndefinitely(ctx)
		if err != nil {
			logger.I.Error("WatchClaimNextJob failed", zap.Error(err))
			errChan <- err
		}
		wg.Done()
	}(parent, &wg)

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for e := range errChan {
		if e != nil {
			return e
		}
	}

	return nil
}

func (w *Watcher) ClaimNextJobIndefinitely(parent context.Context) error {
	queryTicker := time.NewTicker(w.pollingTime)
	defer queryTicker.Stop()

	for {
		func(parent context.Context) {
			done := make(chan error)
			defer close(done)
			ctx, cancel := context.WithTimeout(parent, claimJobMaxTimeout)
			defer cancel()

			go func(ctx context.Context) {
				// Slurm healthcheck first
				if err := w.scheduler.HealthCheck(ctx); err != nil {
					done <- err
					return
				}

				if err := w.metaQueue.Claim(ctx); err != nil {
					logger.I.Info("failed to claim a job", zap.Error(err))
					done <- err
				}
				done <- nil
			}(ctx)

			// Await for the claim response
			select {
			case err := <-done:
				if err != nil {
					logger.I.Error("ClaimNextJobIndefinitely failed", zap.Error(err))
				}

			case <-ctx.Done():
				logger.I.Warn("ClaimNextJobIndefinitely context closed", zap.Error(ctx.Err()))
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

func (w *Watcher) WatchClaimNextJob(parent context.Context) error {
	events := make(chan *metascheduler.MetaSchedulerClaimNextJobEvent)
	sub, err := w.metaQueue.WatchClaimNextJobEvent(parent, events)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-parent.Done():
			logger.I.Warn("WatchClaimNextJobEvent context is done")
			return nil
		case err := <-sub.Err():
			logger.I.Error("WatchClaimNextJobEvent thrown an error", zap.Error(err))
			return err
		case event := <-events:
			err := w.handleClaimNextJob(parent, event)
			if err != nil {
				return err
			}
		}
	}
}

func (w *Watcher) handleClaimNextJob(ctx context.Context, event *metascheduler.MetaSchedulerClaimNextJobEvent) error {
	if event == nil {
		logger.I.Warn(
			"job is nil, we didn't find a job",
		)
		return nil
	}
	// Reject the job if the time limit is incorrect
	if event.MaxDurationMinute <= 0 {
		logger.I.Error(
			"refuse job because the time limit is invalid",
			zap.Any("claim_resp", event),
		)
		if err := w.metaQueue.RefuseJob(ctx, event.JobId); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return nil
	}

	// Fetch the job script
	body, err := w.batchFetcher.Fetch(ctx, event.JobDefinition.BatchLocationHash)
	if err != nil {
		logger.I.Error("slurm fetch job body failed", zap.Error(err))
		if err := w.metaQueue.RefuseJob(ctx, event.JobId); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return nil
	}

	job := eth.JobDefinitionMapToSlurm(event.JobDefinition, event.MaxDurationMinute, body)
	req := &slurm.SubmitJobRequest{
		Name:          hex.EncodeToString(event.JobId[:]),
		User:          strings.ToLower(event.CustomerAddr.Hex()),
		JobDefinition: &job,
	}

	// Submit the job script
	slurmJobID, err := w.scheduler.Submit(ctx, req)
	if err != nil {
		logger.I.Error("slurm submit job failed", zap.Error(err))
		if err := w.metaQueue.RefuseJob(ctx, event.JobId); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return nil
	}
	logger.I.Info(
		"submitted a job successfully",
		zap.Int("JobID", slurmJobID),
		zap.Any("Req", req),
	)
	return nil
}
