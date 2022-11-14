// go:build unit

package eth_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

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
	ms            *mocks.MetaScheduler
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
	suite.ms = mocks.NewMetaScheduler(suite.T())

	suite.impl = eth.New(
		suite.authenticator,
		suite.ms,
		privateKey,
	)
}

func (suite *DataSourceTestSuite) assertMocksExpectations() {
	suite.authenticator.AssertExpectations(suite.T())
	suite.ms.AssertExpectations(suite.T())
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
	suite.ms.On(
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

func (suite *DataSourceTestSuite) TestStartJob() {
	// Arrange
	suite.mustAuthenticate()
	// Must call StartJob
	tx := legacyTx()
	suite.ms.On(
		"StartJob",
		mock.MatchedBy(func(auth *bind.TransactOpts) bool {
			return auth.Nonce.Cmp(big.NewInt(0).SetUint64(nonce)) == 0 && auth.GasPrice == gasPrice
		}),
		jobID,
	).Return(tx, nil)

	// Act
	err := suite.impl.StartJob(context.Background(), jobID)

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestFinishJob() {
	// Arrange
	suite.mustAuthenticate()
	// Must call FinishJob
	tx := legacyTx()
	suite.ms.On(
		"FinishJob",
		mock.MatchedBy(func(auth *bind.TransactOpts) bool {
			return auth.Nonce.Cmp(big.NewInt(0).SetUint64(nonce)) == 0 && auth.GasPrice == gasPrice
		}),
		jobID,
		jobDuration,
	).Return(tx, nil)

	// Act
	err := suite.impl.FinishJob(context.Background(), jobID, jobDuration)

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestFailedJob() {
	// Arrange
	suite.mustAuthenticate()
	// Must call TriggerFailedJob
	tx := legacyTx()
	suite.ms.On(
		"TriggerFailedJob",
		mock.MatchedBy(func(auth *bind.TransactOpts) bool {
			return auth.Nonce.Cmp(big.NewInt(0).SetUint64(nonce)) == 0 && auth.GasPrice == gasPrice
		}),
		jobID,
	).Return(tx, nil)

	// Act
	err := suite.impl.FailJob(context.Background(), jobID)

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestRefuseJob() {
	// Arrange
	suite.mustAuthenticate()
	// Must call RefuseJob
	tx := legacyTx()
	suite.ms.On(
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

func TestDataSourceTestSuite(t *testing.T) {
	suite.Run(t, &DataSourceTestSuite{})
}
