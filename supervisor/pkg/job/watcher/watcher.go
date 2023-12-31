// Copyright (C) 2023 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package watcher

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	loggerv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/logger/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/gridlogger"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/sbatch"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils/try"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
)

const claimJobMaxTimeout = time.Duration(60 * time.Second)

type Watcher struct {
	metascheduler    metascheduler.MetaScheduler
	scheduler        scheduler.Scheduler
	sbatch           sbatch.Client
	pollingTime      time.Duration
	resourceManager  *lock.ResourceManager
	gridLoggerDialer gridlogger.Dialer
	claimMutex       sync.Mutex
}

func New(
	metascheduler metascheduler.MetaScheduler,
	scheduler scheduler.Scheduler,
	sbatch sbatch.Client,
	pollingTime time.Duration,
	resourceManager *lock.ResourceManager,
	gridLoggerDialer gridlogger.Dialer,
) *Watcher {
	if metascheduler == nil {
		logger.I.Panic("metascheduler is nil")
	}
	if scheduler == nil {
		logger.I.Panic("scheduler is nil")
	}
	if sbatch == nil {
		logger.I.Panic("sbatch is nil")
	}
	if resourceManager == nil {
		logger.I.Panic("resourceManager is nil")
	}
	if gridLoggerDialer == nil {
		logger.I.Panic("gridLoggerDialer is nil")
	}
	return &Watcher{
		metascheduler:    metascheduler,
		scheduler:        scheduler,
		sbatch:           sbatch,
		pollingTime:      pollingTime,
		resourceManager:  resourceManager,
		gridLoggerDialer: gridLoggerDialer,
	}
}

func (w *Watcher) logToUser(
	ctx context.Context,
	endpoint string,
	jobName string,
	user string,
	content []byte,
) error {
	log := logger.I.With(
		zap.String("endpoint", endpoint),
		zap.String("jobName", jobName),
		zap.String("user", user),
		zap.ByteString("content", content),
	)
	log.Debug("sending log to user")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	c, close, err := w.gridLoggerDialer.DialContext(ctx, endpoint)
	if err != nil {
		log.Error("log to user: dial fail", zap.Error(err))
		return err
	}
	defer func() {
		_ = close()
	}()

	// Iterate line by line
	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		if err := c.Send(&loggerv1alpha1.WriteRequest{
			LogName:   jobName,
			Data:      scanner.Bytes(),
			User:      user,
			Timestamp: time.Now().UnixNano(),
		}); err != nil {
			log.Error("log to user: send fail", zap.Error(err))
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("log to user: scanner thrown an error", zap.Error(err))
	}

	if _, err := c.CloseAndRecv(); err != nil {
		log.Error("log to user: failed close gracefully", zap.Error(err))
	}
	log.Debug("successfully logged to user")
	return nil
}

// Watch incoming jobs and handles it.
func (w *Watcher) Watch(parent context.Context) error {
	// Ticker for claim
	queryTicker := time.NewTicker(w.pollingTime)
	defer queryTicker.Stop()

	claimNextJobEvents := make(chan *metaschedulerabi.MetaSchedulerClaimJobEvent, 100)
	defer close(claimNextJobEvents)
	claimNextCancellingJobEvents := make(
		chan *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
		100,
	)
	defer close(claimNextCancellingJobEvents)
	claimNextTopUpJobEvents := make(chan *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, 100)
	defer close(claimNextTopUpJobEvents)
	sub, err := w.metascheduler.WatchEvents(
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
				w.metascheduler.GetProviderAddress().Hex(),
			) {
				continue
			}
			logger.I.Debug("Received claimNextJobEvent...", zap.Any("event", event))
			go w.handleClaimNextJob(parent, event)

		// ClaimNextCancellingJobEvents handling
		case event := <-claimNextCancellingJobEvents:
			if !strings.EqualFold(
				event.ProviderAddr.Hex(),
				w.metascheduler.GetProviderAddress().Hex(),
			) {
				continue
			}
			logger.I.Debug("Received claimNextCancellingJobEvent...", zap.Any("event", event))
			go w.handleClaimNextCancellingJobEvent(parent, event)

		// ClaimNextTopUpJobEvents handling
		case event := <-claimNextTopUpJobEvents:
			if !strings.EqualFold(
				event.ProviderAddr.Hex(),
				w.metascheduler.GetProviderAddress().Hex(),
			) {
				continue
			}
			logger.I.Debug("Received claimNextTopUpJobEvent...", zap.Any("event", event))
			go w.handleClaimNextTopUpEvent(parent, event)

		// Claim indefinitely
		case <-queryTicker.C:
			if w.claimMutex.TryLock() {
				w.claimIndefinitely(parent)
			}

		case err := <-sub.Err():
			return err
		}
	}
}

