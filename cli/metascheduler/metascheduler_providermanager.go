package metascheduler

import (
	"context"
	"fmt"

	"github.com/deepsquare-io/the-grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
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
) (provider types.ProviderDetail, err error) {
	p, err := c.GetWaitingForApprovalProvider(
		&bind.CallOpts{Context: ctx},
		address,
	)
	if err != nil {
		return provider, WrapError(err)
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

	provider.Addr = address

	return types.ProviderDetail{
		Provider:             p,
		IsWaitingForApproval: isWaitingForApproval,
		IsValidForScheduling: isValidForScheduling,
		JobCount:             jobCount,
	}, nil
}

func (c *providerManager) GetProviders(
	ctx context.Context,
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
		provider, err := c.GetWaitingForApprovalProvider(
			&bind.CallOpts{Context: ctx},
			it.Event.Addr,
		)
		if err != nil {
			return providers, WrapError(err)
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
