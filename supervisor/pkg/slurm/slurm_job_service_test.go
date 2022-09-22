package slurm_test

import (
	"os"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type JobServiceTestSuite struct {
	suite.Suite
	address   string
	adminUser string
	user      string
	pkB64     string
	impl      slurm.JobService
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

func (suite *JobServiceTestSuite) TestSubmitAndCancel() {
	req := &slurm.SubmitJobRequest{
		Name: "test",
		User: suite.user,
		JobDefinition: slurm.JobDefinition{
			TimeLimit:     5,
			NTasks:        1,
			GPUsPerNode:   0,
			CPUsPerTask:   1,
			MemoryPerNode: 1024,
			Body: `#!/bin/sh

srun hostname
`,
		},
	}
	jobID, err := suite.impl.SubmitJob(req)

	logger.I.Info("job submitted successfully", zap.Int("JobID", jobID))

	suite.NoError(err)

	suite.impl.CancelJob(&slurm.CancelJobRequest{
		Name: req.Name,
		User: suite.user,
	})
}

func (suite *JobServiceTestSuite) TestTopUp() {
	// Arrange
	req := &slurm.SubmitJobRequest{
		Name: "test",
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
	_, err := suite.impl.SubmitJob(req)

	suite.NoError(err)

	// Act
	err = suite.impl.TopUp(&slurm.TopUpRequest{
		Name:           req.Name,
		User:           suite.user,
		AdditionalTime: 5,
	})

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
