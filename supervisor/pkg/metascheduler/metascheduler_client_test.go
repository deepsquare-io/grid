// go:build unit

package metascheduler_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/meta-scheduler/mocks"
	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/mocks/mockbind"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ClientTestSuite struct {
	suite.Suite
	deployBackend   *mockbind.DeployBackend
	contractBackend *mockbind.ContractBackend
	impl            metascheduler.MetaScheduler
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

func (suite *ClientTestSuite) mockProviderJobQueue() {
	// Pack input
	input, err := metascheduler.MetaschedulerABI.Pack("providerJobQueues")
	suite.Require().NoError(err)
	output, err := metascheduler.MetaschedulerABI.Methods["providerJobQueues"].Outputs.Pack(
		providerJobQueuesAddress,
	)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.EXPECT().CallContract(mock.Anything, ethereum.CallMsg{
		To:   &metaschedulerAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *ClientTestSuite) mockProviderJobQueuesContractCall(
	name string,
	inputs []interface{},
	outputs []interface{},
) {
	// Pack input
	abi, err := metaschedulerabi.IProviderJobQueuesMetaData.GetAbi()
	suite.Require().NoError(err)
	input, err := abi.Pack(name, inputs...)
	suite.Require().NoError(err)
	output, err := abi.Methods[name].Outputs.Pack(outputs...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.EXPECT().CallContract(mock.Anything, ethereum.CallMsg{
		To:   &providerJobQueuesAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *ClientTestSuite) mockContractCall(
	name string,
	inputs []interface{},
	outputs []interface{},
) {
	// Pack input
	input, err := metascheduler.MetaschedulerABI.Pack(name, inputs...)
	suite.Require().NoError(err)
	output, err := metascheduler.MetaschedulerABI.Methods[name].Outputs.Pack(outputs...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.EXPECT().CallContract(mock.Anything, ethereum.CallMsg{
		To:   &metaschedulerAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *ClientTestSuite) mockContractTransaction(name string, args ...interface{}) {
	// Mock code presence
	suite.contractBackend.EXPECT().
		PendingCodeAt(
			mock.Anything,
			metaschedulerAddress,
		).Return(
		common.FromHex("0xdeadface"),
		nil,
	)

	// Pack input
	abi, err := metaschedulerabi.MetaSchedulerMetaData.GetAbi()
	suite.Require().NoError(err)
	input, err := abi.Pack(name, args...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.EXPECT().
		EstimateGas(
			mock.Anything,
			mock.MatchedBy(func(call ethereum.CallMsg) bool {
				return call.From == fromAddress && *call.To == metaschedulerAddress &&
					bytes.Equal(call.Data, input)
			}),
		).Return(
		uint64(0),
		nil,
	)
	suite.contractBackend.EXPECT().SendTransaction(mock.Anything, mock.Anything).Return(nil)
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

func (suite *ClientTestSuite) BeforeTest(suiteName, testName string) {
	suite.contractBackend = mockbind.NewContractBackend(suite.T())

	// Assert calling providerJobQueues
	suite.mockProviderJobQueue()

	suite.deployBackend = mockbind.NewDeployBackend(suite.T())
	suite.impl = metascheduler.NewClient(
		chainID,
		metaschedulerAddress,
		suite.deployBackend,
		suite.contractBackend,
		suite.contractBackend,
		privateKey,
	)
}

func (suite *ClientTestSuite) assertMocksExpectations() {
	suite.contractBackend.AssertExpectations(suite.T())
	suite.deployBackend.AssertExpectations(suite.T())
}

func (suite *ClientTestSuite) mustAuthenticate() {
	// Must fetch nonce
	suite.contractBackend.EXPECT().PendingNonceAt(mock.Anything, fromAddress).Return(nonce, nil)
	// Must fetch gas price
	suite.contractBackend.EXPECT().SuggestGasPrice(mock.Anything).Return(gasPrice, nil)
}

func (suite *ClientTestSuite) TestClaim() {
	// Arrange
	suite.mustAuthenticate()
	suite.mockProviderJobQueuesContractCall(
		"hasNextClaimableJob",
		[]interface{}{fromAddress},
		[]interface{}{true},
	)
	suite.mockContractTransaction("claimNextJob")
	// Must wait
	suite.deployBackend.EXPECT().TransactionReceipt(mock.Anything, mock.Anything).Return(nil, nil)

	// Act
	err := suite.impl.Claim(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *ClientTestSuite) TestClaimNoJob() {
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

func (suite *ClientTestSuite) TestSetJobStatus() {
	// Arrange
	suite.mustAuthenticate()
	// Must call StartJob
	suite.mockContractTransaction(
		"providerSetJobStatus",
		jobID,
		uint8(metascheduler.JobStatusFailed),
		jobDuration,
	)
	// Must wait
	suite.deployBackend.EXPECT().TransactionReceipt(mock.Anything, mock.Anything).Return(nil, nil)

	// Act
	err := suite.impl.SetJobStatus(
		context.Background(),
		jobID,
		metascheduler.JobStatusFailed,
		jobDuration,
	)

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *ClientTestSuite) TestRefuseJob() {
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

func (suite *ClientTestSuite) TestWatchEvents() {
	// Arrange
	claimNextJobEvents := make(chan *metaschedulerabi.MetaSchedulerClaimJobEvent, 100)
	claimNextCancellingJobEvents := make(
		chan *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
		100,
	)
	claimNextTopUpJobEvents := make(
		chan *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent,
		100,
	)
	sub := mocks.NewSubscription(suite.T())
	sub.EXPECT().Unsubscribe()
	suite.contractBackend.EXPECT().SubscribeFilterLogs(
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

func (suite *ClientTestSuite) TestClaimCancelling() {
	// Arrange
	suite.mustAuthenticate()
	suite.mockProviderJobQueuesContractCall(
		"hasCancellingJob",
		[]interface{}{fromAddress},
		[]interface{}{true},
	)
	suite.mockContractTransaction("claimNextCancellingJob")
	// Must wait
	suite.deployBackend.EXPECT().TransactionReceipt(mock.Anything, mock.Anything).Return(nil, nil)

	// Act
	err := suite.impl.ClaimCancelling(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *ClientTestSuite) TestClaimTopUp() {
	// Arrange
	suite.mustAuthenticate()
	suite.mockProviderJobQueuesContractCall(
		"hasTopUpJob",
		[]interface{}{fromAddress},
		[]interface{}{true},
	)
	suite.mockContractTransaction("claimNextTopUpJob")
	// Must wait
	suite.deployBackend.EXPECT().TransactionReceipt(mock.Anything, mock.Anything).Return(nil, nil)

	// Act
	err := suite.impl.ClaimTopUp(context.Background())

	// Assert
	suite.NoError(err)
	suite.assertMocksExpectations()
}

func (suite *ClientTestSuite) TestGetJobStatus() {
	// Arrange
	fixtureStatus := metascheduler.JobStatusRunning
	suite.mockContractCall("jobs", []interface{}{jobID}, []interface{}{
		jobID,
		uint8(fixtureStatus),
		common.HexToAddress("0xdeadface"),
		common.HexToAddress("0xdeadface"),
		metaschedulerabi.JobDefinition{},
		false,
		metaschedulerabi.JobCost{
			MaxCost:      new(big.Int),
			FinalCost:    new(big.Int),
			PendingTopUp: new(big.Int),
		},
		metaschedulerabi.JobTime{
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

func (suite *ClientTestSuite) TestClaimCancellingNoCancelling() {
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

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}
