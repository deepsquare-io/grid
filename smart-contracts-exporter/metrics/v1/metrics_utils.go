package metricsv1

import (
	"math/big"

	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/contracts/metascheduler"
	"github.com/ethereum/go-ethereum/common"
)

func MetaToProtoJob(msJob *metascheduler.Job) *Job {
	pJob := &Job{
		JobId:        msJob.JobId[:],
		Status:       uint32(msJob.Status),
		CustomerAddr: msJob.CustomerAddr.Hex(),
		ProviderAddr: msJob.ProviderAddr.Hex(),
		Definition: &JobDefinition{
			GpusPerTask: uint64(msJob.Definition.GpusPerTask),
			MemPerCpu:   uint64(msJob.Definition.MemPerCpu),
			CpusPerTask: uint64(msJob.Definition.CpusPerTask),
			Ntasks:      uint64(msJob.Definition.Ntasks),
		},
		Cost: &JobCost{
			MaxCost:   &BigInt{Bytes: msJob.Cost.MaxCost.Bytes()},
			FinalCost: &BigInt{Bytes: msJob.Cost.FinalCost.Bytes()},
		},
		Time: &JobTime{
			Start: &BigInt{Bytes: msJob.Time.Start.Bytes()},
			End:   &BigInt{Bytes: msJob.Time.End.Bytes()},
		},
		JobName: msJob.JobName[:],
	}

	return pJob
}

func ProtoToMetaJob(protoJob *Job) *metascheduler.Job {
	metaJob := &metascheduler.Job{
		Status:       uint8(protoJob.Status),
		CustomerAddr: common.HexToAddress(protoJob.CustomerAddr),
		ProviderAddr: common.HexToAddress(protoJob.ProviderAddr),
		Definition: metascheduler.JobDefinition{
			GpusPerTask: protoJob.Definition.GpusPerTask,
			MemPerCpu:   protoJob.Definition.MemPerCpu,
			CpusPerTask: protoJob.Definition.CpusPerTask,
			Ntasks:      protoJob.Definition.Ntasks,
		},
		Cost: metascheduler.JobCost{
			MaxCost:      new(big.Int).SetBytes(protoJob.Cost.MaxCost.Bytes),
			FinalCost:    new(big.Int).SetBytes(protoJob.Cost.FinalCost.Bytes),
			PendingTopUp: new(big.Int),
		},
		Time: metascheduler.JobTime{
			Start: new(big.Int).SetBytes(protoJob.Time.Start.Bytes),
			End:   new(big.Int).SetBytes(protoJob.Time.End.Bytes),
			CancelRequestTimestamp: new(
				big.Int,
			),
			BlockNumberStateChange: new(
				big.Int,
			),
		},
	}

	copy(metaJob.JobId[:], protoJob.JobId)
	copy(metaJob.JobName[:], protoJob.JobName)
	return metaJob
}
