//go:build integration

package slurm_test

import (
	"context"
	"os"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/stretchr/testify/suite"
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
			TimeLimit:     uint64(5),
			NTasks:        1,
			GPUsPerNode:   0,
			CPUsPerTask:   1,
			MemoryPerNode: 512,
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
	suite.impl = slurm.New(
		suite.address,
		suite.pkB64,
		suite.adminUser,
		"scancel",
		"sbatch",
		"squeue",
		"scontrol",
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
	address := os.Getenv("SLURM_SSH_ADDRESS")
	user := os.Getenv("SLURM_SSH_USER")
	adminUser := os.Getenv("SLURM_ADMIN_SSH_USER")
	pkB64 := os.Getenv("SLURM_SSH_PRIVATE_KEY")
	// Skip test if not defined
	if address == "" || user == "" || pkB64 == "" {
		logger.I.Warn("mandatory variables are not set!")
	} else {
		suite.Run(t, &ServiceTestSuite{
			address:   address,
			user:      user,
			adminUser: adminUser,
			pkB64:     pkB64,
		})
	}
}
