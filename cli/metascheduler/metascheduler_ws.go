// Copyright (C) 2023 DeepSquare Asociation
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package metascheduler

import (
	"context"
	"fmt"

	"github.com/deepsquare-io/the-grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type eventSubscriber struct {
	rpc Backend
	ws  Backend

	rpcMetascheduler *metaschedulerabi.MetaScheduler
}

func NewEventSubscriber(
	rpc Backend,
	ws Backend,
) types.EventSubscriber {
	m, err := metaschedulerabi.NewMetaScheduler(rpc.MetaschedulerAddress, rpc.EthereumBackend)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	return &eventSubscriber{
		rpc:              rpc,
		ws:               ws,
		rpcMetascheduler: m,
	}
}

func (c *eventSubscriber) SubscribeEvents(
	ctx context.Context,
	opts ...types.SubscriptionOption,
) (ethereum.Subscription, error) {
	logs := make(chan ethtypes.Log, 100)
	var o types.SubscriptionOptions
	for _, opt := range opts {
		opt(&o)
	}

	addresses := []common.Address{
		c.ws.MetaschedulerAddress,
	}
	var topics []common.Hash

	// MetaScheduler events
	if o.NewJobRequestChan != nil {
		topics = append(topics, newJobRequestEvent.ID)
	}

	if o.JobTransitionChan != nil {
		topics = append(topics, jobTransitionEvent.ID)
	}

	// CreditManager events
	var creditAddress common.Address
	if o.ApprovalChan != nil {
		if (creditAddress == common.Address{}) {
			creditAddress, err := c.rpcMetascheduler.Credit(&bind.CallOpts{
				Context: ctx,
			})
			if err != nil {
				panic(fmt.Errorf("failed to find Credit address: %w", err))
			}
			addresses = append(addresses, creditAddress)
		}
		topics = append(topics, approvalEvent.ID)
	}
	if o.TransferChan != nil {
		if (creditAddress == common.Address{}) {
			creditAddress, err := c.rpcMetascheduler.Credit(&bind.CallOpts{
				Context: ctx,
			})
			if err != nil {
				panic(fmt.Errorf("failed to find Credit address: %w", err))
			}
			addresses = append(addresses, creditAddress)
		}
		topics = append(topics, transferEvent.ID)
	}

	query := ethereum.FilterQuery{
		Addresses: addresses,
		Topics: [][]common.Hash{
			topics,
		},
	}

	sub, err := c.ws.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		close(logs)
		return nil, WrapError(err)
	}

	// Pass close signal to channel
	go func() {
		<-sub.Err()
		close(logs)
	}()

	go c.filter(logs, o)

	return sub, nil
}

func (c *eventSubscriber) filter(
	logs <-chan ethtypes.Log,
	o types.SubscriptionOptions,
) {
	var creditFilterer metaschedulerabi.IERC20Filterer
	if o.TransferChan != nil || o.ApprovalChan != nil {
		creditAddress, err := c.rpcMetascheduler.Credit(&bind.CallOpts{})
		if err != nil {
			panic(fmt.Errorf("failed to fetch Credit: %w", err))
		}
		ierc20, err := metaschedulerabi.NewIERC20(creditAddress, c.rpc)
		if err != nil {
			panic(fmt.Errorf("failed to instanciate Credit: %w", err))
		}
		creditFilterer = ierc20.IERC20Filterer
	}
	for log := range logs {
		if len(log.Topics) == 0 {
			return
		}
		switch log.Topics[0].Hex() {
		case newJobRequestEvent.ID.Hex():
			event, err := c.rpcMetascheduler.ParseNewJobRequestEvent(log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}

			if o.NewJobRequestChan != nil {
				o.NewJobRequestChan <- event
			}
		case jobTransitionEvent.ID.Hex():
			event, err := c.rpcMetascheduler.ParseJobTransitionEvent(log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}

			if o.JobTransitionChan != nil {
				o.JobTransitionChan <- event
			}
		case transferEvent.ID.Hex():
			event, err := creditFilterer.ParseTransfer(log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}

			if o.TransferChan != nil {
				o.TransferChan <- event
			}
		case approvalEvent.ID.Hex():
			event, err := creditFilterer.ParseApproval(log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}

			if o.ApprovalChan != nil {
				o.ApprovalChan <- event
			}
		}
	}
}
