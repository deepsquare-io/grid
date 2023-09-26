//go:build integration

package scheduler_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/ssh"
	"github.com/deepsquare-io/grid/supervisor/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ServiceIntegrationTestSuite struct {
	suite.Suite
	address   string
	adminUser string
	user      string
	pkB64     string
	impl      scheduler.Scheduler
}

func (suite *ServiceIntegrationTestSuite) submitJob() *scheduler.SubmitRequest {
	// Arrange
	ctx := context.Background()
	name := utils.GenerateRandomString(6)
	req := &scheduler.SubmitRequest{
		Name:   name,
		User:   suite.user,
		Prefix: "supervisor",
		JobDefinition: &scheduler.JobDefinition{
			TimeLimit:    uint64(5),
			NTasks:       1,
			GPUsPerTask:  utils.Ptr(uint64(0)),
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

func (suite *ServiceIntegrationTestSuite) BeforeTest(suiteName, testName string) {
	service := ssh.New(
		suite.address,
		suite.pkB64,
	)
	suite.impl = scheduler.NewSlurm(
		service,
		suite.adminUser,
		"localhost",
		"main",
	)
}

func (suite *ServiceIntegrationTestSuite) TestSubmit() {
	suite.submitJob()
}

func (suite *ServiceIntegrationTestSuite) TestCancel() {
	// Arrange
	ctx := context.Background()
	req := suite.submitJob()

	// Act
	err := suite.impl.CancelJob(ctx, req.Name, suite.user)

	// Assert
	suite.NoError(err)
}

func (suite *ServiceIntegrationTestSuite) TestTopUp() {
	// Arrange
	ctx := context.Background()
	req := suite.submitJob()

	// Act
	err := suite.impl.TopUp(ctx, req.Name, 5)

	// Assert
	suite.NoError(err)
}

func (suite *ServiceIntegrationTestSuite) TestFindMemPerNode() {
	// Arrange
	ctx := context.Background()

	// Act
	mem, err := suite.impl.FindMemPerNode(ctx)

	// Assert
	suite.NoError(err)
	fmt.Println(mem)
}

func (suite *ServiceIntegrationTestSuite) TestFindCPUsPerNode() {
	// Arrange
	ctx := context.Background()

	// Act
	cpu, err := suite.impl.FindCPUsPerNode(ctx)

	// Assert
	suite.NoError(err)
	fmt.Println(cpu)
}

func (suite *ServiceIntegrationTestSuite) TestFindGPUsPerNode() {
	// Arrange
	ctx := context.Background()

	// Act
	gpu, err := suite.impl.FindGPUsPerNode(ctx)

	// Assert
	suite.NoError(err)
	fmt.Println(gpu)
}

func (suite *ServiceIntegrationTestSuite) TestFindTotalGPUs() {
	// Arrange
	ctx := context.Background()

	// Act
	ret, err := suite.impl.FindTotalGPUs(ctx)

	// Assert
	suite.NoError(err)
	fmt.Println(ret)
}

func (suite *ServiceIntegrationTestSuite) TestFindTotalMem() {
	// Arrange
	ctx := context.Background()

	// Act
	ret, err := suite.impl.FindTotalMem(ctx)

	// Assert
	suite.NoError(err)
	fmt.Println(ret)
}

func (suite *ServiceIntegrationTestSuite) TestFindTotalCPUs() {
	// Arrange
	ctx := context.Background()

	// Act
	ret, err := suite.impl.FindTotalCPUs(ctx)

	// Assert
	suite.NoError(err)
	fmt.Println(ret)
}

func (suite *ServiceIntegrationTestSuite) TestFindTotalNodes() {
	// Arrange
	ctx := context.Background()

	// Act
	ret, err := suite.impl.FindTotalNodes(ctx)

	// Assert
	suite.NoError(err)
	fmt.Println(ret)
}

func TestServiceIntegrationTestSuite(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		// Skip test if not defined
		logger.I.Error("Error loading .env.test file", zap.Error(err))
	} else {
		suite.Run(t, &ServiceIntegrationTestSuite{
			address:   os.Getenv("SLURM_SSH_ADDRESS"),
			user:      os.Getenv("SLURM_SSH_USER"),
			adminUser: os.Getenv("SLURM_ADMIN_SSH_USER"),
			pkB64:     os.Getenv("SLURM_SSH_PRIVATE_KEY"),
		})
	}
}
