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
}

func (suite *BenchmarkIntegrationTestSuite) TestRunPhase1SingleNode() {
	ctx := context.Background()
	cpusPerNode, err := suite.scheduler.FindCPUsPerNode(ctx)
	suite.Require().NoError(err)
	gpusPerNode, err := suite.scheduler.FindGPUsPerNode(ctx)
	suite.Require().NoError(err)
	memPerNode, err := suite.scheduler.FindMemPerNode(ctx)
	suite.Require().NoError(err)
	suite.impl = benchmark.NewLauncher(
		"registry-1.deepsquare.run#library/hpc-benchmarks:23.5",
		"root",
		suite.publicAddress,
		suite.scheduler,
		1,
		cpusPerNode,
		memPerNode,
		gpusPerNode,
		2*time.Hour,
		benchmark.WithNoWait(),
	)

	err = suite.impl.RunPhase1(context.Background())

	suite.Require().NoError(err)
}

func (suite *BenchmarkIntegrationTestSuite) TestRunPhase1ThreeNodes() {
	ctx := context.Background()
	cpusPerNode, err := suite.scheduler.FindCPUsPerNode(ctx)
	suite.Require().NoError(err)
	gpusPerNode, err := suite.scheduler.FindGPUsPerNode(ctx)
	suite.Require().NoError(err)
	memPerNode, err := suite.scheduler.FindMemPerNode(ctx)
	suite.Require().NoError(err)
	suite.impl = benchmark.NewLauncher(
		"registry-1.deepsquare.run#library/hpc-benchmarks:23.5",
		"root",
		suite.publicAddress,
		suite.scheduler,
		3,
		cpusPerNode,
		memPerNode,
		gpusPerNode,
		2*time.Hour,
		// benchmark.WithUCX("mlx5_2:1|mlx5_2:1|mlx5_2:1|mlx5_2:1|mlx5_2:1|mlx5_2:1", ""),
		benchmark.WithUCX("eno2np1|eno2np1|eno2np1|eno2np1|eno2np1|eno2np1", ""),
		benchmark.WithNoWait(),
	)

	err = suite.impl.RunPhase1(context.Background())

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
