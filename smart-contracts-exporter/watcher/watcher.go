package watcher

import (
	"context"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/logger"
	metricsv1 "github.com/deepsquare-io/the-grid/smart-contracts-exporter/metrics/v1"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/utils/metric"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type JobStatus uint8

const (
	JobStatusPending       JobStatus = 0
	JobStatusMetaScheduled JobStatus = 1
	JobStatusScheduled     JobStatus = 2
	JobStatusRunning       JobStatus = 3
	JobStatusCancelled     JobStatus = 4
	JobStatusFinished      JobStatus = 5
	JobStatusFailed        JobStatus = 6
	JobStatusOutOfCredits  JobStatus = 7
	JobStatusUnknown       JobStatus = 255
)

var (
	jobTransitionEvent abi.Event
	newJobRequestEvent abi.Event
	jobRefusedEvent    abi.Event
	billedTooMuchEvent abi.Event
)

func init() {
	contractABI, err := abi.JSON(strings.NewReader(metascheduler.MetaSchedulerABI))
	if err != nil {
		logger.I.Fatal("failed to parse contract ABI", zap.Error(err))
	}

	// Find the event signature dynamically
	var ok bool
	jobTransitionEvent, ok = contractABI.Events["JobTransitionEvent"]
	if !ok {
		logger.I.Fatal("failed to parse contract ABI", zap.Error(err))
	}

	newJobRequestEvent, ok = contractABI.Events["NewJobRequestEvent"]
	if !ok {
		logger.I.Fatal("failed to parse contract ABI", zap.Error(err))
	}

	jobRefusedEvent, ok = contractABI.Events["JobRefusedEvent"]
	if !ok {
		logger.I.Fatal("failed to parse contract ABI", zap.Error(err))
	}

	billedTooMuchEvent, ok = contractABI.Events["BilledTooMuchEvent"]
	if !ok {
		logger.I.Fatal("failed to parse contract ABI", zap.Error(err))
	}
}

type Watcher struct {
	clientRPC            *ethclient.Client
	clientWS             *ethclient.Client
	contractRPC          *metascheduler.MetaScheduler
	contractWS           *metascheduler.MetaScheduler
	metaschedulerAddress common.Address
}

func New(
	clientRPC *ethclient.Client,
	clientWS *ethclient.Client,
	contractRPC *metascheduler.MetaScheduler,
	contractWS *metascheduler.MetaScheduler,
	metaschedulerAddress common.Address,
) *Watcher {
	return &Watcher{
		clientRPC:            clientRPC,
		clientWS:             clientWS,
		contractRPC:          contractRPC,
		contractWS:           contractWS,
		metaschedulerAddress: metaschedulerAddress,
	}
}

func (w *Watcher) WatchNewEvents(ctx context.Context, readyChan chan<- struct{}) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Get Last block
	header, err := w.clientRPC.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	logger.I.Info("found head block", zap.String("head block", header.Number.String()))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{w.metaschedulerAddress},
		Topics: [][]common.Hash{
			{
				jobTransitionEvent.ID,
				newJobRequestEvent.ID,
				jobRefusedEvent.ID,
			},
		},
	}

	lastBlockWatched := metric.GetGaugeValue(metricsv1.LastBlockWatched)
	if lastBlockWatched != 0 {
		logger.I.Info("found last block", zap.Float64("last block watched", lastBlockWatched))
		iLastBlockWatched := math.Round(lastBlockWatched)
		bLastBlockWatched := big.NewInt(int64(iLastBlockWatched) + 1)
		query.FromBlock = bLastBlockWatched
	}

	if query.FromBlock == nil || header.Number.Cmp(query.FromBlock) > 0 {
		logger.I.Info("retrieving old blocks...")

		// Retrieve old blocks
		oldLogs, err := w.clientRPC.FilterLogs(ctx, query)
		if err != nil {
			logger.I.Error("failed to read old logs", zap.Error(err))
			return err
		}
		logger.I.Info("retrieved old blocks", zap.Int("n", len(oldLogs)))
		for _, log := range oldLogs {
			if err := w.handleLog(ctx, log); err != nil {
				return err
			}
		}

		if len(oldLogs) > 0 {
			metricsv1.LastBlockWatched.Set(float64(oldLogs[len(oldLogs)-1].BlockNumber))
			if err := metricsv1.Save(); err != nil {
				logger.I.Warn("failed to checkpoint", zap.Error(err))
			} else {
				logger.I.Debug("checkpoint")
			}
		}

		logger.I.Info("finished replaying old blocks, now watching...")
	} else {
		logger.I.Info("at latest block, no need to retrieve old blocks, now watching...")
	}

	readyChan <- struct{}{}

	// Retrieve new blocks
	logs := make(chan types.Log)

	sub, err := w.clientWS.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		logger.I.Error("failed to subscribe", zap.Error(err))
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case log, ok := <-logs:
			if !ok {
				return nil
			}
			if err := w.handleLog(ctx, log); err != nil {
				return err
			}
		case <-ticker.C:
			if err := metricsv1.Save(); err != nil {
				logger.I.Warn("failed to checkpoint", zap.Error(err))
			} else {
				logger.I.Debug("checkpoint")
			}
		}
	}
}

