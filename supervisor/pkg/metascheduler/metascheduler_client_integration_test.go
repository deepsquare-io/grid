//go:build integration

package metascheduler_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ClientTestSuite struct {
	suite.Suite
	impl          metascheduler.MetaScheduler
	rpcURL        string
	wsURL         string
	ethHexPK      string
	smartContract string
	fromAddress   common.Address
}

func (suite *ClientTestSuite) BeforeTest(suiteName, testName string) {
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
	suite.impl = metascheduler.NewClient(
		chainID,
		common.HexToAddress(suite.smartContract),
		ethClientRPC,
		ethClientRPC,
		ethClientWS,
		pk,
	)
}

func (suite *ClientTestSuite) TestGetJobs() {
	ctx := context.Background()
	it, err := suite.impl.GetJobs(ctx)
	suite.Require().NoError(err)
	defer it.Close()

	for it.Next(ctx) {
		suite.Require().Equal(suite.fromAddress, it.Job.ProviderAddr)
	}
}

func TestClientTestSuite(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		// Skip test if not defined
		logger.I.Error("Error loading .env.test file", zap.Error(err))
	} else {
		suite.Run(t, &ClientTestSuite{
			smartContract: os.Getenv("METASCHEDULER_SMART_CONTRACT"),
			ethHexPK:      os.Getenv("ETH_PRIVATE_KEY"),
			rpcURL:        os.Getenv("METASCHEDULER_ENDPOINT_RPC"),
			wsURL:         os.Getenv("METASCHEDULER_ENDPOINT_WS"),
		})
	}
}
