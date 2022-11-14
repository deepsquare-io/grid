//go:build unit

package job_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	fixtureSchedulerJobID = uint8(1)
	fixtureEvent          = &metascheduler.MetaSchedulerClaimNextJobEvent{
		CustomerAddr:      common.HexToAddress("01"),
		JobId:             [32]byte{1},
		MaxDurationMinute: uint64(5),
		JobDefinition:     metascheduler.JobDefinition{},
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

func (suite *WatcherTestSuite) TestClaimNextJobIndefinitely() {
	// Arrange
	suite.metaQueue.On("Claim", mock.Anything).Return(nil)
	suite.scheduler.On("HealthCheck", mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.ClaimNextJobIndefinitely(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
	}

	// Assert
	suite.metaQueue.AssertExpectations(suite.T())
	suite.scheduler.AssertExpectations(suite.T())
	suite.batchFetcher.AssertExpectations(suite.T())
}

func (suite *WatcherTestSuite) TestClaimNextJobIndefinitelyWithSchedulerHealthCheckFail() {
	// Arrange
	suite.scheduler.On("HealthCheck", mock.Anything).Return(errors.New("expected error"))

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.ClaimNextJobIndefinitely(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertNotCalled(suite.T(), "Claim", mock.Anything)
	suite.batchFetcher.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestClaimNextJobIndefinitelyWithClaimFail() {
	// Arrange
	suite.scheduler.On("HealthCheck", mock.Anything).Return(nil)
	suite.metaQueue.On("Claim", mock.Anything).Return(nil, errors.New("expected error"))

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.ClaimNextJobIndefinitely(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
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
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimNextJobEvent)
		go func() {
			time.Sleep(pollingTime)
			sink <- fixtureEvent
		}()
	}).Return(sub, nil)
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	suite.scheduler.On("Submit", mock.Anything, mock.Anything).Return(int(fixtureSchedulerJobID), nil)
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertExpectations(suite.T())
}

func (suite *WatcherTestSuite) TestWatchClaimNextJobWithTimeLimitFail() {
	// Arrange
	badFixtureEvent := *fixtureEvent
	badFixtureEvent.MaxDurationMinute = 0
	sub := mocks.NewSubscription(suite.T())
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimNextJobEvent)
		go func() {
			time.Sleep(pollingTime)
			sink <- &badFixtureEvent
		}()
	}).Return(sub, nil)
	sub.On("Err").Return(nil)
	sub.On("Unsubscribe")
	// Must refuse job because we couldn't  call
	suite.metaQueue.On("RefuseJob", mock.Anything, mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
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
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimNextJobEvent)
		go func() {
			time.Sleep(pollingTime)
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
	// Must refuse job because we couldn't fetch the job batch script
	suite.metaQueue.On("RefuseJob", mock.Anything, mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
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
	suite.metaQueue.On("WatchClaimNextJobEvent", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		sink := args.Get(1).(chan<- *metascheduler.MetaSchedulerClaimNextJobEvent)
		go func() {
			time.Sleep(pollingTime)
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
	suite.scheduler.On("Submit", mock.Anything, mock.Anything).Return(0, errors.New("expected error"))
	// Must refuse job because we couldn't fetch the job batch script
	suite.metaQueue.On("RefuseJob", mock.Anything, mock.Anything).Return(nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 2*pollingTime)
	defer cancel()
	go suite.impl.WatchClaimNextJob(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
	}

	// Assert
	suite.scheduler.AssertExpectations(suite.T())
	suite.metaQueue.AssertExpectations(suite.T())
	suite.batchFetcher.AssertExpectations(suite.T())
}

func TestWatcherTestSuite(t *testing.T) {
	suite.Run(t, &WatcherTestSuite{})
}
