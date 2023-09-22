package metascheduler

import (
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
)

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
