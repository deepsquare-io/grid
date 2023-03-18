// go:build unit

package eth_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DataSourceTestSuite struct {
	suite.Suite
	authenticator *mocks.EthereumAuthenticator
	msRPC         *mocks.MetaSchedulerRPC
	msWS          *mocks.MetaSchedulerWS
	deployBackend *mocks.DeployBackend
	impl          *eth.DataSource
}

var (
	gasPrice    = big.NewInt(1)
	chainID     = big.NewInt(1)
	nonce       = uint64(1)
	jobID       = [32]byte{1}
	jobDuration = uint64(1000)
	privateKey  *ecdsa.PrivateKey
	fromAddress common.Address
)

func generateAddress() (pk *ecdsa.PrivateKey, address common.Address) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logger.I.Fatal("couldn't create pk", zap.Error(err))
	}
	pk = privateKey
	publicKey := privateKey.Public()
	pubECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.I.Fatal("error casting public key to ECDSA")
	}
	address = crypto.PubkeyToAddress(*pubECDSA)
	return pk, address
}

func init() {
	privateKey, fromAddress = generateAddress()
}

func (suite *DataSourceTestSuite) BeforeTest(suiteName, testName string) {
	suite.authenticator = mocks.NewEthereumAuthenticator(suite.T())
	suite.msRPC = mocks.NewMetaSchedulerRPC(suite.T())
	suite.msWS = mocks.NewMetaSchedulerWS(suite.T())
	suite.deployBackend = mocks.NewDeployBackend(suite.T())
	suite.impl = eth.New(
		suite.authenticator,
		suite.deployBackend,
		suite.msRPC,
		suite.msWS,
		privateKey,
	)
}

func (suite *DataSourceTestSuite) assertMocksExpectations() {
	suite.authenticator.AssertExpectations(suite.T())
	suite.msWS.AssertExpectations(suite.T())
}

func (suite *DataSourceTestSuite) mustAuthenticate() {
	// Must fetch nonce
	suite.authenticator.On("PendingNonceAt", mock.Anything, fromAddress).Return(nonce, nil)
	// Must fetch gas price
	suite.authenticator.On("SuggestGasPrice", mock.Anything).Return(gasPrice, nil)
	// Must fetch chainID
	suite.authenticator.On("ChainID", mock.Anything).Return(chainID, nil)
}

// legacyTx creates a fake transaction
//
// The hash is 0xb4848204c8432070136a41792003caf8dea08f9eb284eb4240845bf64a66a068
func legacyTx() *types.Transaction {
	nonce, err := hexutil.DecodeUint64("0x1216")
	if err != nil {
		panic(err)
	}
	gasPrice, err := hexutil.DecodeBig("0x2bd0875aed")
	if err != nil {
		panic(err)
	}
	gas, err := hexutil.DecodeUint64("0x5208")
	if err != nil {
		panic(err)
	}
	to := common.HexToAddress("0x2f14582947e292a2ecd20c430b46f2d27cfe213c")
	value, err := hexutil.DecodeBig("0x2386f26fc10000")
	if err != nil {
		panic(err)
	}
	data := common.Hex2Bytes("0x")
	v, err := hexutil.DecodeBig("0x1")
	if err != nil {
		panic(err)
	}
	r, err := hexutil.DecodeBig("0x56b5bf9222ce26c3239492173249696740bc7c28cd159ad083a0f4940baf6d03")
	if err != nil {
		panic(err)
	}
	s, err := hexutil.DecodeBig("0x5fcd608b3b638950d3fe007b19ca8c4ead37237eaf89a8426777a594fd245c2a")
	if err != nil {
		panic(err)
	}

	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gas,
		To:       &to,
		Value:    value,
		Data:     data,
		V:        v,
		R:        r,
		S:        s,
	})
}

