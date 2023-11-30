// Copyright (C) 2023 DeepSquare Association
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

	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type providerManager struct {
	*RPCClientSet
	*metaschedulerabi.IProviderManager
}

func (c *providerManager) Approve(ctx context.Context, provider common.Address) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed get auth options: %w", err)
	}
	tx, err := c.IProviderManager.Approve(opts, provider)
	if err != nil {
		return fmt.Errorf("failed to approve provider: %w", err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}
	return nil
}

func (c *providerManager) Remove(ctx context.Context, provider common.Address) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed get auth options: %w", err)
	}
	tx, err := c.IProviderManager.Remove(opts, provider)
	if err != nil {
		return fmt.Errorf("failed to remove provider: %w", err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}
	return nil
}

func (c *providerManager) GetProvider(
	ctx context.Context,
	address common.Address,
	opts ...types.GetProviderOption,
) (provider types.ProviderDetail, err error) {
	var o types.GetProviderOptions
	for _, opt := range opts {
		opt(&o)
	}
	var p metaschedulerabi.Provider
	if o.Proposal {
		p, err = c.GetWaitingForApprovalProvider(
			&bind.CallOpts{Context: ctx},
			address,
		)
		if err != nil {
			return provider, WrapError(err)
		}
	} else {
		p, err = c.IProviderManager.GetProvider(
			&bind.CallOpts{Context: ctx},
			address,
		)
		if err != nil {
			return provider, WrapError(err)
		}
	}

	isWaitingForApproval, err := c.IsWaitingForApproval(
		&bind.CallOpts{Context: ctx},
		address,
	)
	if err != nil {
		return provider, WrapError(err)
	}

	isValidForScheduling, err := c.IsValidForScheduling(
		&bind.CallOpts{Context: ctx},
		address,
	)
	if err != nil {
		return provider, WrapError(err)
	}

	jobCount, err := c.GetJobCount(
		&bind.CallOpts{Context: ctx},
		address,
	)
	if err != nil {
		return provider, WrapError(err)
	}

	p.Addr = address

	return types.ProviderDetail{
		Provider:             p,
		IsWaitingForApproval: isWaitingForApproval,
		IsValidForScheduling: isValidForScheduling,
		JobCount:             jobCount,
	}, nil
}

func (c *providerManager) GetProviders(
	ctx context.Context,
	opts ...types.GetProviderOption,
) (providers []types.ProviderDetail, err error) {
	it, err := c.FilterProviderWaitingForApproval(&bind.FilterOpts{Context: ctx})
	if err != nil {
		return providers, WrapError(err)
	}
	defer func() {
		_ = it.Close()
	}()

	providerMap := make(map[common.Address]types.ProviderDetail)

	for it.Next() {
		var o types.GetProviderOptions
		for _, opt := range opts {
			opt(&o)
		}
		var provider metaschedulerabi.Provider
		if o.Proposal {
			provider, err = c.GetWaitingForApprovalProvider(
				&bind.CallOpts{Context: ctx},
				it.Event.Addr,
			)
			if err != nil {
				return providers, WrapError(err)
			}
		} else {
			provider, err = c.IProviderManager.GetProvider(
				&bind.CallOpts{Context: ctx},
				it.Event.Addr,
			)
			if err != nil {
				return providers, WrapError(err)
			}
		}

		isWaitingForApproval, err := c.IsWaitingForApproval(
			&bind.CallOpts{Context: ctx},
			it.Event.Addr,
		)
		if err != nil {
			return providers, WrapError(err)
		}

		isValidForScheduling, err := c.IsValidForScheduling(
			&bind.CallOpts{Context: ctx},
			it.Event.Addr,
		)
		if err != nil {
			return providers, WrapError(err)
		}

		jobCount, err := c.GetJobCount(
			&bind.CallOpts{Context: ctx},
			it.Event.Addr,
		)
		if err != nil {
			return providers, WrapError(err)
		}

		provider.Addr = it.Event.Addr

		providerMap[it.Event.Addr] = types.ProviderDetail{
			Provider:             provider,
			IsWaitingForApproval: isWaitingForApproval,
			IsValidForScheduling: isValidForScheduling,
			JobCount:             jobCount,
		}
	}

	providers = make([]types.ProviderDetail, 0, len(providerMap))
	for _, v := range providerMap {
		providers = append(providers, v)
	}

	return providers, nil
}
