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
	"math/big"

	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/deepsquare-io/grid/cli/types/allowance"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

var _ allowance.Manager = (*AllowanceManager)(nil)

// AllowanceManager is a manager for allowances.
type AllowanceManager struct {
	*RPCClientSet
	*metaschedulerabi.IERC20
}

// SetAllowance sets the allowance for the metascheduler.
func (c *AllowanceManager) SetAllowance(ctx context.Context, amount *big.Int) error {
	tx, err := c.transact(ctx, func(auth *bind.TransactOpts) (*coretypes.Transaction, error) {
		return c.Approve(auth, c.MetaschedulerAddress, amount)
	})
	if err != nil {
		return WrapError(err)
	}
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return WrapError(err)
	}
	internallog.I.Debug("metascheduled job", zap.Any("receipt", receipt))
	return CheckReceiptError(ctx, c, tx, receipt)
}

// ClearAllowance clears the allowance for the metascheduler.
func (c *AllowanceManager) ClearAllowance(ctx context.Context) error {
	return c.SetAllowance(ctx, big.NewInt(0))
}

// GetAllowance gets the allowance for the metascheduler.
func (c *AllowanceManager) GetAllowance(ctx context.Context) (*big.Int, error) {
	return c.Allowance(&bind.CallOpts{
		Context: ctx,
	}, c.from(), c.MetaschedulerAddress)
}

// ReduceToAllowance reduces the stream of allowance to the allowance.
func (c *AllowanceManager) ReduceToAllowance(
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
