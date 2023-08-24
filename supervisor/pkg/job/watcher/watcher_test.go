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
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockethereum"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockgridlogger"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockloggerv1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockmetascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mocksbatch"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockscheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/watcher"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/sbatch"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/event"
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
	metaScheduler    *mockmetascheduler.MetaScheduler
	scheduler        *mockscheduler.Scheduler
	sbatch           *mocksbatch.Client
	gridLoggerDialer *mockgridlogger.Dialer
	impl             *watcher.Watcher
}

func (suite *WatcherTestSuite) BeforeTest(suiteName, testName string) {
	suite.metaScheduler = mockmetascheduler.NewMetaScheduler(suite.T())
	suite.scheduler = mockscheduler.NewScheduler(suite.T())
	suite.sbatch = mocksbatch.NewClient(suite.T())
	suite.gridLoggerDialer = mockgridlogger.NewDialer(suite.T())
	suite.impl = watcher.New(
		suite.metaScheduler,
		suite.scheduler,
		suite.sbatch,
		pollingTime,
		lock.NewResourceManager(),
		suite.gridLoggerDialer,
	)
}

func (suite *WatcherTestSuite) expectLoggerSend() *mockloggerv1alpha1.LoggerAPI_WriteClient {
	c := mockloggerv1alpha1.NewLoggerAPI_WriteClient(suite.T())
	suite.gridLoggerDialer.EXPECT().
		DialContext(mock.Anything, mock.Anything).
		Return(c, func() error { return nil }, nil)
	c.EXPECT().Send(mock.Anything).Return(nil)
	c.EXPECT().CloseAndRecv().Return(nil, nil)
	return c
}

func (suite *WatcherTestSuite) arrangeNoEvent() {
	// Arrange
	sub := mockethereum.NewSubscription(suite.T())
	suite.metaScheduler.EXPECT().WatchEvents(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(sub, nil)
	errChan := make(chan error, 1)
	var rErrChan <-chan error = errChan
	sub.EXPECT().Err().Return(rErrChan)
}

// arrangeEmitClaimNextJobEvent arrange a claim next job
func (suite *WatcherTestSuite) arrangeEmitClaimNextJobEvent(
	e *metaschedulerabi.MetaSchedulerClaimJobEvent,
) {
	// Arrange
	sub := mockethereum.NewSubscription(suite.T())
	suite.metaScheduler.EXPECT().WatchEvents(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).RunAndReturn(func(ctx context.Context, c1 chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, c2 chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, c3 chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) (event.Subscription, error) {
		go func() {
			c3 <- e
		}()
		return sub, nil
	})
	errChan := make(chan error, 1)
	var rErrChan <-chan error = errChan
	sub.EXPECT().Err().Return(rErrChan)

	// Disable slurm healthcheck
	suite.scheduler.EXPECT().HealthCheck(mock.Anything).Return(errors.New("scheduler disabled"))
}

// arrangeEmitClaimNextCancellingJobEvent arrange a claim next job
func (suite *WatcherTestSuite) arrangeEmitClaimNextCancellingJobEvent(
	e *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
) {
	// Arrange
	sub := mockethereum.NewSubscription(suite.T())
	suite.metaScheduler.EXPECT().
		WatchEvents(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, c1 chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, c2 chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, c3 chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) (event.Subscription, error) {
			go func() {
				c2 <- e
			}()
			return sub, nil
		})
	errChan := make(chan error, 1)
	var rErrChan <-chan error = errChan
	sub.EXPECT().Err().Return(rErrChan)

	// Disable slurm healthcheck
	suite.scheduler.EXPECT().HealthCheck(mock.Anything).Return(errors.New("scheduler disabled"))
}

// arrangeEmitClaimNextTopUpJobEvent arrange a claim next job
func (suite *WatcherTestSuite) arrangeEmitClaimNextTopUpJobEvent(
	e *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent,
) {
	// Arrange
	sub := mockethereum.NewSubscription(suite.T())
	suite.metaScheduler.EXPECT().
		WatchEvents(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, c1 chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, c2 chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, c3 chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) (event.Subscription, error) {
			go func() {
				c1 <- e
			}()
			return sub, nil
		})
	errChan := make(chan error, 1)
	var rErrChan <-chan error = errChan
	sub.EXPECT().Err().Return(rErrChan)

	// Disable slurm healthcheck
	suite.scheduler.EXPECT().HealthCheck(mock.Anything).Return(errors.New("scheduler disabled"))
}

