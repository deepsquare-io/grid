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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// RPCClientSet is a set of clients that interact with DeepSquare.
type RPCClientSet struct {
	*Backend
}

// authOpts generate transact options based on the network.
func (c *RPCClientSet) authOpts(ctx context.Context) (*bind.TransactOpts, error) {
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
	auth.GasLimit = uint64(0)
	auth.GasPrice = gasPrice
	auth.Context = ctx

	return auth, nil
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
