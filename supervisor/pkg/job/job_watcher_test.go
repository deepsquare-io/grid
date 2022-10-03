//go:build unit

package job_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	mocks "github.com/deepsquare-io/the-grid/supervisor/mocks/pkg/job"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	fixtureSlurmJobID = uint8(1)
	fixtureEvent      = &metascheduler.MetaSchedulerClaimNextJobEvent{
		CustomerAddr:      common.HexToAddress("01"),
		JobId:             [32]byte{1},
		MaxDurationMinute: big.NewInt(5),
		JobDefinition:     metascheduler.JobDefinition{},
	}
	fixtureBody = `#!/bin/sh

srun hostname
`
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
	)
}

func (suite *WatcherTestSuite) TestWatchWithClaim() {
	// Arrange
	suite.metaQueue.On("Claim", mock.Anything).Return(fixtureEvent, nil)
	suite.batchFetcher.On(
		"Fetch",
		mock.Anything,
		fixtureEvent.JobDefinition.BatchLocationHash,
	).Return(fixtureBody, nil)
	suite.scheduler.On("Submit", mock.Anything).Return(int(fixtureSlurmJobID), nil)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go suite.impl.Watch(ctx)
	select {
	case <-ctx.Done():
		logger.I.Info("test ended")
	}
}

func TestWatcherTestSuite(t *testing.T) {
	suite.Run(t, &WatcherTestSuite{})
}
