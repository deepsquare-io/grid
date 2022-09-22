package slurm_test

import (
	"os"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils"
	"github.com/stretchr/testify/suite"
)

type JobServiceTestSuite struct {
	suite.Suite
	address   string
	adminUser string
	user      string
	pkB64     string
	impl      slurm.JobService
}

func (suite *JobServiceTestSuite) submitJob() *slurm.SubmitJobRequest {
	// Arrange
	name := utils.GenerateRandomString(6)
	req := &slurm.SubmitJobRequest{
		Name: name,
		User: suite.user,
		JobDefinition: slurm.JobDefinition{
			TimeLimit:     5,
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
	_, err := suite.impl.SubmitJob(req)

	// Assert
	suite.NoError(err)

	return req
}

func (suite *JobServiceTestSuite) BeforeTest(suiteName, testName string) {
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

func (suite *JobServiceTestSuite) TestSubmit() {
	suite.submitJob()
}

func (suite *JobServiceTestSuite) TestCancel() {
	// Arrange
	req := suite.submitJob()

	// Act
	err := suite.impl.CancelJob(&slurm.CancelJobRequest{
		Name: req.Name,
		User: suite.user,
	})

	// Assert
	suite.NoError(err)
}

func (suite *JobServiceTestSuite) TestTopUp() {
	// Arrange
	req := suite.submitJob()

	// Act
	err := suite.impl.TopUp(&slurm.TopUpRequest{
		Name:           req.Name,
		User:           suite.user,
		AdditionalTime: 5,
	})

	// Assert
	suite.NoError(err)
}

func TestJobServiceTestSuite(t *testing.T) {
	address := os.Getenv("SLURM_SSH_ADDRESS")
	user := os.Getenv("SLURM_SSH_USER")
	adminUser := os.Getenv("SLURM_ADMIN_SSH_USER")
	pkB64 := os.Getenv("SLURM_SSH_PRIVATE_KEY")
	// Skip test if not defined
	if address == "" || user == "" || pkB64 == "" {
		logger.I.Warn("mandatory variables are not set!")
	} else {
		suite.Run(t, &JobServiceTestSuite{
			address:   address,
			user:      user,
			adminUser: adminUser,
			pkB64:     pkB64,
		})
	}
}
