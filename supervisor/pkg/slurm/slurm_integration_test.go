//go:build integration

package slurm_test

import (
	"context"
	"os"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/ssh"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ServiceTestSuite struct {
	suite.Suite
	address   string
	adminUser string
	user      string
	pkB64     string
	impl      *slurm.Service
}

func (suite *ServiceTestSuite) submitJob() *slurm.SubmitJobRequest {
	// Arrange
	ctx := context.Background()
	name := utils.GenerateRandomString(6)
	req := &slurm.SubmitJobRequest{
		Name: name,
		User: suite.user,
		JobDefinition: &slurm.JobDefinition{
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

	// Act
	_, err := suite.impl.Submit(ctx, req)

	// Assert
	suite.NoError(err)

	return req
}

func (suite *ServiceTestSuite) BeforeTest(suiteName, testName string) {
	service := ssh.New(
		suite.address,
		suite.pkB64,
	)
	suite.impl = slurm.New(
		service,
		suite.adminUser,
		"scancel",
		"sbatch",
		"squeue",
		"scontrol",
		"localhost",
	)
}

func (suite *ServiceTestSuite) TestSubmit() {
	suite.submitJob()
}

func (suite *ServiceTestSuite) TestCancel() {
	// Arrange
	ctx := context.Background()
	req := suite.submitJob()

	// Act
	err := suite.impl.CancelJob(ctx, &slurm.CancelJobRequest{
		Name: req.Name,
		User: suite.user,
	})

	// Assert
	suite.NoError(err)
}

func (suite *ServiceTestSuite) TestTopUp() {
	// Arrange
	ctx := context.Background()
	req := suite.submitJob()

	// Act
	err := suite.impl.TopUp(ctx, &slurm.TopUpRequest{
		Name:           req.Name,
		User:           suite.user,
		AdditionalTime: 5,
	})

	// Assert
	suite.NoError(err)
}

func TestServiceTestSuite(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		// Skip test if not defined
		logger.I.Error("Error loading .env.test file", zap.Error(err))
	} else {
		suite.Run(t, &ServiceTestSuite{
			address:   os.Getenv("SLURM_SSH_ADDRESS"),
			user:      os.Getenv("SLURM_SSH_USER"),
			adminUser: os.Getenv("SLURM_ADMIN_SSH_USER"),
			pkB64:     os.Getenv("SLURM_SSH_PRIVATE_KEY"),
		})
	}
}
