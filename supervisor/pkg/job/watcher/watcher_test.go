//go:build unit

package watcher_test

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/watcher"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	fixtureSchedulerJobID    = 1
	fixtureClaimNextJobEvent = &metaschedulerabi.MetaSchedulerClaimJobEvent{
		CustomerAddr:      common.HexToAddress("01"),
		ProviderAddr:      common.HexToAddress("02"),
		JobId:             [32]byte{1},
		MaxDurationMinute: uint64(5),
		JobDefinition:     metaschedulerabi.JobDefinition{},
	}
	fixtureCancellingEvent = &metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent{
		CustomerAddr: common.HexToAddress("01"),
		ProviderAddr: common.HexToAddress("02"),
		JobId:        [32]byte{1},
	}
	fixtureClaimNextTopUpJobEvent = &metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent{
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
	metaScheduler *mocks.MetaScheduler
	scheduler     *mocks.Scheduler
	sbatch        *mocks.Client
	impl          *watcher.Watcher
}

func (suite *WatcherTestSuite) BeforeTest(suiteName, testName string) {
	suite.metaScheduler = mocks.NewMetaScheduler(suite.T())
	suite.scheduler = mocks.NewScheduler(suite.T())
	suite.sbatch = mocks.NewClient(suite.T())
	suite.impl = watcher.New(
		suite.metaScheduler,
		suite.scheduler,
		suite.sbatch,
		pollingTime,
	)
}

func (suite *WatcherTestSuite) arrangeNoEvent() {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaScheduler.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(sub, nil)
	errChan := make(chan error)
	var rErrChan <-chan error = errChan
	sub.On("Err").Return(rErrChan)
}

