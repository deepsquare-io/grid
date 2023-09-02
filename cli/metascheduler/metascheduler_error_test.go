package metascheduler_test

import (
	"context"
	"errors"
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	errorsabi "github.com/deepsquare-io/the-grid/cli/internal/abi/errors"
	"github.com/deepsquare-io/the-grid/cli/metascheduler"
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
	contract        *errorsabi.ErrorContract
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
	balance.SetString("10000000000000000000", 10) // 10 metascheduler in wei

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
	contractAddress, tx, contract, err := errorsabi.DeployErrorContract(auth, backend)
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
				return &metascheduler.DoubleEndedQueueEmpty{}
			},
		},
		{
			name: "ParseDoubleEndedQueueOutOfBounds",
			act: func(r []interface{}) error {
				return suite.contract.ThrowOutOfBounds(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.DoubleEndedQueueOutOfBounds{}
			},
		},
		{
			name: "ParseInvalidJob",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidJob(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.InvalidJob{}
			},
		},
		{
			name: "ParseNoJob",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoJob(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.NoJob{}
			},
		},
		{
			name: "ParseNoProvider",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoProvider(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.NoProvider{}
			},
		},
		{
			name: "ParseWaitingApprovalOnly",
			act: func(r []interface{}) error {
				return suite.contract.ThrowWaitingApprovalOnly(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.WaitingApprovalOnly{}
			},
		},
		{
			name: "ParseBanned",
			act: func(r []interface{}) error {
				return suite.contract.ThrowBanned(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.Banned{}
			},
		},
		{
			name: "ParseJobHotStatusOnly",
			arrange: []interface{}{
				metascheduler.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowJobHotStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(metascheduler.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &metascheduler.JobHotStatusOnly{
					r[0].(metascheduler.JobStatus),
				}
			},
		},
		{
			name: "ParseInvalidTransition",
			arrange: []interface{}{
				metascheduler.JobStatus(1),
				metascheduler.JobStatus(2),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidTransition(
					&bind.CallOpts{},
					uint8(r[0].(metascheduler.JobStatus)),
					uint8(r[1].(metascheduler.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &metascheduler.InvalidTransition{
					r[0].(metascheduler.JobStatus),
					r[1].(metascheduler.JobStatus),
				}
			},
		},
		{
			name: "ParseSameStatusError",
			act: func(r []interface{}) error {
				return suite.contract.ThrowSameStatusError(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.SameStatusError{}
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
				return &metascheduler.InsufficientFunds{
					r[0].(*big.Int),
					r[1].(*big.Int),
				}
			},
		},
		{
			name: "ParseInvalidJobDefinition",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidJobDefinition(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.InvalidJobDefinition{}
			},
		},
		{
			name: "ParseRunningScheduledStatusOnly",
			arrange: []interface{}{
				metascheduler.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowRunningScheduledStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(metascheduler.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &metascheduler.RunningScheduledStatusOnly{
					r[0].(metascheduler.JobStatus),
				}
			},
		},
		{
			name: "ParseMetaScheduledScheduledStatusOnly",
			arrange: []interface{}{
				metascheduler.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowMetaScheduledScheduledStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(metascheduler.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &metascheduler.MetaScheduledScheduledStatusOnly{
					r[0].(metascheduler.JobStatus),
				}
			},
		},
		{
			name: "ParseRunningColdStatusOnly",
			arrange: []interface{}{
				metascheduler.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowRunningColdStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(metascheduler.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &metascheduler.RunningColdStatusOnly{
					r[0].(metascheduler.JobStatus),
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
				return &metascheduler.CustomerOnly{
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
				return &metascheduler.JobProviderOnly{
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
				return &metascheduler.CustomerMetaSchedulerProviderOnly{}
			},
		},
		{
			name: "ParseProviderNotJoined",
			act: func(r []interface{}) error {
				return suite.contract.ThrowProviderNotJoined(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.ProviderNotJoined{}
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
				return &metascheduler.RemainingTimeAboveLimit{
					r[0].(*big.Int),
					r[1].(*big.Int),
				}
			},
		},
		{
			name: "ParseNoSpendingAuthority",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoSpendingAuthority(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.NoSpendingAuthority{}
			},
		},
		{
			name: "ParseNewJobRequestDisabled",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNewJobRequestDisabled(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &metascheduler.NewJobRequestDisabled{}
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
			err = metascheduler.WrapError(err)
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
