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

//go:build integration

package benchmark_test

import (
	"context"
	"os"
	"slices"
	"testing"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/ssh"
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

func (suite *BenchmarkIntegrationTestSuite) TestRunHPLSingleNode() {
	ctx := context.Background()
	cpusPerNode, err := suite.scheduler.FindCPUsPerNode(ctx)
	suite.Require().NoError(err)
	gpusPerNode, err := suite.scheduler.FindGPUsPerNode(ctx)
	suite.Require().NoError(err)
	memPerNode, err := suite.scheduler.FindMemPerNode(ctx)
	suite.Require().NoError(err)
	suite.impl = benchmark.NewLauncher(
		"root",
		suite.publicAddress,
		suite.scheduler,
		benchmark.WithNoWait(),
	)

	b, err := benchmark.GenerateHPLBenchmark(
		benchmark.WithClusterSpecs(
			1,
			slices.Min(cpusPerNode),
			slices.Min(gpusPerNode),
			slices.Min(memPerNode),
		),
		benchmark.WithSupervisorPublicAddress(suite.publicAddress),
		benchmark.WithUCX("eno2np1|eno2np1|eno2np1", ""),
	)
	suite.Require().NoError(err)

	err = suite.impl.Launch(ctx, "test", b)
	suite.Require().NoError(err)
}

func (suite *BenchmarkIntegrationTestSuite) TestRunHPL() {
	ctx := context.Background()
	nodes, err := suite.scheduler.FindTotalNodes(ctx, scheduler.WithOnlyResponding())
	suite.Require().NoError(err)
	cpusPerNode, err := suite.scheduler.FindCPUsPerNode(ctx)
	suite.Require().NoError(err)
	gpusPerNode, err := suite.scheduler.FindGPUsPerNode(ctx)
	suite.Require().NoError(err)
	memPerNode, err := suite.scheduler.FindMemPerNode(ctx)
	suite.Require().NoError(err)
	suite.impl = benchmark.NewLauncher(
		"root",
		suite.publicAddress,
		suite.scheduler,
		benchmark.WithNoWait(),
	)

	b, err := benchmark.GenerateHPLBenchmark(
		benchmark.WithClusterSpecs(
			nodes,
			slices.Min(cpusPerNode),
			slices.Min(gpusPerNode),
			slices.Min(memPerNode),
		),
		benchmark.WithSupervisorPublicAddress(suite.publicAddress),
		benchmark.WithUCX("mlx5_2:1|mlx5_2:1|mlx5_2:1", ""),
		benchmark.WithTrace(),
	)
	suite.Require().NoError(err)

	err = suite.impl.Launch(ctx, "test", b)
	suite.Require().NoError(err)
}

func (suite *BenchmarkIntegrationTestSuite) TestRunSpeedTest() {
	ctx := context.Background()
	suite.impl = benchmark.NewLauncher(
		"root",
		suite.publicAddress,
		suite.scheduler,
		benchmark.WithNoWait(),
	)

	b, err := benchmark.GenerateSpeedTestBenchmark(
		benchmark.WithSupervisorPublicAddress(suite.publicAddress),
	)
	suite.Require().NoError(err)

	err = suite.impl.Launch(ctx, "test", b)
	suite.Require().NoError(err)
}

func (suite *BenchmarkIntegrationTestSuite) TestRunOSU() {
	ctx := context.Background()
	nodes, err := suite.scheduler.FindTotalNodes(ctx, scheduler.WithOnlyResponding())
	suite.Require().NoError(err)
	cpusPerNode, err := suite.scheduler.FindCPUsPerNode(ctx)
	suite.Require().NoError(err)
	gpusPerNode, err := suite.scheduler.FindGPUsPerNode(ctx)
	suite.Require().NoError(err)
	memPerNode, err := suite.scheduler.FindMemPerNode(ctx)
	suite.Require().NoError(err)
	suite.impl = benchmark.NewLauncher(
		"root",
		suite.publicAddress,
		suite.scheduler,
		benchmark.WithNoWait(),
	)

	b, err := benchmark.GenerateOSUBenchmark(
		benchmark.WithClusterSpecs(
			nodes,
			slices.Min(cpusPerNode),
			slices.Min(gpusPerNode),
			slices.Min(memPerNode),
		),
		benchmark.WithSupervisorPublicAddress(suite.publicAddress),
		benchmark.WithUCX("eno2np1|eno2np1|eno2np1", ""),
	)
	suite.Require().NoError(err)

	err = suite.impl.Launch(ctx, "test", b)
	suite.Require().NoError(err)
}

func (suite *BenchmarkIntegrationTestSuite) TestRunIOR() {
	ctx := context.Background()
	nodes, err := suite.scheduler.FindTotalNodes(ctx, scheduler.WithOnlyResponding())
	suite.Require().NoError(err)
	cpusPerNode, err := suite.scheduler.FindCPUsPerNode(ctx)
	suite.Require().NoError(err)
	gpusPerNode, err := suite.scheduler.FindGPUsPerNode(ctx)
	suite.Require().NoError(err)
	memPerNode, err := suite.scheduler.FindMemPerNode(ctx)
	suite.Require().NoError(err)
	suite.impl = benchmark.NewLauncher(
		"root",
		suite.publicAddress,
		suite.scheduler,
		benchmark.WithNoWait(),
	)

	b, err := benchmark.GenerateIORBenchmark(
		benchmark.WithClusterSpecs(
			nodes,
			slices.Min(cpusPerNode),
			slices.Min(gpusPerNode),
			slices.Min(memPerNode),
		),
		benchmark.WithSupervisorPublicAddress(suite.publicAddress),
		benchmark.WithUCX("eno2np1|eno2np1|eno2np1", ""),
	)
	suite.Require().NoError(err)

	err = suite.impl.Launch(ctx, "test", b)
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
