//go:build unit

package jobapi_test

import (
	"context"
	"errors"
	"testing"
	"time"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/jobapi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JobHandlerTestSuite struct {
	suite.Suite
	jobHandler *mocks.JobHandler
	impl       *jobapi.JobAPIServer
}

func (suite *JobHandlerTestSuite) BeforeTest(suiteName, testName string) {
	suite.jobHandler = mocks.NewJobHandler(suite.T())
	suite.impl = jobapi.New(suite.jobHandler)
	suite.impl.Timeout = time.Second
	suite.impl.Delay = time.Second
}

func (suite *JobHandlerTestSuite) TestSetJobStatus() {
	ctx := context.Background()

	suite.jobHandler.On(
		"SetJobStatus",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func (suite *JobHandlerTestSuite) TestSetJobStatusFailure() {
	ctx := context.Background()
	err := errors.New("problem")

	suite.jobHandler.On(
		"SetJobStatus",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(err)

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})
	suite.Require().Error(err)
	suite.Require().Nil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func (suite *JobHandlerTestSuite) TestSetJobStatusBlocking() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	suite.jobHandler.On(
		"SetJobStatus",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(func(args mock.Arguments) {
		<-args[0].(context.Context).Done()
	}).Return(context.Canceled)

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})
	suite.Require().Error(err)
	suite.Require().Nil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func TestJobHandlerTestSuite(t *testing.T) {
	suite.Run(t, &JobHandlerTestSuite{})
}
