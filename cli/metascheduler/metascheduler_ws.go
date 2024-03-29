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

	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/deepsquare-io/grid/cli/types/event"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ event.Subscriber = (*EventSubscriber)(nil)

// EventSubscriber is a subscriber for events.
type EventSubscriber struct {
	rpc Backend
	ws  Backend

	rpcMetascheduler *metaschedulerabi.MetaScheduler
}

// NewEventSubscriber creates an EventSubscriber used to watch events.
func NewEventSubscriber(
	rpc Backend,
	ws Backend,
) *EventSubscriber {
	m, err := metaschedulerabi.NewMetaScheduler(rpc.MetaschedulerAddress, rpc.EthereumBackend)
	if err != nil {
		panic(fmt.Errorf("failed to instanciate MetaScheduler: %w", err))
	}
	return &EventSubscriber{
		rpc:              rpc,
		ws:               ws,
		rpcMetascheduler: m,
	}
}

// SubscribeEvents subscribes to events.
//
//nolint:ireturn
func (c *EventSubscriber) SubscribeEvents(
	ctx context.Context,
	opts ...event.SubscriptionOption,
) (ethereum.Subscription, error) {
	var err error
	logs := make(chan ethtypes.Log, 100)
	var o event.SubscriptionOptions
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

	// JobRepository events
	var jobsAddress common.Address
	if o.JobTransitionChan != nil {
		if (jobsAddress == common.Address{}) {
			jobsAddress, err = c.rpcMetascheduler.Jobs(&bind.CallOpts{
				Context: ctx,
			})
			if err != nil {
				panic(fmt.Errorf("failed to find Credit address: %w", err))
			}
			addresses = append(addresses, jobsAddress)
		}
		topics = append(topics, jobTransitionEvent.ID)
	}

	// CreditManager events
	var creditAddress common.Address
	if o.ApprovalChan != nil {
		if (creditAddress == common.Address{}) {
			creditAddress, err = c.rpcMetascheduler.Credit(&bind.CallOpts{
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
			creditAddress, err = c.rpcMetascheduler.Credit(&bind.CallOpts{
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

func (c *EventSubscriber) filter(
	logs <-chan ethtypes.Log,
	o event.SubscriptionOptions,
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
	var jobsFilterer metaschedulerabi.IJobRepositoryFilterer
	if o.JobTransitionChan != nil {
		jobsAddress, err := c.rpcMetascheduler.Jobs(&bind.CallOpts{})
		if err != nil {
			panic(fmt.Errorf("failed to fetch Jobs: %w", err))
		}
		jobs, err := metaschedulerabi.NewIJobRepository(jobsAddress, c.rpc)
		if err != nil {
			panic(fmt.Errorf("failed to instanciate Credit: %w", err))
		}
		jobsFilterer = jobs.IJobRepositoryFilterer
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
				select {
				case o.NewJobRequestChan <- event:
				default:
				}
			}
		case jobTransitionEvent.ID.Hex():
			event, err := jobsFilterer.ParseJobTransitionEvent(log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}

			if o.JobTransitionChan != nil {
				select {
				case o.JobTransitionChan <- event:
				default:
				}
			}
		case transferEvent.ID.Hex():
			event, err := creditFilterer.ParseTransfer(log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}

			if o.TransferChan != nil {
				select {
				case o.TransferChan <- event:
				default:
				}
			}
		case approvalEvent.ID.Hex():
			event, err := creditFilterer.ParseApproval(log)
			if err != nil {
				panic(fmt.Errorf("failed to parse event: %w", err))
			}

			if o.ApprovalChan != nil {
				select {
				case o.ApprovalChan <- event:
				default:
				}
			}
		}
	}
}
