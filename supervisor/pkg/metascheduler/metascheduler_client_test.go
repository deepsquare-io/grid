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

//go:build unit

package metascheduler_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/deepsquare-io/grid/supervisor/mocks/mockbind"
	"github.com/deepsquare-io/grid/supervisor/mocks/mockethereum"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	jobRepositoryAddress     = common.HexToAddress("0x3")
	providerManagerAddress   = common.HexToAddress("0x4")
)

func (suite *ClientTestSuite) mockProviderJobQueues() {
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

func (suite *ClientTestSuite) mockProviderManager() {
	// Pack input
	input, err := metascheduler.MetaschedulerABI.Pack("providerManager")
	suite.Require().NoError(err)
	output, err := metascheduler.MetaschedulerABI.Methods["providerManager"].Outputs.Pack(
		providerManagerAddress,
	)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.EXPECT().CallContract(mock.Anything, ethereum.CallMsg{
		To:   &metaschedulerAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *ClientTestSuite) mockIJobRepository() {
	// Pack input
	input, err := metascheduler.MetaschedulerABI.Pack("jobs")
	suite.Require().NoError(err)
	output, err := metascheduler.MetaschedulerABI.Methods["jobs"].Outputs.Pack(
		jobRepositoryAddress,
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

func (suite *ClientTestSuite) mockProviderManagerContractCall(
	name string,
	inputs []interface{},
	outputs []interface{},
) {
	// Pack input
	abi, err := metaschedulerabi.IProviderManagerMetaData.GetAbi()
	suite.Require().NoError(err)
	input, err := abi.Pack(name, inputs...)
	suite.Require().NoError(err)
	output, err := abi.Methods[name].Outputs.Pack(outputs...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.EXPECT().CallContract(mock.Anything, ethereum.CallMsg{
		To:   &providerManagerAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *ClientTestSuite) mockProviderManagerContractTransaction(
	name string,
	args ...interface{},
) {
	suite.contractBackend.EXPECT().
		EstimateGas(mock.Anything, mock.Anything).
		Return(uint64(1), nil).
		Once()
	suite.contractBackend.EXPECT().SendTransaction(mock.Anything, mock.Anything).Return(nil)
}

func (suite *ClientTestSuite) mockMetaschedulerContractCall(
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

func (suite *ClientTestSuite) mockJobRepositoryContractCall(
	name string,
	inputs []interface{},
	outputs []interface{},
) {
	// Pack input
	abi, err := metaschedulerabi.IJobRepositoryMetaData.GetAbi()
	suite.Require().NoError(err)
	input, err := abi.Pack(name, inputs...)
	suite.Require().NoError(err)
	output, err := abi.Methods[name].Outputs.Pack(outputs...)
	suite.Require().NoError(err)

	// Mock
	suite.contractBackend.EXPECT().CallContract(mock.Anything, ethereum.CallMsg{
		To:   &jobRepositoryAddress,
		Data: input,
	}, mock.Anything).Return(output, nil)
}

func (suite *ClientTestSuite) mockContractTransaction(name string, args ...interface{}) {
	suite.contractBackend.EXPECT().
		EstimateGas(mock.Anything, mock.Anything).
		Return(uint64(1), nil).
		Once()
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
	suite.mockProviderJobQueues()
	// Assert calling jobs
	suite.mockIJobRepository()
	// Assert calling providermanager
	suite.mockProviderManager()

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
	suite.deployBackend.EXPECT().
		TransactionReceipt(mock.Anything, mock.Anything).
		Return(&types.Receipt{
			Status: 1,
		}, nil)

	// Act
	err := suite.impl.Claim(context.Background())

	// Assert
	suite.NoError(err)
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
}

func (suite *ClientTestSuite) TestRegister() {
	// Arrange
	hardware := metaschedulerabi.ProviderHardware{
		Nodes:       1,
		GpusPerNode: []uint64{2},
		CpusPerNode: []uint64{3},
		MemPerNode:  []uint64{4},
	}
	prices := metaschedulerabi.ProviderPrices{
		GpuPricePerMin: big.NewInt(1),
		CpuPricePerMin: big.NewInt(2),
		MemPricePerMin: big.NewInt(3),
	}
	labels := []metaschedulerabi.Label{
		{
			Key:   "key",
			Value: "value",
		},
	}
	suite.mustAuthenticate()
	suite.mockProviderManagerContractTransaction(
		"register",
		hardware,
		prices,
		labels,
	)
	// Must wait
	suite.deployBackend.EXPECT().
		TransactionReceipt(mock.Anything, mock.Anything).
		Return(&types.Receipt{
			Status: 1,
		}, nil)

	// Act
	err := suite.impl.Register(context.Background(), hardware, prices, labels)

	// Assert
	suite.NoError(err)
}

func (suite *ClientTestSuite) TestSetJobStatus() {
	// Arrange
	suite.mustAuthenticate()
	suite.mockContractTransaction(
		"providerSetJobStatus",
		jobID,
		uint8(metascheduler.JobStatusFailed),
		jobDuration,
		"",
		int64(0),
	)
	// Must wait
	suite.deployBackend.EXPECT().
		TransactionReceipt(mock.Anything, mock.Anything).
		Return(&types.Receipt{
			Status: 1,
		}, nil)

	// Act
	err := suite.impl.SetJobStatus(
		context.Background(),
		jobID,
		metascheduler.JobStatusFailed,
		jobDuration,
	)

	// Assert
	suite.NoError(err)
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
	sub := mockethereum.NewSubscription(suite.T())
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
	suite.deployBackend.EXPECT().
		TransactionReceipt(mock.Anything, mock.Anything).
		Return(&types.Receipt{
			Status: 1,
		}, nil)

	// Act
	err := suite.impl.ClaimCancelling(context.Background())

	// Assert
	suite.NoError(err)
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
	suite.deployBackend.EXPECT().
		TransactionReceipt(mock.Anything, mock.Anything).
		Return(&types.Receipt{
			Status: 1,
		}, nil)

	// Act
	err := suite.impl.ClaimTopUp(context.Background())

	// Assert
	suite.NoError(err)
}

func (suite *ClientTestSuite) TestGetJobStatus() {
	// Arrange
	fixtureStatus := metascheduler.JobStatusRunning
	suite.mockJobRepositoryContractCall(
		"get",
		[]interface{}{jobID},
		[]interface{}{
			metaschedulerabi.Job{
				JobId:        jobID,
				Status:       uint8(fixtureStatus),
				CustomerAddr: common.HexToAddress("0xdeadface"),
				ProviderAddr: common.HexToAddress("0xdeadface"),
				Definition:   metaschedulerabi.JobDefinition{},
				Cost: metaschedulerabi.JobCost{
					MaxCost:      new(big.Int),
					FinalCost:    new(big.Int),
					PendingTopUp: new(big.Int),
				},
				Time: metaschedulerabi.JobTime{
					Submit:                 new(big.Int),
					Start:                  new(big.Int),
					End:                    new(big.Int),
					CancelRequestTimestamp: new(big.Int),
					BlockNumberStateChange: new(big.Int),
					PanicTimestamp:         new(big.Int),
				},
				JobName:          [32]byte{1},
				HasCancelRequest: false,
				LastError:        "",
				ExitCode:         0,
			},
		},
	)
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
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}