func (suite *WatcherTestSuite) TestClaimIndefinitely() {
	// Arrange
	suite.arrangeNoEvent()
	suite.metaScheduler.EXPECT().Claim(mock.Anything).Return(nil)
	suite.metaScheduler.EXPECT().ClaimCancelling(mock.Anything).Return(nil)
	suite.metaScheduler.EXPECT().ClaimTopUp(mock.Anything).Return(nil)
	suite.scheduler.EXPECT().HealthCheck(mock.Anything).Return(nil)

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
	suite.scheduler.EXPECT().HealthCheck(mock.Anything).Return(errors.New("expected error"))

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
	suite.scheduler.EXPECT().HealthCheck(mock.Anything).Return(nil)
	suite.metaScheduler.EXPECT().Claim(mock.Anything).Return(errors.New("expected error"))

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
	suite.metaScheduler.EXPECT().GetProviderAddress().Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.sbatch.EXPECT().
		Fetch(
			mock.Anything,
			fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
		).Return(sbatch.FetchResponse{
		SBatch: fixtureBody,
	}, nil)
	d := watcher.MapJobDefinitionToScheduler(
		fixtureClaimNextJobEvent.JobDefinition,
		fixtureClaimNextJobEvent.MaxDurationMinute,
		fixtureBody,
	)
	suite.scheduler.EXPECT().Submit(mock.Anything, &scheduler.SubmitRequest{
		Name:          hexutil.Encode(fixtureClaimNextJobEvent.JobId[:]),
		User:          strings.ToLower(fixtureClaimNextJobEvent.CustomerAddr.Hex()),
		Prefix:        "supervisor",
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
	suite.metaScheduler.EXPECT().GetProviderAddress().Return(common.HexToAddress("123"))

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
	suite.metaScheduler.EXPECT().GetProviderAddress().Return(fixtureClaimNextJobEvent.ProviderAddr)
	// Must fail job
	suite.metaScheduler.EXPECT().
		SetJobStatus(mock.Anything, badFixtureEvent.JobId, metascheduler.JobStatusFailed, uint64(0)).
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
	suite.metaScheduler.EXPECT().GetProviderAddress().Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.sbatch.EXPECT().Fetch(
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return(sbatch.FetchResponse{}, errors.New("expected error"))
	// Must fail job
	suite.metaScheduler.EXPECT().
		SetJobStatus(mock.Anything, fixtureClaimNextJobEvent.JobId, metascheduler.JobStatusFailed, uint64(0)).
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
	suite.expectLoggerSend()
	suite.metaScheduler.EXPECT().GetProviderAddress().Return(fixtureClaimNextJobEvent.ProviderAddr)
	suite.arrangeEmitClaimNextJobEvent(fixtureClaimNextJobEvent)
	suite.sbatch.EXPECT().Fetch(
		mock.Anything,
		fixtureClaimNextJobEvent.JobDefinition.BatchLocationHash,
	).Return(sbatch.FetchResponse{
		SBatch: fixtureBody,
	}, nil)
	d := watcher.MapJobDefinitionToScheduler(
		fixtureClaimNextJobEvent.JobDefinition,
		fixtureClaimNextJobEvent.MaxDurationMinute,
		fixtureBody,
	)
	suite.scheduler.EXPECT().Submit(mock.Anything, &scheduler.SubmitRequest{
		Name:          hexutil.Encode(fixtureClaimNextJobEvent.JobId[:]),
		User:          strings.ToLower(fixtureClaimNextJobEvent.CustomerAddr.Hex()),
		Prefix:        "supervisor",
		JobDefinition: &d,
	}).Return("0", errors.New("expected error"))
	// Must refuse job because we couldn't fetch the job batch script
	suite.metaScheduler.EXPECT().RefuseJob(mock.Anything, mock.Anything).Return(nil)

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
	suite.metaScheduler.EXPECT().GetJobStatus(mock.Anything, fixtureCancellingEvent.JobId).
		Return(status, nil)
	suite.metaScheduler.EXPECT().GetProviderAddress().Return(fixtureCancellingEvent.ProviderAddr)
	suite.scheduler.EXPECT().
		CancelJob(mock.Anything, mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, s1, s2 string) error {
			status = metascheduler.JobStatusCancelled
			return nil
		})

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
	suite.metaScheduler.EXPECT().GetProviderAddress().Return(common.HexToAddress("123"))

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
	suite.scheduler.EXPECT().TopUp(mock.Anything, mock.Anything, mock.Anything).Return(nil)
	suite.metaScheduler.EXPECT().
		GetProviderAddress().
		Return(fixtureClaimNextTopUpJobEvent.ProviderAddr)

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
