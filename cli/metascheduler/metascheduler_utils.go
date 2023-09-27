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
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
)

// DurationToCredit converts a job duration to credits based on pricing and resources allocation.
func DurationToCredit(
	prices metaschedulerabi.ProviderPrices,
	definition metaschedulerabi.JobDefinition,
	durationMinutes *big.Int,
) *big.Int {
	// GpuPricePerMin * GpusPerTask
	gpusPricePerMinPerTask := new(big.Int).
		Mul(prices.GpuPricePerMin, new(big.Int).SetUint64(definition.GpusPerTask))

	// CpuPricePerMin * CpusPerTask
	cpusPricePerMinPerTask := new(big.Int).
		Mul(prices.CpuPricePerMin, new(big.Int).SetUint64(definition.CpusPerTask))

	// MemPerCpu * CpusPerTask * MemPricePerMin
	memPricePerMinPerTask := new(
		big.Int,
	).Mul(prices.MemPricePerMin, new(big.Int).SetUint64(definition.MemPerCpu))
	memPricePerMinPerTask = memPricePerMinPerTask.Mul(
		memPricePerMinPerTask,
		new(big.Int).SetUint64(definition.CpusPerTask),
	)

	// pricePerMin = Ntasks * (cpusPricePerMinPerTask + memPricePerMinPerTask + gpusPricePerMinPerTask)
	b := new(big.Int).Add(gpusPricePerMinPerTask, cpusPricePerMinPerTask)
	b = b.Add(b, memPricePerMinPerTask)
	b = b.Mul(b, new(big.Int).SetUint64(definition.Ntasks))

	// price = duration*pricePerMin
	b.Mul(b, durationMinutes)
	return b
}

// CreditToDuration converts credits to a job duration based on pricing and resources allocation.
func CreditToDuration(
	prices metaschedulerabi.ProviderPrices,
	definition metaschedulerabi.JobDefinition,
	creditsWei *big.Int,
) *big.Int {
	// GpuPricePerMin * GpusPerTask
	gpusPricePerMinPerTask := new(big.Int).
		Mul(prices.GpuPricePerMin, new(big.Int).SetUint64(definition.GpusPerTask))

	// CpuPricePerMin * CpusPerTask
	cpusPricePerMinPerTask := new(big.Int).
		Mul(prices.CpuPricePerMin, new(big.Int).SetUint64(definition.CpusPerTask))

	// MemPerCpu * CpusPerTask * MemPricePerMin
	memPricePerMinPerTask := new(
		big.Int,
	).Mul(prices.MemPricePerMin, new(big.Int).SetUint64(definition.MemPerCpu))
	memPricePerMinPerTask = memPricePerMinPerTask.Mul(
		memPricePerMinPerTask,
		new(big.Int).SetUint64(definition.CpusPerTask),
	)

	// pricePerMin = Ntasks * (cpusPricePerMinPerTask + memPricePerMinPerTask + gpusPricePerMinPerTask)
	pricePerMin := new(big.Int).Add(gpusPricePerMinPerTask, cpusPricePerMinPerTask)
	pricePerMin = pricePerMin.Add(pricePerMin, memPricePerMinPerTask)
	pricePerMin = pricePerMin.Mul(pricePerMin, new(big.Int).SetUint64(definition.Ntasks))

	// duration = price/pricePerMin
	return new(big.Int).Div(creditsWei, pricePerMin)
}
