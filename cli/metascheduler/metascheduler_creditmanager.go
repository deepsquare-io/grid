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
	"github.com/deepsquare-io/grid/cli/types/credit"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

var _ credit.Manager = (*CreditManager)(nil)

// CreditManager is a manager for credits.
type CreditManager struct {
	*RPCClientSet
	*metaschedulerabi.IERC20
}

// Balance returns the balance of the user.
func (c *CreditManager) Balance(ctx context.Context) (*big.Int, error) {
	return c.BalanceOf(ctx, c.from())
}

// BalanceOf returns the balance of the given address.
func (c *CreditManager) BalanceOf(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := c.IERC20.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, address)
	return balance, WrapError(err)
}

// Transfer transfers the given amount to the given address.
func (c *CreditManager) Transfer(ctx context.Context, to common.Address, amount *big.Int) error {
	tx, err := c.transact(ctx, func(auth *bind.TransactOpts) (*coretypes.Transaction, error) {
		return c.IERC20.Transfer(auth, to, amount)
	})
	if err != nil {
		return WrapError(err)
	}
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return WrapError(err)
	}
	internallog.I.Debug("transfer", zap.Any("receipt", receipt))
	return CheckReceiptError(ctx, c, tx, receipt)
}

// ReduceToBalance reduces the balance to the given value.
func (c *CreditManager) ReduceToBalance(
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
