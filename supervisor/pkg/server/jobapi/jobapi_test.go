// Copyright (C) 2023 DeepSquare Asociation
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

//go:build unit

package jobapi_test

import (
	"context"
	"errors"
	"testing"
	"time"

	supervisorv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/supervisor/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/mocks/mockmetascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/lock"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/server/jobapi"
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

	done := make(chan struct{}, 1)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64, sjso ...metascheduler.SetJobStatusOption) error {
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
		suite.T().Fatal("Timed out waiting for goroutine to complete")
	}
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)
}

func (suite *ServerTestSuite) TestSetJobStatusFailure() {
	ctx := context.Background()

	done := make(chan struct{}, 1)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64, sjso ...metascheduler.SetJobStatusOption) error {
		done <- struct{}{}
		return errors.New("problem")
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
		suite.T().Fatal("Timed out waiting for goroutine to complete")
	}
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)
}

func (suite *ServerTestSuite) TestSetJobStatusBlocking() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{}, 1)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64, sjso ...metascheduler.SetJobStatusOption) error {
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
		suite.T().Fatal("Timed out waiting for goroutine to complete")
	}
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)
}

func (suite *ServerTestSuite) TestSetJobStatusThrowSameStatusError() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{}, 1)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64, sjso ...metascheduler.SetJobStatusOption) error {
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
		suite.T().Fatal("Timed out waiting for goroutine to complete")
	}
	suite.Require().Nil(err)
	suite.Require().NotNil(resp)
}

func (suite *ServerTestSuite) TestSetJobStatusThrowTransitionError() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{}, 1)

	suite.jobHandler.EXPECT().
		SetJobStatus(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).RunAndReturn(func(ctx context.Context, b [32]byte, js metascheduler.JobStatus, u uint64, sjso ...metascheduler.SetJobStatusOption) error {
		done <- struct{}{}
		return &metascheduler.InvalidTransition{
			From: metascheduler.JobStatusScheduled,
			To:   js,
		}
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
			suite.T().Fatal("Timed out waiting for goroutine to complete")
		}
	}

	suite.NoError(err)
	suite.NotNil(resp)
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, &ServerTestSuite{})
}