func (w *Watcher) handleLog(ctx context.Context, log types.Log) error {
	switch log.Topics[0].Hex() {
	case jobTransitionEvent.ID.Hex():
		event, err := w.contractRPC.ParseJobTransitionEvent(log)
		if err != nil {
			logger.I.Panic("failed to parse event", zap.Error(err))
		}
		if err := w.handleJobTransition(ctx, event); err != nil {
			return err
		}

	case newJobRequestEvent.ID.Hex():
		event, err := w.contractRPC.ParseNewJobRequestEvent(log)
		if err != nil {
			logger.I.Panic("failed to parse event", zap.Error(err))
		}
		if err := w.handleNewJobRequest(ctx, event); err != nil {
			return err
		}

	case jobRefusedEvent.ID.Hex():
		event, err := w.contractRPC.ParseJobRefusedEvent(log)
		if err != nil {
			logger.I.Panic("failed to parse event", zap.Error(err))
		}
		if err := w.handleJobRefused(ctx, event); err != nil {
			return err
		}

	case billedTooMuchEvent.ID.Hex():
		event, err := w.contractRPC.ParseBilledTooMuchEvent(log)
		if err != nil {
			logger.I.Panic("failed to parse event", zap.Error(err))
		}
		if err := w.handleBilledTooMuch(ctx, event); err != nil {
			return err
		}

	default:
		logger.I.Warn("unknown log", zap.Any("log", log))
	}
	metricsv1.LastBlockWatched.Set(float64(log.BlockNumber))
	return nil
}