func (w *Watcher) claimIndefinitely(parent context.Context) {
	defer w.claimMutex.Unlock()
	ctx, cancel := context.WithTimeout(parent, claimJobMaxTimeout)
	defer cancel()

	// Slurm healthcheck
	if err := w.scheduler.HealthCheck(ctx); err != nil {
		logger.I.Error("failed healthcheck", zap.Error(err))
		return
	}

	// MetaScheduler healthcheck
	val, err := w.metascheduler.IsRequestNewJobEnabled(ctx)
	if err != nil {
		logger.I.Error("failed to check metascheduler health", zap.Error(err))
		return
	} else if !val {
		logger.I.Panic(`The smart-contract has closed the submission of a new job.
The smart-contract is therefore deprecated.
Please check for new smart-contract versions.`)
	}

	if err := w.metascheduler.Claim(ctx); err != nil {
		logger.I.Error("failed to claim a job", zap.Error(err))
		return
	}

	if err := w.metascheduler.ClaimCancelling(ctx); err != nil {
		logger.I.Error("failed to claim cancelling job", zap.Error(err))
		return
	}

	if err := w.metascheduler.ClaimTopUp(ctx); err != nil {
		logger.I.Error("failed to claim topup job", zap.Error(err))
		return
	}
}

func (w *Watcher) handleClaimNextJob(
	ctx context.Context,
	event *metaschedulerabi.MetaSchedulerClaimJobEvent,
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
		if err := w.metascheduler.SetJobStatus(ctx, event.JobId, metascheduler.JobStatusFailed, 0); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return
	}

	// Fetch the job script
	fResp, err := w.sbatch.Fetch(
		ctx,
		event.JobDefinition.BatchLocationHash,
		event.CustomerAddr,
		event.JobId,
	)
	if err != nil {
		logger.I.Error("slurm fetch job body failed, setting job to failed", zap.Error(err))
		if err := w.metascheduler.SetJobStatus(ctx, event.JobId, metascheduler.JobStatusFailed, 0); err != nil {
			logger.I.Error("failed to refuse a job", zap.Error(err))
		}
		return
	}

	jobName := hexutil.Encode(event.JobId[:])
	user := strings.ToLower(event.CustomerAddr.Hex())

	definition := MapJobDefinitionToScheduler(
		event.JobDefinition,
		event.MaxDurationMinute,
		fResp.SBatch,
	)
	req := &scheduler.SubmitRequest{
		Name:          jobName,
		User:          user,
		Prefix:        "supervisor",
		JobDefinition: &definition,
	}

	// Lock the job: avoid any mutation of the job until we receive a response from sbatch
	w.resourceManager.Lock(jobName)
	defer w.resourceManager.Unlock(jobName)

	// Submit the job script
	resp, err := w.scheduler.Submit(ctx, req)
	if err != nil {
		logger.I.Error("slurm submit job failed", zap.Error(err))
		if err := w.logToUser(
			ctx,
			fResp.GridLoggerURL,
			jobName,
			user,
			[]byte(fmt.Sprintf("slurm failed to submit the job: %s", err)),
		); err != nil {
			logger.I.Error("failed to log error to the user", zap.Error(err))
		}
		if err := w.metascheduler.RefuseJob(ctx, event.JobId); err != nil {
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
	event *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
) {
	status, err := w.metascheduler.GetJobStatus(ctx, event.JobId)
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
					return w.scheduler.CancelJob(ctx, hexutil.Encode(event.JobId[:]), strings.ToLower(event.CustomerAddr.Hex()))
				}); err != nil {
				logger.I.Error("CancelJob failed, abort handleClaimNextCancellingJobEvent", zap.Error(err))
				return err
			}

			time.Sleep(5 * time.Second)

			status, err = w.metascheduler.GetJobStatus(ctx, event.JobId)
			if err != nil {
				logger.I.Error("GetJobStatus failed, considering as Cancelled", zap.Error(err))
				status = metascheduler.JobStatusCancelled
			}
			if status != metascheduler.JobStatusCancelled {
				return errors.New("failed to cancel job")
			}
			return nil
		}); err != nil {
		logger.I.Error("failed to cancel, considering as CANCELLED", zap.Error(err))
		if err := w.metascheduler.SetJobStatus(ctx, event.JobId, metascheduler.JobStatusCancelled, 0); err != nil {
			logger.I.Error("even considering as CANCELLED, it failed", zap.Error(err))
			return
		}
	}
}

func (w *Watcher) handleClaimNextTopUpEvent(
	ctx context.Context,
	event *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent,
) {
	if err := try.Do(5, 5*time.Second, func(_ int) error {
		return w.scheduler.TopUp(ctx, hexutil.Encode(event.JobId[:]), event.MaxDurationMinute)
	}); err != nil {
		logger.I.Error("failed to topup", zap.Error(err))
	}
}

func MapJobDefinitionToScheduler(
	j metaschedulerabi.JobDefinition,
	t uint64,
	body string,
) scheduler.JobDefinition {
	return scheduler.JobDefinition{
		NTasks:       j.Ntasks,
		GPUs:         &j.Gpus,
		CPUsPerTask:  j.CpusPerTask,
		TimeLimit:    t,
		MemoryPerCPU: j.MemPerCpu,
		Body:         body,
	}
}
