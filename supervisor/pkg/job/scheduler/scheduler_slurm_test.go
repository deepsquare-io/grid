// Copyright (C) 2023 DeepSquare Association
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

package scheduler_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/deepsquare-io/grid/supervisor/mocks/mockscheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	user    = "fakeUser"
	admin   = "fakeAdmin"
	address = "127.0.0.1:22"
	pkB64   = "private key"
)

type ServiceTestSuite struct {
	suite.Suite
	ssh  *mockscheduler.Executor
	impl scheduler.Scheduler
}

func (suite *ServiceTestSuite) BeforeTest(suiteName, testName string) {
	suite.ssh = mockscheduler.NewExecutor(suite.T())
	suite.impl = scheduler.NewSlurm(
		suite.ssh,
		admin,
		"localhost",
		"main",
	)
}

func (suite *ServiceTestSuite) TestCancel() {
	// Arrange
	name := utils.GenerateRandomString(6)
	suite.ssh.EXPECT().ExecAs(
		mock.Anything,
		user,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "scancel") &&
				strings.Contains(cmd, name)
		}),
	).Return("ok", nil)
	ctx := context.Background()

	// Act
	err := suite.impl.CancelJob(ctx, name, user)

	// Assert
	suite.NoError(err)
	suite.ssh.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestSubmit() {
	// Arrange
	name := utils.GenerateRandomString(6)
	expectedJobID := "123"
	req := &scheduler.SubmitRequest{
		Name:   name,
		User:   user,
		Prefix: "supervisor",
		JobDefinition: &scheduler.JobDefinition{
			TimeLimit:    uint64(5),
			NTasks:       1,
			GPUs:         utils.Ptr(uint64(0)),
			CPUsPerTask:  1,
			MemoryPerCPU: 512,
			Body: `#!/bin/sh

srun sleep infinity
`,
		},
	}
	suite.ssh.EXPECT().ExecAs(
		mock.Anything,
		user,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "sbatch") &&
				strings.Contains(cmd, "--cpus-per-task="+strconv.FormatUint(req.CPUsPerTask, 10)) &&
				strings.Contains(cmd, "--gpus="+strconv.FormatUint(*req.GPUs, 10)) &&
				strings.Contains(cmd, "--mem-per-cpu="+strconv.FormatUint(req.MemoryPerCPU, 10)) &&
				strings.Contains(cmd, "--ntasks="+strconv.FormatUint(req.NTasks, 10)) &&
				strings.Contains(cmd, "--time="+strconv.FormatUint(req.TimeLimit, 10)) &&
				strings.Contains(cmd, req.Name) &&
				strings.Contains(cmd, req.Prefix) &&
				strings.Contains(cmd, req.Body)
		}),
	).Return(fmt.Sprintf("%s\n", expectedJobID), nil)
	ctx := context.Background()

	// Act
	jobID, err := suite.impl.Submit(ctx, req)

	// Assert
	suite.NoError(err)
	suite.Equal(expectedJobID, jobID)
	suite.ssh.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestTopUp() {
	// Arrange
	name := utils.GenerateRandomString(6)
	jobID := "123"
	additionalTime := uint64(30)
	suite.ssh.EXPECT().ExecAs(
		mock.Anything,
		admin,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "squeue") &&
				strings.Contains(cmd, name)
		}),
	).Return(jobID, nil)
	suite.ssh.EXPECT().ExecAs(
		mock.Anything,
		admin,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "scontrol") &&
				strings.Contains(cmd, jobID) &&
				strings.Contains(cmd, strconv.FormatUint(additionalTime, 10))
		}),
	).Return("ok", nil)
	ctx := context.Background()

	// Act
	err := suite.impl.TopUp(ctx, name, additionalTime)

	// Assert
	suite.NoError(err)
	suite.ssh.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestHealthCheck() {
	// Arrange
	suite.ssh.EXPECT().ExecAs(
		mock.Anything,
		admin,
		"timeout 10 squeue",
	).Return("ok", nil)
	ctx := context.Background()

	// Act
	err := suite.impl.HealthCheck(ctx)

	// Assert
	suite.NoError(err)
	suite.ssh.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestFindRunningJobByName() {
	// Arrange
	name := utils.GenerateRandomString(6)
	jobID := 123
	suite.ssh.EXPECT().ExecAs(
		mock.Anything,
		user,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "squeue") &&
				strings.Contains(cmd, name)
		}),
	).Return(fmt.Sprintf("%d\n", jobID), nil)
	ctx := context.Background()

	// Act
	out, err := suite.impl.FindRunningJobByName(ctx, name, user)

	// Assert
	suite.NoError(err)
	suite.Equal(jobID, out)
	suite.ssh.AssertExpectations(suite.T())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, &ServiceTestSuite{})
}
