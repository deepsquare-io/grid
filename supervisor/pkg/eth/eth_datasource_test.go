// go:build unit

package eth_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/gen/go/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DataSourceTestSuite struct {
	suite.Suite
	msRPC           *metascheduler.MetaScheduler
	msWS            *metascheduler.MetaScheduler
	deployBackend   *mocks.DeployBackend
	contractBackend *mocks.ContractBackend
	impl            *eth.DataSource
}

var (
	gasPrice                 = big.NewInt(1)
	chainID                  = big.NewInt(1)
	nonce                    = uint64(1)
	jobID                    = [32]byte{1}
	jobDuration              = uint64(1000)
	privateKey               *ecdsa.PrivateKey
	fromAddress              common.Address
	metaschedulerAddress     = common.HexToAddress("0x1")
	providerJobQueuesAddress = common.HexToAddress("0x2")
)

func (suite *DataSourceTestSuite) mockProviderJobQueue() {
	// Pack input
	input, err := eth.MetaschedulerABI.Pack("providerJobQueues")
	suite.Require().NoError(err)
	output, err := eth.MetaschedulerABI.Methods["providerJobQueues"].Outputs.Pack(
		providerJobQueuesAddress,
	)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.On("CallContract", mock.Anything, ethereum.CallMsg{
		To:   &metaschedulerAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *DataSourceTestSuite) mockProviderJobQueuesContractCall(
	name string,
	inputs []interface{},
	outputs []interface{},
) {
	// Pack input
	abi, err := metascheduler.IProviderJobQueuesMetaData.GetAbi()
	suite.Require().NoError(err)
	input, err := abi.Pack(name, inputs...)
	suite.Require().NoError(err)
	output, err := abi.Methods[name].Outputs.Pack(outputs...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.On("CallContract", mock.Anything, ethereum.CallMsg{
		To:   &providerJobQueuesAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *DataSourceTestSuite) mockContractCall(
	name string,
	inputs []interface{},
	outputs []interface{},
) {
	// Pack input
	input, err := eth.MetaschedulerABI.Pack(name, inputs...)
	suite.Require().NoError(err)
	output, err := eth.MetaschedulerABI.Methods[name].Outputs.Pack(outputs...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.On("CallContract", mock.Anything, ethereum.CallMsg{
		To:   &metaschedulerAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *DataSourceTestSuite) mockContractTransaction(name string, args ...interface{}) {
	// Mock code presence
	suite.contractBackend.On(
		"PendingCodeAt",
		mock.Anything,
		metaschedulerAddress,
	).Return(
		common.FromHex("0xdeadface"),
		nil,
	)

	// Pack input
	abi, err := metascheduler.MetaSchedulerMetaData.GetAbi()
	suite.Require().NoError(err)
	input, err := abi.Pack(name, args...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.On(
		"EstimateGas",
		mock.Anything,
		mock.MatchedBy(func(call ethereum.CallMsg) bool {
			return call.From == fromAddress && *call.To == metaschedulerAddress &&
				bytes.Equal(call.Data, input)
		}),
	).Return(
		uint64(0),
		nil,
	)
	suite.contractBackend.On("SendTransaction", mock.Anything, mock.Anything).Return(nil)
}

func generateAddress() (pk *ecdsa.PrivateKey, address common.Address) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logger.I.Fatal("couldn't create pk", zap.Error(err))
	}
	pk = privateKey
	address = crypto.PubkeyToAddress(privateKey.PublicKey)
	return pk, address
}

func init() {
	privateKey, fromAddress = generateAddress()
}

func (suite *DataSourceTestSuite) BeforeTest(suiteName, testName string) {
	suite.contractBackend = mocks.NewContractBackend(suite.T())

	// Assert calling providerJobQueues
	suite.mockProviderJobQueue()

	msRPC, err := metascheduler.NewMetaScheduler(metaschedulerAddress, suite.contractBackend)
	suite.Require().NoError(err)
	suite.msRPC = msRPC
	msWS, err := metascheduler.NewMetaScheduler(metaschedulerAddress, suite.contractBackend)
	suite.Require().NoError(err)
	suite.msWS = msWS
	suite.deployBackend = mocks.NewDeployBackend(suite.T())
	suite.impl = eth.New(
		chainID,
		metaschedulerAddress,
		suite.deployBackend,
		suite.contractBackend,
		suite.contractBackend,
		suite.msRPC,
		suite.msWS,
		privateKey,
	)
}

func (suite *DataSourceTestSuite) assertMocksExpectations() {
	suite.contractBackend.AssertExpectations(suite.T())
	suite.deployBackend.AssertExpectations(suite.T())
}

func (suite *DataSourceTestSuite) mustAuthenticate() {
	// Must fetch nonce
	suite.contractBackend.On("PendingNonceAt", mock.Anything, fromAddress).Return(nonce, nil)
	// Must fetch gas price
	suite.contractBackend.On("SuggestGasPrice", mock.Anything).Return(gasPrice, nil)
}

func (suite *DataSourceTestSuite) TestClaim() {
	// Arrange
	suite.mustAuthenticate()
	suite.mockProviderJobQueuesContractCall(
		"hasNextClaimableJob",
		[]interface{}{fromAddress},
		[]interface{}{true},
	)
	suite.mockContractTransaction("claimNextJob")

	// Act
	err := suite.impl.Claim(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestClaimNoJob() {
	// Arrange
	// Must call ClaimNextJob
	suite.mockProviderJobQueuesContractCall(
		"hasNextClaimableJob",
		[]interface{}{fromAddress},
		[]interface{}{false},
	)

	// Act
	err := suite.impl.Claim(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestSetJobStatus() {
	// Arrange
	suite.mustAuthenticate()
	// Must call StartJob
	suite.mockContractTransaction(
		"providerSetJobStatus",
		jobID,
		uint8(eth.JobStatusFailed),
		jobDuration,
	)
	// Must wait
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
	suite.mockContractTransaction("refuseJob", jobID)

	// Act
	err := suite.impl.RefuseJob(context.Background(), jobID)

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestWatchEvents() {
	// Arrange
	claimNextJobEvents := make(chan *metascheduler.MetaSchedulerClaimJobEvent, 100)
	claimNextCancellingJobEvents := make(
		chan *metascheduler.MetaSchedulerClaimNextCancellingJobEvent,
		100,
	)
	claimNextTopUpJobEvents := make(chan *metascheduler.MetaSchedulerClaimNextTopUpJobEvent, 100)
	sub := mocks.NewSubscription(suite.T())
	errChan := make(chan error)
	var rErrChan <-chan error = errChan
	sub.On("Err").Return(rErrChan)
	sub.On("Unsubscribe")
	suite.contractBackend.On(
		"SubscribeFilterLogs",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(sub, nil)

	ctx, cancel := context.WithCancel(context.Background())

	// Act
	res, err := suite.impl.WatchEvents(
		ctx,
		claimNextTopUpJobEvents,
		claimNextCancellingJobEvents,
		claimNextJobEvents,
	)
	cancel()

	time.Sleep(time.Second)

	// Assert
	suite.NoError(err)
	suite.Equal(res, sub)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestClaimCancelling() {
	// Arrange
	suite.mustAuthenticate()
	suite.mockProviderJobQueuesContractCall(
		"hasCancellingJob",
		[]interface{}{fromAddress},
		[]interface{}{true},
	)
	suite.mockContractTransaction("claimNextCancellingJob")

	// Act
	err := suite.impl.ClaimCancelling(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestClaimTopUp() {
	// Arrange
	suite.mustAuthenticate()
	suite.mockProviderJobQueuesContractCall(
		"hasTopUpJob",
		[]interface{}{fromAddress},
		[]interface{}{true},
	)
	suite.mockContractTransaction("claimNextTopUpJob")

	// Act
	err := suite.impl.ClaimTopUp(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *DataSourceTestSuite) TestGetJobStatus() {
	// Arrange
	fixtureStatus := eth.JobStatusRunning
	suite.mockContractCall("jobs", []interface{}{jobID}, []interface{}{
		jobID,
		uint8(fixtureStatus),
		common.HexToAddress("0xdeadface"),
		common.HexToAddress("0xdeadface"),
		metascheduler.JobDefinition{},
		false,
		metascheduler.JobCost{
			MaxCost:      new(big.Int),
			FinalCost:    new(big.Int),
			PendingTopUp: new(big.Int),
		},
		metascheduler.JobTime{
			Start:                  new(big.Int),
			End:                    new(big.Int),
			CancelRequestTimestamp: new(big.Int),
			BlockNumberStateChange: new(big.Int),
		},
		[32]byte{1},
		false,
	})
	// Act
	status, err := suite.impl.GetJobStatus(context.Background(), jobID)
	// Assert
	suite.NoError(err)
	suite.Equal(fixtureStatus, status)
}

func (suite *DataSourceTestSuite) TestClaimCancellingNoCancelling() {
	// Arrange
	suite.mockProviderJobQueuesContractCall(
		"hasCancellingJob",
		[]interface{}{fromAddress},
		[]interface{}{false},
	)

	// Act
	err := suite.impl.ClaimCancelling(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func TestDataSourceTestSuite(t *testing.T) {
	suite.Run(t, &DataSourceTestSuite{})
}
