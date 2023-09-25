package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	"github.com/deepsquare-io/the-grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type creditManager struct {
	*RPCClientSet
	*metaschedulerabi.IERC20
}

func (c *creditManager) Balance(ctx context.Context) (*big.Int, error) {
	return c.BalanceOf(ctx, c.from())
}

func (c *creditManager) BalanceOf(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := c.IERC20.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, c.from())
	return balance, WrapError(err)
}

func (c *creditManager) Transfer(ctx context.Context, to common.Address, amount *big.Int) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create auth options: %w", err)
	}
	tx, err := c.IERC20.Transfer(opts, c.from(), amount)
	if err != nil {
		return WrapError(err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	return WrapError(err)
}
func (c *creditManager) ReduceToBalance(
	ctx context.Context,
	transfers <-chan types.Transfer,
) (<-chan *big.Int, error) {
	rChan := make(chan *big.Int, 2)
	errChan := make(chan error, 1)

	// Fetch initial value
	value, err := c.Balance(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(rChan)
		defer close(errChan)

		// Send initial value
		rChan <- value

		// Track value
		for transfer := range transfers {
			// User is sending data
			if c.from() == transfer.From {
				value = new(big.Int).Sub(value, transfer.Value)
				rChan <- value
			} else if c.from() == transfer.To {
				value = new(big.Int).Add(value, transfer.Value)
				rChan <- value
			}
		}
	}()

	return rChan, nil
}