func (w *Watcher) handleJobTransition(ctx context.Context, event *metascheduler.MetaSchedulerJobTransitionEvent) error {
	from := JobStatus(event.From)
	to := JobStatus(event.To)
	job, err := w.contractRPC.Jobs(&bind.CallOpts{
		Context: ctx,
	}, event.JobId)
	if err != nil {
		return err
	}

	// Move a state from
	switch from {
	case JobStatusPending:
		metricsv1.TotalJobsPending(job.CustomerAddr.Hex()).Dec()
	case JobStatusMetaScheduled:
		metricsv1.TotalJobsMetaScheduled(job.CustomerAddr.Hex()).Dec()
	case JobStatusScheduled:
		metricsv1.TotalJobsScheduled(job.CustomerAddr.Hex()).Dec()
	case JobStatusRunning:
		metricsv1.TotalJobsRunning(job.CustomerAddr.Hex()).Dec()
	case JobStatusCancelled:
		metricsv1.TotalJobsCancelled(job.CustomerAddr.Hex()).Dec()
	case JobStatusFinished:
		metricsv1.TotalJobsFinished(job.CustomerAddr.Hex()).Dec()
	case JobStatusFailed:
		metricsv1.TotalJobsFailed(job.CustomerAddr.Hex()).Dec()
	case JobStatusOutOfCredits:
		metricsv1.TotalJobsOutOfCredits(job.CustomerAddr.Hex()).Dec()
	}

	// Move a state to
	switch to {
	case JobStatusPending:
		metricsv1.TotalJobsPending(job.CustomerAddr.Hex()).Inc()
	case JobStatusMetaScheduled:
		metricsv1.TotalJobsMetaScheduled(job.CustomerAddr.Hex()).Inc()
	case JobStatusScheduled:
		metricsv1.TotalJobsScheduled(job.CustomerAddr.Hex()).Inc()
	case JobStatusRunning:
		metricsv1.TotalJobsRunning(job.CustomerAddr.Hex()).Inc()
	case JobStatusCancelled:
		metricsv1.TotalJobsCancelled(job.CustomerAddr.Hex()).Inc()
	case JobStatusFinished:
		metricsv1.TotalJobsFinished(job.CustomerAddr.Hex()).Inc()
	case JobStatusFailed:
		metricsv1.TotalJobsFailed(job.CustomerAddr.Hex()).Inc()
	case JobStatusOutOfCredits:
		metricsv1.TotalJobsOutOfCredits(job.CustomerAddr.Hex()).Inc()
	}

	// Add statistics from cold states
	switch to {
	case JobStatusCancelled:
		fallthrough
	case JobStatusFinished:
		fallthrough
	case JobStatusFailed:
		fallthrough
	case JobStatusOutOfCredits:
		bf := new(big.Float).SetInt(job.Cost.FinalCost)
		f, _ := bf.Float64()
		metricsv1.TotalCreditSpent(job.CustomerAddr.Hex()).Add(f)

		bduration := new(big.Int).Div(new(big.Int).Sub(job.Time.End, job.Time.Start), big.NewInt(60))
		if bduration.Sign() == -1 {
			logger.I.Error("job duration is negative", zap.String("duration", bduration.String()), zap.Any("job", job))
			return nil
		}
		f, _ = new(big.Float).SetInt(bduration).Float64()
		metricsv1.TotalJobDuration(job.CustomerAddr.Hex()).Add(f)

		cpus := new(big.Int).SetUint64(job.Definition.CpuPerTask * job.Definition.Ntasks)
		cpuTime := new(big.Int).Mul(bduration, cpus)
		f, _ = new(big.Float).SetInt(cpuTime).Float64()
		metricsv1.TotalCPUTime(job.CustomerAddr.Hex()).Add(f)

		gpus := new(big.Int).SetUint64(job.Definition.GpuPerTask * job.Definition.Ntasks)
		gpuTime := new(big.Int).Mul(bduration, gpus)
		f, _ = new(big.Float).SetInt(gpuTime).Float64()
		metricsv1.TotalGPUTime(job.CustomerAddr.Hex()).Add(f)
	}

	return nil
}

func (w *Watcher) handleNewJobRequest(ctx context.Context, event *metascheduler.MetaSchedulerNewJobRequestEvent) error {
	metricsv1.TotalJobsPending(event.CustomerAddr.Hex()).Inc()
	metricsv1.TotalNumberOfJobs(event.CustomerAddr.Hex()).Inc()
	job, err := w.contractRPC.Jobs(&bind.CallOpts{
		Context: ctx,
	}, event.JobId)
	if err != nil {
		return err
	}
	metricsv1.AddJob(&job)

	return nil
}

func (w *Watcher) handleJobRefused(ctx context.Context, event *metascheduler.MetaSchedulerJobRefusedEvent) error {
	job, err := w.contractRPC.Jobs(&bind.CallOpts{
		Context: ctx,
	}, event.JobId)
	if err != nil {
		return err
	}
	metricsv1.TotalNumberOfJobs(job.CustomerAddr.Hex()).Dec()
	metricsv1.TotalJobsPending(job.CustomerAddr.Hex()).Dec()
	metricsv1.TotalJobRefused(event.ProviderAddr.Hex()).Inc()
	return nil
}

func (w *Watcher) handleBilledTooMuch(ctx context.Context, event *metascheduler.MetaSchedulerBilledTooMuchEvent) error {
	job, err := w.contractRPC.Jobs(&bind.CallOpts{
		Context: ctx,
	}, event.JobId)
	if err != nil {
		return err
	}

	metricsv1.TotalBilledTooMuch(job.ProviderAddr.Hex()).Inc()
	return nil
}