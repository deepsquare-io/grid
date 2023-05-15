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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	fixtureSchedulerJobID    = 1
	fixtureClaimNextJobEvent = &metascheduler.MetaSchedulerClaimJobEvent{
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
	fixtureClaimNextTopUpJobEvent = &metascheduler.MetaSchedulerClaimNextTopUpJobEvent{
		JobId:             [32]byte{1},
		ProviderAddr:      common.HexToAddress("02"),
		MaxDurationMinute: uint64(5),
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

func (suite *WatcherTestSuite) arrangeNoEvent() {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(sub, nil)
	errChan := make(chan error)
	var rErrChan <-chan error = errChan
	sub.On("Err").Return(rErrChan)
}

// arrangeEmitClaimNextJobEvent arrange a claim next job
func (suite *WatcherTestSuite) arrangeEmitClaimNextJobEvent(
	event *metascheduler.MetaSchedulerClaimJobEvent,
) {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			sink := args.Get(3).(chan<- *metascheduler.MetaSchedulerClaimJobEvent)
			go func() {
				sink <- event
			}()
		}).
		Return(sub, nil)
	errChan := make(chan error)
	var rErrChan <-chan error = errChan
	sub.On("Err").Return(rErrChan)

	// Disable slurm healthcheck
	suite.scheduler.On("HealthCheck", mock.Anything).Return(errors.New("scheduler disabled"))
}

// arrangeEmitClaimNextCancellingJobEvent arrange a claim next job
func (suite *WatcherTestSuite) arrangeEmitClaimNextCancellingJobEvent(
	event *metascheduler.MetaSchedulerClaimNextCancellingJobEvent,
) {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			sink := args.Get(2).(chan<- *metascheduler.MetaSchedulerClaimNextCancellingJobEvent)
			go func() {
				sink <- event
			}()
		}).
		Return(sub, nil)
	errChan := make(chan error)
	var rErrChan <-chan error = errChan
	sub.On("Err").Return(rErrChan)

	// Disable slurm healthcheck
	suite.scheduler.On("HealthCheck", mock.Anything).Return(errors.New("scheduler disabled"))
}

// arrangeEmitClaimNextTopUpJobEvent arrange a claim next job
func (suite *WatcherTestSuite) arrangeEmitClaimNextTopUpJobEvent(
	event *metascheduler.MetaSchedulerClaimNextTopUpJobEvent,
) {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimNextTopUpJobEvent)
			go func() {
				sink <- event
			}()
		}).
		Return(sub, nil)
	errChan := make(chan error)
	var rErrChan <-chan error = errChan
	sub.On("Err").Return(rErrChan)

	// Disable slurm healthcheck
	suite.scheduler.On("HealthCheck", mock.Anything).Return(errors.New("scheduler disabled"))
}

func (suite *WatcherTestSuite) TestClaimIndefinitely() {
	// Arrange
	suite.arrangeNoEvent()
	suite.metaQueue.On("Claim", mock.Anything).Return(nil)
	suite.metaQueue.On("ClaimCancelling", mock.Anything).Return(nil)
	suite.metaQueue.On("ClaimTopUp", mock.Anything).Return(nil)
	suite.scheduler.On("HealthCheck", mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.arrangeNoEvent()
	suite.scheduler.On("HealthCheck", mock.Anything).Return(errors.New("expected error"))

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.arrangeNoEvent()
	suite.scheduler.On("HealthCheck", mock.Anything).Return(nil)
	suite.metaQueue.On("Claim", mock.Anything).Return(nil, errors.New("expected error"))
	suite.metaQueue.On("ClaimCancelling", mock.Anything).Return(nil)
	suite.metaQueue.On("ClaimTopUp", mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.metaQueue.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	d := job.DefinitionFromMetascheduler(
		fixtureClaimNextJobEvent.JobDefinition,
		fixtureClaimNextJobEvent.MaxDurationMinute,
		fixtureBody,
	)
	suite.scheduler.On("Submit", mock.Anything, &job.SubmitRequest{
		Name:       hexutil.Encode(fixtureClaimNextJobEvent.JobId[:]),
		User:       strings.ToLower(fixtureClaimNextJobEvent.CustomerAddr.Hex()),
		Definition: &d,
	}).Return(strconv.Itoa(fixtureSchedulerJobID), nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.metaQueue.On("GetProviderAddress").Return(common.HexToAddress("123"))

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	badFixtureEvent := *fixtureClaimNextJobEvent
	badFixtureEvent.MaxDurationMinute = 0
	suite.arrangeEmitClaimNextJobEvent(&badFixtureEvent)
	suite.metaQueue.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	// Must fail job
	suite.metaQueue.On("SetJobStatus", mock.Anything, badFixtureEvent.JobId, eth.JobStatusFailed, uint64(0)).
		Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.metaQueue.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return("", errors.New("expected error"))
	// Must fail job
	suite.metaQueue.On("SetJobStatus", mock.Anything, fixtureClaimNextJobEvent.JobId, eth.JobStatusFailed, uint64(0)).
		Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.metaQueue.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	d := job.DefinitionFromMetascheduler(
		fixtureClaimNextJobEvent.JobDefinition,
		fixtureClaimNextJobEvent.MaxDurationMinute,
		fixtureBody,
	)
	suite.scheduler.On("Submit", mock.Anything, &job.SubmitRequest{
		Name:       hexutil.Encode(fixtureClaimNextJobEvent.JobId[:]),
		User:       strings.ToLower(fixtureClaimNextJobEvent.CustomerAddr.Hex()),
		Definition: &d,
	}).Return("0", errors.New("expected error"))
	// Must refuse job because we couldn't fetch the job batch script
	suite.metaQueue.On("RefuseJob", mock.Anything, mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.arrangeEmitClaimNextCancellingJobEvent(fixtureCancellingEvent)
	suite.metaQueue.On("GetJobStatus", mock.Anything, fixtureCancellingEvent.JobId).
		Return(status, nil)
	suite.metaQueue.On("GetProviderAddress").Return(fixtureCancellingEvent.ProviderAddr)
	suite.scheduler.On("CancelJob", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		status = eth.JobStatusCancelled
	}).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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
	suite.arrangeEmitClaimNextCancellingJobEvent(fixtureCancellingEvent)
	suite.metaQueue.On("GetProviderAddress").Return(common.HexToAddress("123"))

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
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

func (suite *WatcherTestSuite) TestWatchClaimNextTopUpJobEvent() {
	// Arrange
	suite.arrangeEmitClaimNextTopUpJobEvent(fixtureClaimNextTopUpJobEvent)
	suite.scheduler.On("TopUp", mock.Anything, mock.Anything).Return(nil)
	suite.metaQueue.On("GetProviderAddress").Return(fixtureClaimNextTopUpJobEvent.ProviderAddr)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.Watch(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
		time.Sleep(5 * pollingTime)
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
}

func TestWatcherTestSuite(t *testing.T) {
	suite.Run(t, &WatcherTestSuite{})
}
