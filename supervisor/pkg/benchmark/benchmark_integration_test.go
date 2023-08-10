//go:build integration

package benchmark_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/ssh"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type BenchmarkIntegrationTestSuite struct {
	suite.Suite
	address       string
	adminUser     string
	user          string
	pkB64         string
	publicAddress string
	scheduler     scheduler.Scheduler
	impl          benchmark.Launcher
}

func (suite *BenchmarkIntegrationTestSuite) BeforeTest(suiteName, testName string) {
	ctx := context.Background()
	service := ssh.New(
		suite.address,
		suite.pkB64,
	)
	suite.scheduler = scheduler.NewSlurm(
		service,
		suite.adminUser,
		suite.publicAddress,
		"main",
	)
	cpusPerNode, err := suite.scheduler.FindCPUsPerNode(ctx)
	suite.Require().NoError(err)
	gpusPerNode, err := suite.scheduler.FindGPUsPerNode(ctx)
	suite.Require().NoError(err)
	memPerNode, err := suite.scheduler.FindMemPerNode(ctx)
	suite.Require().NoError(err)
	suite.impl = benchmark.NewLauncher(
		"registry-1.deepsquare.run#library/hpc-benchmarks:21.4-hpl",
		"root",
		suite.publicAddress,
		suite.scheduler,
		1,
		cpusPerNode,
		memPerNode,
		gpusPerNode,
		time.Hour,
	)
}

func (suite *BenchmarkIntegrationTestSuite) TestRunPhase1() {
	err := suite.impl.RunPhase1(context.Background())

	suite.Require().NoError(err)
}

func (suite *BenchmarkIntegrationTestSuite) TestRunPhase2() {
	err := suite.impl.RunPhase2(
		context.Background(),
		2,
		2,
		95000,
		1024,
	)

	suite.Require().NoError(err)
}

func TestBenchmarkIntegrationTestSuite(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		// Skip test if not defined
		logger.I.Error("Error loading .env.test file", zap.Error(err))
	} else {
		suite.Run(t, &BenchmarkIntegrationTestSuite{
			address:       os.Getenv("SLURM_SSH_ADDRESS"),
			user:          os.Getenv("SLURM_SSH_USER"),
			adminUser:     os.Getenv("SLURM_ADMIN_SSH_USER"),
			pkB64:         os.Getenv("SLURM_SSH_PRIVATE_KEY"),
			publicAddress: os.Getenv("PUBLIC_ADDRESS"),
		})
	}
}