func (suite *DataSourceTestSuite) TestClaim() {
	// Arrange
	suite.mustAuthenticate()
	// Must call ClaimNextJob
	tx := legacyTx()
	suite.msRPC.On("HasNextJob", mock.Anything, fromAddress).Return(true, nil)
	suite.msRPC.On(
		"ClaimNextJob",
		mock.MatchedBy(func(auth *bind.TransactOpts) bool {
			return auth.Nonce.Cmp(big.NewInt(0).SetUint64(nonce)) == 0 && auth.GasPrice == gasPrice
		}),
	).Return(tx, nil)

	// Act
	err := suite.impl.Claim(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestClaimNoJob() {
	// Arrange
	// Must call ClaimNextJob
	suite.msRPC.On("HasNextJob", mock.Anything, fromAddress).Return(false, nil)

	// Act
	err := suite.impl.Claim(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
	suite.msRPC.AssertNotCalled(suite.T(), "ClaimNextJob", mock.Anything)
}

func (suite *DataSourceTestSuite) TestSetJobStatus() {
	// Arrange
	suite.mustAuthenticate()
	// Must call StartJob
	tx := legacyTx()
	suite.msRPC.On(
		"ProviderSetJobStatus",
		mock.MatchedBy(func(auth *bind.TransactOpts) bool {
			return auth.Nonce.Cmp(big.NewInt(0).SetUint64(nonce)) == 0 && auth.GasPrice == gasPrice
		}),
		jobID,
		uint8(eth.JobStatusFailed),
		jobDuration,
	).Return(tx, nil)
	suite.deployBackend.On("TransactionReceipt", mock.Anything, mock.Anything).Return(nil, nil)

	// Act
	err := suite.impl.SetJobStatus(context.Background(), jobID, eth.JobStatusFailed, jobDuration)

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestRefuseJob() {
	// Arrange
	suite.mustAuthenticate()
	// Must call RefuseJob
	tx := legacyTx()
	suite.msRPC.On(
		"RefuseJob",
		mock.MatchedBy(func(auth *bind.TransactOpts) bool {
			return auth.Nonce.Cmp(big.NewInt(0).SetUint64(nonce)) == 0 && auth.GasPrice == gasPrice
		}),
		jobID,
	).Return(tx, nil)

	// Act
	err := suite.impl.RefuseJob(context.Background(), jobID)

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestWatchClaimNextJobEvent() {
	// Arrange
	sink := make(chan *metascheduler.MetaSchedulerClaimJobEvent)
	defer close(sink)
	sub := mocks.NewSubscription(suite.T())
	suite.msWS.On(
		"WatchClaimJobEvent",
		mock.Anything,
		mock.Anything,
	).Return(sub, nil)

	// Act
	res, err := suite.impl.WatchClaimNextJobEvent(context.Background(), sink)

	// Assert
	suite.NoError(err)
	suite.Equal(res, sub)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestWatchClaimNextCancellingJobEvent() {
	// Arrange
	sink := make(chan *metascheduler.MetaSchedulerClaimNextCancellingJobEvent)
	defer close(sink)
	sub := mocks.NewSubscription(suite.T())
	suite.msWS.On(
		"WatchClaimNextCancellingJobEvent",
		mock.Anything,
		mock.Anything,
	).Return(sub, nil)

	// Act
	res, err := suite.impl.WatchClaimNextCancellingJobEvent(context.Background(), sink)

	// Assert
	suite.NoError(err)
	suite.Equal(res, sub)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestClaimCancelling() {
	// Arrange
	suite.mustAuthenticate()
	tx := legacyTx()
	suite.msRPC.On("HasCancellingJob", mock.Anything, fromAddress).Return(true, nil)
	suite.msRPC.On(
		"ClaimNextCancellingJob",
		mock.MatchedBy(func(auth *bind.TransactOpts) bool {
			return auth.Nonce.Cmp(big.NewInt(0).SetUint64(nonce)) == 0 && auth.GasPrice == gasPrice
		}),
	).Return(tx, nil)

	// Act
	err := suite.impl.ClaimCancelling(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestGetJobStatus() {
	// Arrange
	fixtureStatus := eth.JobStatusRunning
	suite.msRPC.On("Jobs", mock.Anything, jobID).Return(struct {
		JobId            [32]byte
		Status           uint8
		CustomerAddr     common.Address
		ProviderAddr     common.Address
		Definition       metascheduler.JobDefinition
		Valid            bool
		Cost             metascheduler.JobCost
		Time             metascheduler.JobTime
		JobName          [32]byte
		HasCancelRequest bool
	}{
		JobId:  jobID,
		Status: uint8(fixtureStatus),
	}, nil)
	// Act
	status, err := suite.impl.GetJobStatus(context.Background(), jobID)
	// Assert
	suite.NoError(err)
	suite.Equal(fixtureStatus, status)
}

func (suite *DataSourceTestSuite) TestClaimCancellingNoCancelling() {
	// Arrange
	suite.msRPC.On("HasCancellingJob", mock.Anything, fromAddress).Return(false, nil)

	// Act
	err := suite.impl.ClaimCancelling(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
	suite.msRPC.AssertNotCalled(suite.T(), "ClaimNextCancellingJob", mock.Anything)
}

func TestDataSourceTestSuite(t *testing.T) {
	suite.Run(t, &DataSourceTestSuite{})
}