// arrangeEmitClaimNextJobEvent arrange a claim next job
func (suite *WatcherTestSuite) arrangeEmitClaimNextJobEvent(
	event *metaschedulerabi.MetaSchedulerClaimJobEvent,
) {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaScheduler.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			sink := args.Get(3).(chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent)
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
	event *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
) {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaScheduler.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			sink := args.Get(2).(chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent)
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
	event *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent,
) {
	// Arrange
	sub := mocks.NewSubscription(suite.T())
	suite.metaScheduler.On("WatchEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			sink := args.Get(1).(chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent)
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
	suite.metaScheduler.On("Claim", mock.Anything).Return(nil)
	suite.metaScheduler.On("ClaimCancelling", mock.Anything).Return(nil)
	suite.metaScheduler.On("ClaimTopUp", mock.Anything).Return(nil)
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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.scheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertExpectations(suite.T())
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
	suite.metaScheduler.AssertNotCalled(suite.T(), "Claim", mock.Anything)
	suite.metaScheduler.AssertNotCalled(suite.T(), "ClaimCancelling", mock.Anything)
	suite.sbatch.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestClaimIndefinitelyWithClaimFail() {
	// Arrange
	suite.arrangeNoEvent()
	suite.scheduler.On("HealthCheck", mock.Anything).Return(nil)
	suite.metaScheduler.On("Claim", mock.Anything).Return(nil, errors.New("expected error"))
	suite.metaScheduler.On("ClaimCancelling", mock.Anything).Return(nil)
	suite.metaScheduler.On("ClaimTopUp", mock.Anything).Return(nil)

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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextJob() {
	// Arrange
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.metaScheduler.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.sbatch.On(
		"Fetch",
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	d := watcher.MapJobDefinitionToScheduler(
		fixtureClaimNextJobEvent.JobDefinition,
		fixtureClaimNextJobEvent.MaxDurationMinute,
		fixtureBody,
	)
	suite.scheduler.On("Submit", mock.Anything, &scheduler.SubmitRequest{
		Name:          hexutil.Encode(fixtureClaimNextJobEvent.JobId[:]),
		User:          strings.ToLower(fixtureClaimNextJobEvent.CustomerAddr.Hex()),
		JobDefinition: &d,
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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertExpectations(suite.T())
}

func (suite *WatcherTestSuite) TestWatchClaimNextJobIgnoresEvent() {
	// Arrange
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.metaScheduler.On("GetProviderAddress").Return(common.HexToAddress("123"))

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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextJobWithTimeLimitFail() {
	// Arrange
	badFixtureEvent := *fixtureClaimNextJobEvent
	badFixtureEvent.MaxDurationMinute = 0
	suite.arrangeEmitClaimNextJobEvent(&badFixtureEvent)
	suite.metaScheduler.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	// Must fail job
	suite.metaScheduler.On("SetJobStatus", mock.Anything, badFixtureEvent.JobId, metascheduler.JobStatusFailed, uint64(0)).
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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextJobWithBatchFetchFail() {
	// Arrange
	suite.metaScheduler.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.sbatch.On(
		"Fetch",
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return("", errors.New("expected error"))
	// Must fail job
	suite.metaScheduler.On("SetJobStatus", mock.Anything, fixtureClaimNextJobEvent.JobId, metascheduler.JobStatusFailed, uint64(0)).
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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertExpectations(suite.T())
	suite.scheduler.AssertNotCalled(suite.T(), "Submit", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchWithSchedulerSubmitFail() {
	// Arrange
	suite.metaScheduler.On("GetProviderAddress").Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.sbatch.On(
		"Fetch",
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	d := watcher.MapJobDefinitionToScheduler(
		fixtureClaimNextJobEvent.JobDefinition,
		fixtureClaimNextJobEvent.MaxDurationMinute,
		fixtureBody,
	)
	suite.scheduler.On("Submit", mock.Anything, &scheduler.SubmitRequest{
		Name:          hexutil.Encode(fixtureClaimNextJobEvent.JobId[:]),
		User:          strings.ToLower(fixtureClaimNextJobEvent.CustomerAddr.Hex()),
		JobDefinition: &d,
	}).Return("0", errors.New("expected error"))
	// Must refuse job because we couldn't fetch the job batch script
	suite.metaScheduler.On("RefuseJob", mock.Anything, mock.Anything).Return(nil)

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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertExpectations(suite.T())
}

func (suite *WatcherTestSuite) TestWatchClaimNextCancellingJobEvent() {
	// Arrange
	status := metascheduler.JobStatusRunning
	suite.arrangeEmitClaimNextCancellingJobEvent(fixtureCancellingEvent)
	suite.metaScheduler.On("GetJobStatus", mock.Anything, fixtureCancellingEvent.JobId).
		Return(status, nil)
	suite.metaScheduler.On("GetProviderAddress").Return(fixtureCancellingEvent.ProviderAddr)
	suite.scheduler.On("CancelJob", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		status = metascheduler.JobStatusCancelled
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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextCancellingJobEventIgnoresEvent() {
	// Arrange
	suite.arrangeEmitClaimNextCancellingJobEvent(fixtureCancellingEvent)
	suite.metaScheduler.On("GetProviderAddress").Return(common.HexToAddress("123"))

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
	suite.metaScheduler.AssertExpectations(suite.T())
	suite.sbatch.AssertNotCalled(suite.T(), "Fetch", mock.Anything, mock.Anything)
}

func (suite *WatcherTestSuite) TestWatchClaimNextTopUpJobEvent() {
	// Arrange
	suite.arrangeEmitClaimNextTopUpJobEvent(fixtureClaimNextTopUpJobEvent)
	suite.scheduler.On("TopUp", mock.Anything, mock.Anything).Return(nil)
	suite.metaScheduler.On("GetProviderAddress").Return(fixtureClaimNextTopUpJobEvent.ProviderAddr)

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
	suite.metaScheduler.AssertExpectations(suite.T())
}

func TestWatcherTestSuite(t *testing.T) {
	suite.Run(t, &WatcherTestSuite{})
}
