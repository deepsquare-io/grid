// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	"github.com/deepsquare-io/grid/cli/sbatch"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/deepsquare-io/grid/cli/types/allowance"
	"github.com/deepsquare-io/grid/cli/types/credit"
	"github.com/deepsquare-io/grid/cli/types/job"
	"github.com/deepsquare-io/grid/cli/types/provider"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// RPCClientSet is a set of clients that interact with DeepSquare.
type RPCClientSet struct {
	*Backend
}

// authOpts generate transact options based on the network.
func (c *RPCClientSet) transact(
	ctx context.Context,
	exec func(auth *bind.TransactOpts) (*types.Transaction, error),
) (*types.Transaction, error) {
	nonce, err := c.PendingNonceAt(ctx, c.from())
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.UserPrivateKey, c.ChainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	// TODO: remove hack
	// Conflict with GoQuorum (Paris), Avax (Shanghai) and SKALED (Istanbul)
	auth.GasLimit = uint64(0x1312D00)
	auth.GasPrice = gasPrice
	auth.Context = ctx

	simulated := &bind.TransactOpts{
		From:      auth.From,
		Signer:    auth.Signer,
		Nonce:     auth.Nonce,
		Value:     auth.Value,
		GasPrice:  auth.GasPrice,
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
		GasLimit:  auth.GasLimit,
		Context:   auth.Context,
		NoSend:    true,
	}

	// Simuate the transaction
	tx, err := exec(simulated)
	if err != nil {
		return nil, err
	}

	// Play fake transaction to find error reason
	_, err = c.EstimateGas(ctx, ethereum.CallMsg{
		To:         tx.To(),
		From:       auth.From,
		Gas:        tx.Gas(),
		GasPrice:   tx.GasPrice(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	})
	if err != nil {
		return nil, err
	}

	return exec(auth)
}

// NewRPCClientSet creates an RPCClientSet.
func NewRPCClientSet(b Backend) *RPCClientSet {
	return &RPCClientSet{
		Backend: &b,
	}
}

// JobScheduler creates a [types.JobScheduler].
func (c *RPCClientSet) JobScheduler(
	sbatch sbatch.Service,
) job.Scheduler {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	return &jobScheduler{
		RPCClientSet:  c,
		MetaScheduler: m,
		Service:       sbatch,
	}
}

// JobFetcher creates a [job.Fetcher].
func (c *RPCClientSet) JobFetcher() job.Fetcher {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	jobsAddress, err := m.Jobs(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch JobRepository contract address: %w", err))
	}
	jobs, err := metaschedulerabi.NewIJobRepository(jobsAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate JobRepository: %w", err))
	}

	return &jobFetcher{
		RPCClientSet:   c,
		IJobRepository: jobs,
	}
}

// CreditManager creates a [credit.Manager].
func (c *RPCClientSet) CreditManager() credit.Manager {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	creditAddress, err := m.Credit(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch Credit contract address: %w", err))
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate Credit: %w", err))
	}
	return &creditManager{
		RPCClientSet: c,
		IERC20:       ierc20,
	}
}

// AllowanceManager creates an [allowance.Manager].
func (c *RPCClientSet) AllowanceManager() allowance.Manager {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	creditAddress, err := m.Credit(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch CreditManager contract address: %w", err))
	}
	ierc20, err := metaschedulerabi.NewIERC20(creditAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate CreditManager: %w", err))
	}
	return &allowanceManager{
		RPCClientSet: c,
		IERC20:       ierc20,
	}
}

// ProviderManager creates a [provider.Manager].
func (c *RPCClientSet) ProviderManager() provider.Manager {
	m, err := metaschedulerabi.NewMetaScheduler(c.MetaschedulerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	providerManagerAddress, err := m.ProviderManager(&bind.CallOpts{})
	if err != nil {
		panic(fmt.Errorf("failed to fetch CreditManager contract address: %w", err))
	}
	pm, err := metaschedulerabi.NewIProviderManager(providerManagerAddress, c)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate CreditManager: %w", err))
	}
	return &providerManager{
		RPCClientSet:     c,
		IProviderManager: pm,
	}
}
