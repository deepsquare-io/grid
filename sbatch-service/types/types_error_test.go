// Copyright (C) 2024 DeepSquare Association
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

package types_test

import (
	"context"
	"errors"
	"math/big"
	"math/rand"
	"reflect"
	"testing"

	errorabi "github.com/deepsquare-io/grid/sbatch-service/abi/error"
	"github.com/deepsquare-io/grid/sbatch-service/types"
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
	contract        *errorabi.ErrorContract
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
	contractAddress, tx, contract, err := errorabi.DeployErrorContract(auth, backend)
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
				return &types.DoubleEndedQueueEmpty{}
			},
		},
		{
			name: "ParseDoubleEndedQueueOutOfBounds",
			act: func(r []interface{}) error {
				return suite.contract.ThrowOutOfBounds(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.DoubleEndedQueueOutOfBounds{}
			},
		},
		{
			name: "ParseInvalidJob",
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidJob(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.InvalidJob{}
			},
		},
		{
			name: "ParseNoJob",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoJob(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.NoJob{}
			},
		},
		{
			name: "ParseNoProvider",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNoProvider(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.NoProvider{}
			},
		},
		{
			name: "ParseWaitingApprovalOnly",
			act: func(r []interface{}) error {
				return suite.contract.ThrowWaitingApprovalOnly(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.WaitingApprovalOnly{}
			},
		},
		{
			name: "ParseBanned",
			act: func(r []interface{}) error {
				return suite.contract.ThrowBanned(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.Banned{}
			},
		},
		{
			name: "ParseAlreadyDone",
			act: func(r []interface{}) error {
				return suite.contract.ThrowAlreadyDone(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.AlreadyDone{}
			},
		},
		{
			name: "ParseJobHotStatusOnly",
			arrange: []interface{}{
				types.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowJobHotStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(types.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &types.JobHotStatusOnly{
					r[0].(types.JobStatus),
				}
			},
		},
		{
			name: "ParseInvalidTransition",
			arrange: []interface{}{
				types.JobStatus(1),
				types.JobStatus(2),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowInvalidTransition(
					&bind.CallOpts{},
					uint8(r[0].(types.JobStatus)),
					uint8(r[1].(types.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &types.InvalidTransition{
					r[0].(types.JobStatus),
					r[1].(types.JobStatus),
				}
			},
		},
		{
			name: "ParseSameStatusError",
			act: func(r []interface{}) error {
				return suite.contract.ThrowSameStatusError(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.SameStatusError{}
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
				return &types.InsufficientFunds{
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
				return &types.InvalidJobDefinition{}
			},
		},
		{
			name: "ParseRunningScheduledStatusOnly",
			arrange: []interface{}{
				types.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowRunningScheduledStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(types.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &types.RunningScheduledStatusOnly{
					r[0].(types.JobStatus),
				}
			},
		},
		{
			name: "ParseMetaScheduledScheduledStatusOnly",
			arrange: []interface{}{
				types.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowMetaScheduledScheduledStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(types.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &types.MetaScheduledScheduledStatusOnly{
					r[0].(types.JobStatus),
				}
			},
		},
		{
			name: "ParseRunningColdStatusOnly",
			arrange: []interface{}{
				types.JobStatus(1),
			},
			act: func(r []interface{}) error {
				return suite.contract.ThrowRunningColdStatusOnly(
					&bind.CallOpts{},
					uint8(r[0].(types.JobStatus)),
				)
			},
			expected: func(r []interface{}) error {
				return &types.RunningColdStatusOnly{
					r[0].(types.JobStatus),
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
				return &types.CustomerOnly{
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
				return &types.JobProviderOnly{
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
				return &types.CustomerMetaSchedulerProviderOnly{}
			},
		},
		{
			name: "ParseProviderNotJoined",
			act: func(r []interface{}) error {
				return suite.contract.ThrowProviderNotJoined(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.ProviderNotJoined{}
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
				return &types.RemainingTimeAboveLimit{
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
				return &types.NoSpendingAuthority{}
			},
		},
		{
			name: "ParseNewJobRequestDisabled",
			act: func(r []interface{}) error {
				return suite.contract.ThrowNewJobRequestDisabled(&bind.CallOpts{})
			},
			expected: func(r []interface{}) error {
				return &types.NewJobRequestDisabled{}
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
			err = types.WrapError(err)
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
