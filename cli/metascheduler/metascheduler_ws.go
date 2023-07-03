package metascheduler

import (
	"context"
	"fmt"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/cli/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type wsClient struct {
	*metaschedulerabi.MetaScheduler
	Backend
}

func (c *wsClient) from() (addr common.Address) {
	if c.UserPrivateKey == nil {
		return addr
	}
	return crypto.PubkeyToAddress(c.UserPrivateKey.PublicKey)
}

func (c *wsClient) SubscribeEvents(
	ctx context.Context,
	ch chan<- ethtypes.Log,
) (ethereum.Subscription, error) {
	creditAddress, err := c.Credit(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, WrapError(err)
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			c.MetaschedulerAddress,
			creditAddress,
		},
		Topics: [][]common.Hash{
			{
				newJobRequestEvent.ID,
				jobTransitionEvent.ID,
				transferEvent.ID,
				approvalEvent.ID,
			},
		},
	}

	sub, err := c.SubscribeFilterLogs(ctx, query, ch)
	return sub, WrapError(err)
}

func (c *wsClient) FilterNewJobRequests(
	ch <-chan ethtypes.Log,
) (filtered <-chan *metaschedulerabi.MetaSchedulerNewJobRequestEvent, rest <-chan ethtypes.Log) {
	fChan := make(chan *metaschedulerabi.MetaSchedulerNewJobRequestEvent)
	rChan := make(chan ethtypes.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == newJobRequestEvent.ID.Hex() {
				event, err := c.ParseNewJobRequestEvent(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

func (c *wsClient) FilterJobTransition(
	ch <-chan ethtypes.Log,
) (filtered <-chan *metaschedulerabi.MetaSchedulerJobTransitionEvent, rest <-chan ethtypes.Log) {
	fChan := make(chan *metaschedulerabi.MetaSchedulerJobTransitionEvent)
	rChan := make(chan ethtypes.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == jobTransitionEvent.ID.Hex() {
				event, err := c.ParseJobTransitionEvent(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

type creditFilterer struct {
	types.CreditManager
	wsClient
	*metaschedulerabi.IERC20Filterer
}

func (c *creditFilterer) FilterTransfer(
	ctx context.Context,
	ch <-chan ethtypes.Log,
) (filtered <-chan *metaschedulerabi.IERC20Transfer, rest <-chan ethtypes.Log) {
	fChan := make(chan *metaschedulerabi.IERC20Transfer)
	rChan := make(chan ethtypes.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == transferEvent.ID.Hex() {
				event, err := c.ParseTransfer(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

func (c *creditFilterer) ReduceToBalance(
	ctx context.Context,
	transfers <-chan *metaschedulerabi.IERC20Transfer,
) (<-chan *big.Int, error) {
	rChan := make(chan *big.Int, 2)
	errChan := make(chan error, 1)

	// Fetch initial value
	value, err := c.CreditManager.Balance(ctx)
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

type allowanceFilterer struct {
	types.AllowanceManager
	wsClient
	*metaschedulerabi.IERC20Filterer
}

func (c *allowanceFilterer) FilterApproval(
	ctx context.Context,
	ch <-chan ethtypes.Log,
) (filtered <-chan *metaschedulerabi.IERC20Approval, rest <-chan ethtypes.Log) {
	fChan := make(chan *metaschedulerabi.IERC20Approval)
	rChan := make(chan ethtypes.Log)

	go func() {
		defer close(fChan)
		defer close(rChan)
		for log := range ch {
			if len(log.Topics) == 0 {
				return
			}
			if log.Topics[0].Hex() == transferEvent.ID.Hex() {
				event, err := c.ParseApproval(log)
				if err != nil {
					panic(fmt.Errorf("failed to parse event: %w", err))
				}

				fChan <- event
			} else {
				rChan <- log
			}
		}
	}()

	return fChan, rChan
}

func (c *allowanceFilterer) ReduceToAllowance(
	ctx context.Context,
	approvals <-chan *metaschedulerabi.IERC20Approval,
) (<-chan *big.Int, error) {
	rChan := make(chan *big.Int, 2)
	errChan := make(chan error, 1)

	// Fetch initial value
	value, err := c.AllowanceManager.GetAllowance(ctx)
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
