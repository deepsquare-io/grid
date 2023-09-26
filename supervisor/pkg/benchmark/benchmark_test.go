package benchmark_test

import (
	"context"
	"testing"
	"time"

	"github.com/deepsquare-io/grid/supervisor/mocks/mockscheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LauncherTestSuite struct {
	suite.Suite
	user      string
	scheduler *mockscheduler.Scheduler
	timeLimit time.Duration
	impl      benchmark.Launcher
}

func (suite *LauncherTestSuite) BeforeTest(suiteName, testName string) {
	suite.user = "root"
	suite.scheduler = mockscheduler.NewScheduler(suite.T())
	suite.timeLimit = time.Hour
	suite.impl = benchmark.NewLauncher(
		suite.user,
		"localhost:3000",
		suite.scheduler,
		benchmark.WithNoWait(),
		benchmark.WithTimeLimit(suite.timeLimit),
	)
}

func (suite *LauncherTestSuite) TestGetJobName() {
	res := suite.impl.GetJobName("test")
	suite.Require().NotEmpty(res)
}

func (suite *LauncherTestSuite) TestCancel() {
	// Arrange
	suite.scheduler.EXPECT().
		CancelJob(mock.Anything, suite.impl.GetJobName("test"), suite.user).
		Return(nil)

	// Act
	err := suite.impl.Cancel(context.Background(), "test")

	// Assert
	suite.Require().NoError(err)
}

func (suite *LauncherTestSuite) TestLaunch() {
	// Arrange
	suite.scheduler.EXPECT().
		Submit(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, sr *scheduler.SubmitRequest) (string, error) {
			suite.Assert().Equal(&scheduler.SubmitRequest{
				Name:   suite.impl.GetJobName("test"),
				User:   suite.user,
				Prefix: "benchmark",
				JobDefinition: &scheduler.JobDefinition{
					TimeLimit: uint64(suite.timeLimit.Minutes()),
					Body:      "test",
				},
			}, sr)
			return "success", nil
		})

	// Act
	err := suite.impl.Launch(context.Background(), "test", &benchmark.Benchmark{
		Body: "test",
	})

	// Assert
	suite.Require().NoError(err)
}

func TestLauncherTestSuite(t *testing.T) {
	suite.Run(t, &LauncherTestSuite{})
}
