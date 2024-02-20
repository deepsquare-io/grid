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
	"sort"
	"strconv"
	"strings"

	internallog "github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/deepsquare-io/grid/cli/types/provider"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

var _ provider.Manager = (*ProviderManager)(nil)

// ProviderManager is a manager for providers.
type ProviderManager struct {
	*RPCClientSet
	*metaschedulerabi.IProviderManager
}

// ApproveProvider approves a provider.
func (c *ProviderManager) ApproveProvider(ctx context.Context, provider common.Address) error {
	tx, err := c.transact(ctx, func(auth *bind.TransactOpts) (*coretypes.Transaction, error) {
		return c.IProviderManager.Approve(auth, provider)
	})
	if err != nil {
		return WrapError(err)
	}
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}
	internallog.I.Debug("approve provider", zap.Any("receipt", receipt))
	return CheckReceiptError(ctx, c, tx, receipt)
}

// RemoveProvider removes a provider.
func (c *ProviderManager) RemoveProvider(ctx context.Context, provider common.Address) error {
	tx, err := c.transact(ctx, func(auth *bind.TransactOpts) (*coretypes.Transaction, error) {
		return c.IProviderManager.Remove(auth, provider)
	})
	if err != nil {
		return WrapError(err)
	}
	receipt, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}
	internallog.I.Debug("removed provider", zap.Any("receipt", receipt))
	return CheckReceiptError(ctx, c, tx, receipt)
}

// GetProvider returns a provider.
func (c *ProviderManager) GetProvider(
	ctx context.Context,
	address common.Address,
	opts ...provider.GetProviderOption,
) (detail provider.Detail, err error) {
	var o provider.GetProviderOptions
	for _, opt := range opts {
		opt(&o)
	}
	var p metaschedulerabi.Provider
	if o.Proposal {
		p, err = c.GetWaitingForApprovalProvider(
			&bind.CallOpts{Context: ctx},
			address,
		)
		if err != nil {
			return detail, WrapError(err)
		}
	} else {
		p, err = c.IProviderManager.GetProvider(
			&bind.CallOpts{Context: ctx},
			address,
		)
		if err != nil {
			return detail, WrapError(err)
		}
	}

	isWaitingForApproval, err := c.IsWaitingForApproval(
		&bind.CallOpts{Context: ctx},
		address,
	)
	if err != nil {
		return detail, WrapError(err)
	}

	isValidForScheduling, err := c.IsValidForScheduling(
		&bind.CallOpts{Context: ctx},
		address,
	)
	if err != nil {
		return detail, WrapError(err)
	}

	jobCount, err := c.GetJobCount(
		&bind.CallOpts{Context: ctx},
		address,
	)
	if err != nil {
		return detail, WrapError(err)
	}

	// Sort labels
	sort.Slice(p.Labels, func(i, j int) bool { return p.Labels[i].Key < p.Labels[j].Key })

	p.Addr = address

	return provider.Detail{
		Provider:             p,
		IsWaitingForApproval: isWaitingForApproval,
		IsValidForScheduling: isValidForScheduling,
		JobCount:             jobCount,
	}, nil
}

// GetProviders returns all providers.
func (c *ProviderManager) GetProviders(
	ctx context.Context,
	opts ...provider.GetProviderOption,
) (providers []provider.Detail, err error) {
	var o provider.GetProviderOptions
	for _, opt := range opts {
		opt(&o)
	}

	it, err := c.FilterProviderWaitingForApproval(&bind.FilterOpts{Context: ctx})
	if err != nil {
		return providers, WrapError(err)
	}
	defer func() {
		_ = it.Close()
	}()

	providerMap := make(map[common.Address]provider.Detail)

	for it.Next() {
		var prov metaschedulerabi.Provider
		if o.Proposal {
			prov, err = c.GetWaitingForApprovalProvider(
				&bind.CallOpts{Context: ctx},
				it.Event.Addr,
			)
			if err != nil {
				return providers, WrapError(err)
			}
		} else {
			prov, err = c.IProviderManager.GetProvider(
				&bind.CallOpts{Context: ctx},
				it.Event.Addr,
			)
			if err != nil {
				return providers, WrapError(err)
			}
		}

		// Check if provider matches affinities
		if len(o.Affinities) != 0 && !CheckAffinities(o.Affinities, prov.Labels) {
			continue
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
		sort.Slice(
			prov.Labels,
			func(i, j int) bool { return prov.Labels[i].Key < prov.Labels[j].Key },
		)

		prov.Addr = it.Event.Addr

		providerMap[it.Event.Addr] = provider.Detail{
			Provider:             prov,
			IsWaitingForApproval: isWaitingForApproval,
			IsValidForScheduling: isValidForScheduling,
			JobCount:             jobCount,
		}
	}

	providers = make([]provider.Detail, 0, len(providerMap))
	for _, v := range providerMap {
		providers = append(providers, v)
	}

	return providers, nil
}

// CompareValues compares two values using the given operator.
func CompareValues(op, valueA, valueB string) bool {
	numA, errA := strconv.ParseFloat(valueA, 64)
	numB, errB := strconv.ParseFloat(valueB, 64)
	if errA == nil && errB == nil {
		switch op {
		case "=", "==":
			return numB == numA
		case "in":
			return strings.Contains(valueB, valueA)
		case ">":
			return numB > numA
		case "<":
			return numB < numA
		case ">=":
			return numB >= numA
		case "<=":
			return numB <= numA
		case "!=":
			return numB != numA
		default:
			return numB == numA
		}
	} else {
		// Perform simple string comparison
		switch op {
		case "=", "==":
			return valueB == valueA
		case "in":
			return strings.Contains(valueB, valueA)
		case ">":
			return valueB > valueA
		case "<":
			return valueB < valueA
		case ">=":
			return valueB >= valueA
		case "<=":
			return valueB <= valueA
		case "!=":
			return valueB != valueA
		default:
			return valueB == valueA
		}
	}
}

// CheckAffinities checks if the given labels satisfy the given affinities.
func CheckAffinities(
	affinities []types.Affinity,
	labels []metaschedulerabi.Label,
) bool {
	kv := make(map[string]string)
	for _, affinity := range affinities {
		op := strings.Trim(string(affinity.Op[:]), "\x00")

		for _, label := range labels {
			// If key found and condition satisfied
			if affinity.Label.Key == label.Key &&
				CompareValues(op, affinity.Label.Value, label.Value) {
				kv[affinity.Label.Key] = affinity.Label.Value
				break
			}
		}
	}

	for _, item := range affinities {
		if _, found := kv[item.Label.Key]; !found {
			return false
		}
	}

	return true
}
