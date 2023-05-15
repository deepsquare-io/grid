package eth_test

import (
	"context"
	"errors"
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
)

type ErrorTestSuite struct {
	suite.Suite
	contractAddress common.Address
	contract        *eth.ErrorContract
	backend         *backends.SimulatedBackend
}

func (suite *ErrorTestSuite) BeforeTest(suiteName, testName string) {
	// Genesis Account
	privateKey, err := crypto.GenerateKey()
	suite.Require().NoError(err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	suite.Require().NoError(err)

	// Create a simulated Ethereum backend
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}
	blockGasLimit := uint64(4712388)
	backend := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	suite.backend = backend

	// Deploy contract
	contractAddress, tx, contract, err := eth.DeployErrorContract(auth, backend)
	suite.Require().NoError(err)
	backend.Commit()

	_, err = bind.WaitDeployed(context.Background(), backend, tx)
	suite.Require().NoError(err)

	suite.contractAddress = contractAddress
	suite.contract = contract
}

func (suite *ErrorTestSuite) TestParseErrors() {
	tests := []struct {
		name     string
		arrange  []interface{}
		act      func(r []interface{}) error
		expected func(r []interface{}) error
	}{
		{
			name: "ParseDoubleEndedQueueEmpty",
			act: func(r []interface{}) error {
				return suite.contract.ThrowEmpty(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.DoubleEndedQueueEmpty{}
			},
		},
		{
			name: "ParseDoubleEndedQueueOutOfBounds",
			act: func(r []interface{}) error {
				return suite.contract.ThrowOutOfBounds(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.DoubleEndedQueueOutOfBounds{}
			},
		},
		{
			name: "ParseInsufficientFunds",
			arrange: []interface{}{
				big.NewInt(rand.Int63()),
				big.NewInt(rand.Int63()),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowInsufficientFunds(
					&bind.CallOpts{},
					r[0].(*big.Int),
					r[1].(*big.Int),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.InsufficientFunds{
					r[0].(*big.Int),
					r[1].(*big.Int),
				}
			},
		},
		{
			name: "ParseNoJob",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoJob(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.NoJob{}
			},
		},
		{
			name: "ParseInvalidJob",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidJob(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidJob{}
			},
		},
		{
			name: "ParseInvalidJobDefinition",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidJobDefinition(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidJobDefinition{}
			},
		},
		{
			name: "ParseJobHotStatusOnly",
			arrange: []interface{}{
				eth.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowJobHotStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(eth.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.JobHotStatusOnly{
					r[0].(eth.JobStatus),
				}
			},
		},
		{
			name: "ParseRunningScheduledStatusOnly",
			arrange: []interface{}{
				eth.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowRunningScheduledStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(eth.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.RunningScheduledStatusOnly{
					r[0].(eth.JobStatus),
				}
			},
		},
		{
			name: "ParseRunningColdStatusOnly",
			arrange: []interface{}{
				eth.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowMetaScheduledScheduledStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(eth.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.MetaScheduledScheduledStatusOnly{
					r[0].(eth.JobStatus),
				}
			},
		},
		{
			name: "ParseRunningColdStatusOnly",
			arrange: []interface{}{
				eth.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowRunningColdStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(eth.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.RunningColdStatusOnly{
					r[0].(eth.JobStatus),
				}
			},
		},
		{
			name: "ParseInvalidNNodes",
			arrange: []interface{}{
				big.NewInt(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidNNodes(
					&bind.CallOpts{},
					r[0].(*big.Int),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidNNodes{
					r[0].(*big.Int),
				}
			},
		},
		{
			name: "ParseInvalidNCpu",
			arrange: []interface{}{
				big.NewInt(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidNCpu(
					&bind.CallOpts{},
					r[0].(*big.Int),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidNCpu{
					r[0].(*big.Int),
				}
			},
		},
		{
			name: "ParseInvalidNMem",
			arrange: []interface{}{
				big.NewInt(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidNMem(
					&bind.CallOpts{},
					r[0].(*big.Int),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidNMem{
					r[0].(*big.Int),
				}
			},
		},
		{
			name: "ParseCustomerOnly",
			arrange: []interface{}{
				common.HexToAddress("0x1"),
				common.HexToAddress("0x2"),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowCustomerOnly(
					&bind.CallOpts{},
					r[0].(common.Address),
					r[1].(common.Address),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.CustomerOnly{
					r[0].(common.Address),
					r[1].(common.Address),
				}
			},
		},
		{
			name: "ParseJobProviderOnly",
			arrange: []interface{}{
				common.HexToAddress("0x1"),
				common.HexToAddress("0x2"),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowJobProviderOnly(
					&bind.CallOpts{},
					r[0].(common.Address),
					r[1].(common.Address),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.JobProviderOnly{
					r[0].(common.Address),
					r[1].(common.Address),
				}
			},
		},
		{
			name: "ParseJobProviderThisOnly",
			arrange: []interface{}{
				common.HexToAddress("0x1"),
				common.HexToAddress("0x2"),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowJobProviderThisOnly(
					&bind.CallOpts{},
					r[0].(common.Address),
					r[1].(common.Address),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.JobProviderThisOnly{
					r[0].(common.Address),
					r[1].(common.Address),
				}
			},
		},
		{
			name: "ParseOwnerOnly",
			arrange: []interface{}{
				common.HexToAddress("0x1"),
				common.HexToAddress("0x2"),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowOwnerOnly(
					&bind.CallOpts{},
					r[0].(common.Address),
					r[1].(common.Address),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.OwnerOnly{
					r[0].(common.Address),
					r[1].(common.Address),
				}
			},
		},
		{
			name:    "ParseCustomerMetaSchedulerProviderOnly",
			arrange: []interface{}{},
			act: func(r []interface{}) error {
				return suite.contract.ThrowCustomerMetaSchedulerProviderOnly(
					&bind.CallOpts{},
				)
			},
			expected: func(r []interface{}) error {
				return &eth.CustomerMetaSchedulerProviderOnly{}
			},
		},
		{
			name:    "ParseMetashedulerProviderOnly",
			arrange: []interface{}{},
			act: func(r []interface{}) error {
				return suite.contract.ThrowMetashedulerProviderOnly(
					&bind.CallOpts{},
				)
			},
			expected: func(r []interface{}) error {
				return &eth.MetashedulerProviderOnly{}
			},
		},
		{
			name: "ParseProviderAddrIsZero",
			act: func(r []interface{}) error {
				return suite.contract.ThrowProviderAddrIsZero(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.ProviderAddrIsZero{}
			},
		},
		{
			name: "ParseProviderNotJoined",
			act: func(r []interface{}) error {
				return suite.contract.ThrowProviderNotJoined(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.ProviderNotJoined{}
			},
		},
		{
			name: "ParseNoProvider",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoProvider(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.NoProvider{}
			},
		},
		{
			name: "ParseWaitingApprovalOnly",
			act: func(r []interface{}) error {
				return suite.contract.ThrowWaitingApprovalOnly(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.WaitingApprovalOnly{}
			},
		},
		{
			name: "ParseBanned",
			act: func(r []interface{}) error {
				return suite.contract.ThrowBanned(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.Banned{}
			},
		},
		{
			name: "ParseRemainingTimeAboveLimit",
			arrange: []interface{}{
				big.NewInt(rand.Int63()),
				big.NewInt(rand.Int63()),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowRemainingTimeAboveLimit(
					&bind.CallOpts{},
					r[0].(*big.Int),
					r[1].(*big.Int),
				)
			},
			expected: func(r []interface{}) error {
				return &eth.RemainingTimeAboveLimit{
					r[0].(*big.Int),
					r[1].(*big.Int),
				}
			},
		},
		{
			name: "ParseCreditAddrIsZero",
			act: func(r []interface{}) error {
				return suite.contract.ThrowCreditAddrIsZero(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.CreditAddrIsZero{}
			},
		},
		{
			name: "ParseNoSpendingAuthority",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoSpendingAuthority(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.NoSpendingAuthority{}
			},
		},
		{
			name: "ParseDivisionByZeroError",
			act: func(r []interface{}) error {
				return suite.contract.ThrowDivisionByZeroError(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.DivisionByZeroError{}
			},
		},
		{
			name: "ParseUninitialized",
			act: func(r []interface{}) error {
				return suite.contract.ThrowUninitialized(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.Uninitialized{}
			},
		},
		{
			name: "ParseSameStatusError",
			act: func(r []interface{}) error {
				return suite.contract.ThrowSameStatusError(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.SameStatusError{}
			},
		},
		{
			name: "ParseInvalidTransitionFromPending",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidTransitionFromPending(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidTransitionFromPending{}
			},
		},
		{
			name: "ParseInvalidTransitionFromMetascheduled",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidTransitionFromMetascheduled(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidTransitionFromMetascheduled{}
			},
		},
		{
			name: "ParseInvalidTransitionFromScheduled",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidTransitionFromScheduled(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidTransitionFromScheduled{}
			},
		},
		{
			name: "ParseInvalidTransitionFromRunning",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidTransitionFromRunning(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &eth.InvalidTransitionFromRunning{}
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			// Arrange
			var r []interface{}
			if test.arrange == nil {
				r = []interface{}{}
			} else {
				r = test.arrange
			}

			// Act
			err := test.act(r)

			// Assert
			suite.Error(err)
			err = eth.WrapError(err)
			out := reflect.New(reflect.TypeOf(test.expected(r)).Elem()).Interface()
			ok := errors.As(err, &out)
			suite.NotEmpty(err.Error())
			suite.True(ok)
			suite.Equal(test.expected(r), out)
		})
	}
}

func (suite *ErrorTestSuite) AfterTest(suiteName, testName string) {
	_ = suite.backend.Close()
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, &ErrorTestSuite{})
}
