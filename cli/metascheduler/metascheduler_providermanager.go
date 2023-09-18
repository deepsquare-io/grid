package metascheduler

import (
	"context"
	"fmt"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
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

func (c *providerManager) GetWaitingForApprovalProviders(
	ctx context.Context,
) (waiting []metaschedulerabi.Provider, notWaiting []metaschedulerabi.Provider, err error) {
	it, err := c.FilterProviderWaitingForApproval(&bind.FilterOpts{Context: ctx})
	if err != nil {
		return waiting, notWaiting, WrapError(err)
	}
	defer func() {
		_ = it.Close()
	}()

	for it.Next() {
		provider, err := c.GetWaitingForApprovalProvider(
			&bind.CallOpts{Context: ctx},
			it.Event.Addr,
		)
		if err != nil {
			return waiting, notWaiting, WrapError(err)
		}

		if (provider.Addr != common.Address{}) {
			waiting = append(waiting, provider)
		} else {
			notWaiting = append(notWaiting, provider)
		}
	}

	return waiting, notWaiting, nil
}
