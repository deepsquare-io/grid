//go:build unit

package scheduler_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
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
	ssh  *mocks.Executor
	impl *scheduler.Slurm
}

func (suite *ServiceTestSuite) BeforeTest(suiteName, testName string) {
	suite.ssh = mocks.NewExecutor(suite.T())
	suite.impl = scheduler.NewSlurm(
		suite.ssh,
		admin,
		"scancel",
		"sbatch",
		"squeue",
		"scontrol",
		"localhost",
	)
}

func (suite *ServiceTestSuite) TestCancel() {
	// Arrange
	name := utils.GenerateRandomString(6)
	req := &job.CancelRequest{
		Name: name,
		User: user,
	}
	suite.ssh.On(
		"ExecAs",
		mock.Anything,
		user,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "scancel") &&
				strings.Contains(cmd, req.Name)
		}),
	).Return("ok", nil)
	ctx := context.Background()

	// Act
	err := suite.impl.CancelJob(ctx, req)

	// Assert
	suite.NoError(err)
	suite.ssh.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestSubmit() {
	// Arrange
	name := utils.GenerateRandomString(6)
	expectedJobID := "123"
	req := &job.SubmitRequest{
		Name: name,
		User: user,
		Definition: &job.Definition{
			TimeLimit:    uint64(5),
			NTasks:       1,
			GPUsPerTask:  0,
			CPUsPerTask:  1,
			MemoryPerCPU: 512,
			Body: `#!/bin/sh

srun sleep infinity
`,
		},
	}
	suite.ssh.On(
		"ExecAs",
		mock.Anything,
		user,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "sbatch") &&
				strings.Contains(cmd, strconv.FormatUint(req.CPUsPerTask, 10)) &&
				strings.Contains(cmd, strconv.FormatUint(req.GPUsPerTask, 10)) &&
				strings.Contains(cmd, strconv.FormatUint(req.MemoryPerCPU, 10)) &&
				strings.Contains(cmd, strconv.FormatUint(req.NTasks, 10)) &&
				strings.Contains(cmd, strconv.FormatUint(req.TimeLimit, 10)) &&
				strings.Contains(cmd, req.Name) &&
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
	req := &job.TopUpRequest{
		Name:           name,
		AdditionalTime: 30,
	}
	suite.ssh.On(
		"ExecAs",
		mock.Anything,
		admin,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "squeue") &&
				strings.Contains(cmd, req.Name)
		}),
	).Return(jobID, nil)
	suite.ssh.On(
		"ExecAs",
		mock.Anything,
		admin,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "scontrol") &&
				strings.Contains(cmd, jobID) &&
				strings.Contains(cmd, strconv.FormatUint(req.AdditionalTime, 10))
		}),
	).Return("ok", nil)
	ctx := context.Background()

	// Act
	err := suite.impl.TopUp(ctx, req)

	// Assert
	suite.NoError(err)
	suite.ssh.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestHealthCheck() {
	// Arrange
	suite.ssh.On(
		"ExecAs",
		mock.Anything,
		admin,
		"squeue",
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
	req := &job.FindRunningJobByNameRequest{
		Name: name,
		User: user,
	}
	suite.ssh.On(
		"ExecAs",
		mock.Anything,
		user,
		mock.MatchedBy(func(cmd string) bool {
			return strings.Contains(cmd, "squeue") &&
				strings.Contains(cmd, name)
		}),
	).Return(fmt.Sprintf("%d\n", jobID), nil)
	ctx := context.Background()

	// Act
	out, err := suite.impl.FindRunningJobByName(ctx, req)

	// Assert
	suite.NoError(err)
	suite.Equal(jobID, out)
	suite.ssh.AssertExpectations(suite.T())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, &ServiceTestSuite{})
}