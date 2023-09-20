package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type allowanceManager struct {
	*RPCClientSet
	*metaschedulerabi.IERC20
}

func (c *allowanceManager) SetAllowance(ctx context.Context, amount *big.Int) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed get auth options: %w", err)
	}
	tx, err := c.Approve(opts, c.MetaschedulerAddress, amount)
	if err != nil {
		return fmt.Errorf("failed to approve credit: %w", err)
	}
	_, err = bind.WaitMined(ctx, c, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}
	return nil
}

func (c *allowanceManager) ClearAllowance(ctx context.Context) error {
	return c.SetAllowance(ctx, big.NewInt(0))
}

func (c *allowanceManager) GetAllowance(ctx context.Context) (*big.Int, error) {
	return c.Allowance(&bind.CallOpts{
		Context: ctx,
	}, c.from(), c.MetaschedulerAddress)
}

func (c *allowanceManager) ReduceToAllowance(
	ctx context.Context,
	approvals <-chan types.Approval,
) (<-chan *big.Int, error) {
	rChan := make(chan *big.Int, 2)
	errChan := make(chan error, 1)

	// Fetch initial value
	value, err := c.GetAllowance(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(rChan)
		defer close(errChan)

		// Send initial value
		rChan <- value

		// Track value
		for approval := range approvals {
			if approval.Owner == c.from() && approval.Spender == c.MetaschedulerAddress {
				rChan <- approval.Value
			}
		}
	}()

	return rChan, nil
}
