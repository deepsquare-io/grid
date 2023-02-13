//go:build unit

package job_test

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	fixtureSchedulerJobID = 1
	fixtureEvent          = &metascheduler.MetaSchedulerClaimJobEvent{
		CustomerAddr:      common.HexToAddress("01"),
		ProviderAddr:      common.HexToAddress("02"),
		JobId:             [32]byte{1},
		MaxDurationMinute: uint64(5),
		JobDefinition:     metascheduler.JobDefinition{},
	}
	fixtureCancellingEvent = &metascheduler.MetaSchedulerClaimNextCancellingJobEvent{
		CustomerAddr: common.HexToAddress("01"),
		ProviderAddr: common.HexToAddress("02"),
		JobId:        [32]byte{1},
	}
	fixtureBody = `#!/bin/sh

srun hostname
`
	pollingTime = time.Duration(100 * time.Millisecond)
)

type WatcherTestSuite struct {
	suite.Suite
	metaQueue    *mocks.JobMetaQueue
	scheduler    *mocks.JobScheduler
	batchFetcher *mocks.JobBatchFetcher
	impl         *job.Watcher
}

func (suite *WatcherTestSuite) BeforeTest(suiteName, testName string) {
	suite.metaQueue = mocks.NewJobMetaQueue(suite.T())
	suite.scheduler = mocks.NewJobScheduler(suite.T())
	suite.batchFetcher = mocks.NewJobBatchFetcher(suite.T())
	suite.impl = job.New(
		suite.metaQueue,
		suite.scheduler,
		suite.batchFetcher,
		pollingTime,
	)
}

func (suite *WatcherTestSuite) TestClaimIndefinitely() {
	// Arrange
	suite.metaQueue.On("Claim", mock.Anything).Return(nil)
	suite.metaQueue.On("ClaimCancelling", mock.Anything).Return(nil)
	suite.scheduler.On("HealthCheck", mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.ClaimIndefinitely(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.metaQueue.AssertExpectations(suite.T())
	suite.scheduler.AssertExpectations(suite.T())
	suite.batchFetcher.AssertExpectations(suite.T())
}

func (suite *WatcherTestSuite) TestClaimIndefinitelyWithSchedulerHealthCheckFail() {
	// Arrange
	suite.scheduler.On("HealthCheck", mock.Anything).Return(errors.New("expected error"))

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.ClaimIndefinitely(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertNotCalled(suite.T(), "Claim", mock.Anything)
	suite.metaQueue.AssertNotCalled(suite.T(), "ClaimCancelling", mock.Anything)
	suite.batchFetcher.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestClaimIndefinitelyWithClaimFail() {
	// Arrange
	suite.scheduler.On("HealthCheck", mock.Anything).Return(nil)
	suite.metaQueue.On("Claim", mock.Anything).Return(nil, errors.New("expected error"))
	suite.metaQueue.On("ClaimCancelling", mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.ClaimIndefinitely(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextJob() {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimJobEvent)
		go func() {
			sink <- fixtureEvent
		}()
	}).Return(sub, nil)
	suite.metaQueue.On("GetProviderAddress").Return(fixtureEvent.ProviderAddr)
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	job := eth.JobDefinitionMapToSlurm(fixtureEvent.JobDefinition, fixtureEvent.MaxDurationMinute, fixtureBody)
	suite.scheduler.On("Submit", mock.Anything, &slurm.SubmitJobRequest{
		Name:          hexutil.Encode(fixtureEvent.JobId[:]),
		User:          strings.ToLower(fixtureEvent.CustomerAddr.Hex()),
		JobDefinition: &job,
	}).Return(strconv.Itoa(fixtureSchedulerJobID), nil)
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertExpectations(suite.T())
}

func (suite *WatcherTestSuite) TestWatchClaimNextJobIgnoresEvent() {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimJobEvent)
		go func() {
			sink <- fixtureEvent
		}()
	}).Return(sub, nil)
	suite.metaQueue.On("GetProviderAddress").Return(common.HexToAddress("123"))
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextJobWithTimeLimitFail() {
	// Arrange
	badFixtureEvent := *fixtureEvent
	badFixtureEvent.MaxDurationMinute = 0
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("GetProviderAddress").Return(fixtureEvent.ProviderAddr)
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimJobEvent)
		go func() {
			sink <- &badFixtureEvent
		}()
	}).Return(sub, nil)
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")
	// Must fail job
	suite.metaQueue.On("SetJobStatus", mock.Anything, badFixtureEvent.JobId, eth.JobStatusFailed, uint64(0)).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextJobWithBatchFetchFail() {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("GetProviderAddress").Return(fixtureEvent.ProviderAddr)
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimJobEvent)
		go func() {
			sink <- fixtureEvent
		}()
	}).Return(sub, nil)
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureEvent.JobDefinition.BatchLocationHash,
	).Return("", errors.New("expected error"))
	// Must fail job
	suite.metaQueue.On("SetJobStatus", mock.Anything, fixtureEvent.JobId, eth.JobStatusFailed, uint64(0)).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertExpectations(suite.T())
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchWithSchedulerSubmitFail() {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("GetProviderAddress").Return(fixtureEvent.ProviderAddr)
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimJobEvent)
		go func() {
			sink <- fixtureEvent
		}()
	}).Return(sub, nil)
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	job := eth.JobDefinitionMapToSlurm(fixtureEvent.JobDefinition, fixtureEvent.MaxDurationMinute, fixtureBody)
	suite.scheduler.On("Submit", mock.Anything, &slurm.SubmitJobRequest{
		Name:          hexutil.Encode(fixtureEvent.JobId[:]),
		User:          strings.ToLower(fixtureEvent.CustomerAddr.Hex()),
		JobDefinition: &job,
	}).Return("0", errors.New("expected error"))
	// Must refuse job because we couldn't fetch the job batch script
	suite.metaQueue.On("RefuseJob", mock.Anything, mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertExpectations(suite.T())
}

func (suite *WatcherTestSuite) TestWatchClaimNextCancellingJobEvent() {
	// Arrange
	status := eth.JobStatusRunning
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchClaimNextCancellingJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimNextCancellingJobEvent)
		go func() {
			sink <- fixtureCancellingEvent
		}()
	}).Return(sub, nil)
	suite.metaQueue.On("GetJobStatus", mock.Anything, fixtureCancellingEvent.JobId).Return(status, nil)
	suite.metaQueue.On("GetProviderAddress").Return(fixtureCancellingEvent.ProviderAddr)
	suite.scheduler.On("CancelJob", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		status = eth.JobStatusCancelled
	}).Return(nil)
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextCancellingJobEvent(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextCancellingJobEventIgnoresEvent() {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchClaimNextCancellingJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimNextCancellingJobEvent)
		go func() {
			sink <- fixtureCancellingEvent
		}()
	}).Return(sub, nil)
	suite.metaQueue.On("GetProviderAddress").Return(common.HexToAddress("123"))
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextCancellingJobEvent(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertNotCalled(suite.T(), "GetJobStatus", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "CancelJob", mock.Anything, mock.Anything)
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
}

func TestWatcherTestSuite(t *testing.T) {
	suite.Run(t, &WatcherTestSuite{})
}
