//go:build unit

package jobapi_test

import (
	"context"
	"errors"
	"testing"
	"time"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockmetascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/jobapi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	jobHandler *mockmetascheduler.MetaScheduler
	impl       *jobapi.Server
}

func (suite *ServerTestSuite) BeforeTest(suiteName, testName string) {
	suite.jobHandler = mockmetascheduler.NewMetaScheduler(suite.T())
	suite.impl = jobapi.New(suite.jobHandler, lock.NewResourceManager())
	suite.impl.Timeout = time.Second
	suite.impl.Delay = time.Second
}

func (suite *ServerTestSuite) TestSetJobStatus() {
	ctx := context.Background()

	done := make(chan struct{})
	defer close(done)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64) error {
		done <- struct{}{}
		return nil
	})

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})

	select {
	case <-done:
	case <-time.After(5 * time.Second): // Set an appropriate timeout duration
		suite.T().Error("Timed out waiting for goroutine to complete")
	}
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func (suite *ServerTestSuite) TestSetJobStatusFailure() {
	ctx := context.Background()
	err := errors.New("problem")

	done := make(chan struct{})
	defer close(done)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64) error {
		done <- struct{}{}
		return err
	})

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})

	select {
	case <-done:
	case <-time.After(5 * time.Second): // Set an appropriate timeout duration
		suite.T().Error("Timed out waiting for goroutine to complete")
	}
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func (suite *ServerTestSuite) TestSetJobStatusBlocking() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	defer close(done)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64) error {
		<-ctx.Done()
		done <- struct{}{}
		return context.Canceled
	})

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})

	select {
	case <-done:
	case <-time.After(5 * time.Second): // Set an appropriate timeout duration
		suite.T().Error("Timed out waiting for goroutine to complete")
	}
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func (suite *ServerTestSuite) TestSetJobStatusThrowSameStatusError() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	defer close(done)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64) error {
		done <- struct{}{}
		return &metascheduler.SameStatusError{}
	})

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})
	select {
	case <-done:
	case <-time.After(5 * time.Second): // Set an appropriate timeout duration
		suite.T().Error("Timed out waiting for goroutine to complete")
	}
	suite.Require().Nil(err)
	suite.Require().NotNil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func (suite *ServerTestSuite) TestSetJobStatusThrowTransitionError() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	defer close(done)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64) error {
		done <- struct{}{}
		return &metascheduler.InvalidTransitionFromScheduled{}
	})

	// Act
	resp, err := suite.impl.SetJobStatus(ctx, &supervisorv1alpha1.SetJobStatusRequest{
		Name:     "0xb7b91cfc7853b6ec7115c5f33e092fd0",
		Id:       1,
		Duration: 1,
		Status:   supervisorv1alpha1.JobStatus_JOB_STATUS_RUNNING,
	})

	for range [2]struct{}{} {
		select {
		case <-done:
		case <-time.After(5 * time.Second): // Set an appropriate timeout duration
			suite.T().Error("Timed out waiting for goroutine to complete")
		}
	}

	suite.NoError(err)
	suite.NotNil(resp)

	suite.jobHandler.AssertExpectations(suite.T())
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, &ServerTestSuite{})
}
