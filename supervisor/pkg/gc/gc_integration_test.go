// Copyright (C) 2023 DeepSquare Asociation
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

package gc_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/pkg/gc"
	"github.com/deepsquare-io/grid/supervisor/pkg/job/scheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/ssh"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type GCTestSuite struct {
	suite.Suite
	impl          *gc.GC
	rpcURL        string
	wsURL         string
	ethHexPK      string
	smartContract string
	address       string
	adminUser     string
	user          string
	pkB64         string
	fromAddress   common.Address
}

func (suite *GCTestSuite) BeforeTest(suiteName, testName string) {
	ctx := context.Background()
	rpcClient, err := rpc.DialOptions(
		ctx,
		suite.rpcURL,
		rpc.WithHTTPClient(http.DefaultClient),
	)
	if err != nil {
		logger.I.Fatal("ethclientRPC dial failed", zap.Error(err))
	}
	ethClientRPC := ethclient.NewClient(rpcClient)
	wsClient, err := rpc.DialOptions(
		ctx,
		suite.wsURL,
		rpc.WithHTTPClient(http.DefaultClient),
	)
	if err != nil {
		logger.I.Fatal("ethclientWS dial failed", zap.Error(err))
	}
	ethClientWS := ethclient.NewClient(wsClient)
	pk, err := crypto.HexToECDSA(suite.ethHexPK)
	if err != nil {
		logger.I.Fatal("couldn't decode private key", zap.Error(err))
	}
	chainID, err := ethClientRPC.ChainID(ctx)
	if err != nil {
		logger.I.Fatal("couldn't fetch chainID", zap.Error(err))
	}
	suite.fromAddress = crypto.PubkeyToAddress(pk.PublicKey)
	ms := metascheduler.NewClient(
		chainID,
		common.HexToAddress(suite.smartContract),
		ethClientRPC,
		ethClientRPC,
		ethClientWS,
		pk,
	)

	service := ssh.New(
		suite.address,
		suite.pkB64,
	)
	sched := scheduler.NewSlurm(
		service,
		suite.adminUser,
		"localhost",
		"main",
	)

	suite.impl = gc.NewGC(
		ms,
		sched,
	)
}

func (suite *GCTestSuite) TestFindUnhandledJobs() {
	ctx := context.Background()
	err := suite.impl.FindAndCancelUnhandledJobs(ctx)
	suite.Require().NoError(err)
}

func (suite *GCTestSuite) TestLoop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := suite.impl.Loop(ctx)
	suite.Require().EqualError(err, context.DeadlineExceeded.Error())
}

func TestGCTestSuite(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		// Skip test if not defined
		logger.I.Error("Error loading .env.test file", zap.Error(err))
	} else {
		suite.Run(t, &GCTestSuite{
			smartContract: os.Getenv("METASCHEDULER_SMART_CONTRACT"),
			ethHexPK:      os.Getenv("ETH_PRIVATE_KEY"),
			rpcURL:        os.Getenv("METASCHEDULER_ENDPOINT_RPC"),
			wsURL:         os.Getenv("METASCHEDULER_ENDPOINT_WS"),
			address:       os.Getenv("SLURM_SSH_ADDRESS"),
			user:          os.Getenv("SLURM_SSH_USER"),
			adminUser:     os.Getenv("SLURM_ADMIN_SSH_USER"),
			pkB64:         os.Getenv("SLURM_SSH_PRIVATE_KEY"),
		})
	}
}
