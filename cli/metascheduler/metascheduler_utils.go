// Copyright (C) 2023 DeepSquare Association
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
	"errors"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
)

var zero = new(big.Int)
var DivByZero = errors.New("division by zero")

func creditsPerMin(
	prices metaschedulerabi.ProviderPrices,
	definition metaschedulerabi.JobDefinition,
) *big.Int {
	// GpuPricePerMin * GpusPerTask
	gpusCostPerMinPerTask := new(big.Int).
		Mul(prices.GpuPricePerMin, new(big.Int).SetUint64(definition.GpusPerTask))

	// CpuPricePerMin * CpusPerTask
	cpusCostPerMinPerTask := new(big.Int).
		Mul(prices.CpuPricePerMin, new(big.Int).SetUint64(definition.CpusPerTask))

	// MemPerCpu * CpusPerTask * MemPricePerMin
	memCostPerMinPerTask := new(big.Int).
		Mul(prices.MemPricePerMin, new(big.Int).SetUint64(definition.MemPerCpu))
	memCostPerMinPerTask.Mul(memCostPerMinPerTask, new(big.Int).SetUint64(definition.CpusPerTask))

	// creditsPerMin = Ntasks * (cpusPricePerMinPerTask + memPricePerMinPerTask + gpusPricePerMinPerTask)
	creditsPerMin := new(big.Int).Set(gpusCostPerMinPerTask)
	creditsPerMin.Add(creditsPerMin, cpusCostPerMinPerTask)
	creditsPerMin.Add(creditsPerMin, memCostPerMinPerTask)
	creditsPerMin.Mul(creditsPerMin, new(big.Int).SetUint64(definition.Ntasks))

	return creditsPerMin
}

// DurationToCredit converts a job duration to credits based on pricing and resources allocation.
func DurationToCredit(
	prices metaschedulerabi.ProviderPrices,
	definition metaschedulerabi.JobDefinition,
	durationMinutes *big.Int,
) *big.Int {
	creditsPerMin := creditsPerMin(prices, definition)
	// price = duration*pricePerMin
	return creditsPerMin.Mul(creditsPerMin, durationMinutes)
}

// CreditToDuration converts credits to a job duration based on pricing and resources allocation.
func CreditToDuration(
	prices metaschedulerabi.ProviderPrices,
	definition metaschedulerabi.JobDefinition,
	creditsWei *big.Int,
) (*big.Int, error) {
	creditsPerMin := creditsPerMin(prices, definition)
	// duration = price/pricePerMin
	if creditsPerMin.Cmp(zero) == 0 {
		return nil, DivByZero
	}
	return new(big.Int).Div(creditsWei, creditsPerMin), nil
}
