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

	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
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
	}, address)
	return balance, WrapError(err)
}

func (c *creditManager) Transfer(ctx context.Context, to common.Address, amount *big.Int) error {
	opts, err := c.authOpts(ctx)
	if err != nil {
		return fmt.Errorf("failed to create auth options: %w", err)
	}
	tx, err := c.IERC20.Transfer(opts, to, amount)
	if err != nil {
		return WrapError(err)
	}
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return WrapError(err)
	}
	if receipt.Status != 1 {
		internallog.I.Error("transaction failed", zap.Any("receipt", receipt))
		return fmt.Errorf("transaction failed: %v", receipt.TxHash.String())
	}
	internallog.I.Debug("transfer", zap.Any("receipt", receipt))
	return nil
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
