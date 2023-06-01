package metricsv1

import (
	"math/big"

	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/contracts/metascheduler"
	"github.com/ethereum/go-ethereum/common"
)

func MetaToProtoJob(msJob *struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       metascheduler.JobDefinition
	Valid            bool
	Cost             metascheduler.JobCost
	Time             metascheduler.JobTime
	JobName          [32]byte
	HasCancelRequest bool
}) *Job {
	pJob := &Job{
		JobId:            msJob.JobId[:],
		Status:           uint32(msJob.Status),
		CustomerAddr:     msJob.CustomerAddr.Hex(),
		ProviderAddr:     msJob.ProviderAddr.Hex(),
		Valid:            msJob.Valid,
		HasCancelRequest: msJob.HasCancelRequest,
		Definition: &JobDefinition{
			GpuPerTask: uint64(msJob.Definition.GpuPerTask),
			MemPerCpu:  uint64(msJob.Definition.MemPerCpu),
			CpuPerTask: uint64(msJob.Definition.CpuPerTask),
			Ntasks:     uint64(msJob.Definition.Ntasks),
		},
		Cost: &JobCost{
			MaxCost:                   &BigInt{Bytes: msJob.Cost.MaxCost.Bytes()},
			FinalCost:                 &BigInt{Bytes: msJob.Cost.FinalCost.Bytes()},
			PendingTopUp:              &BigInt{Bytes: msJob.Cost.PendingTopUp.Bytes()},
			DelegateSpendingAuthority: msJob.Cost.DelegateSpendingAuthority,
		},
		Time: &JobTime{
			Start:                  &BigInt{Bytes: msJob.Time.Start.Bytes()},
			End:                    &BigInt{Bytes: msJob.Time.End.Bytes()},
			CancelRequestTimestamp: &BigInt{Bytes: msJob.Time.CancelRequestTimestamp.Bytes()},
			BlockNumberStateChange: &BigInt{Bytes: msJob.Time.BlockNumberStateChange.Bytes()},
		},
		JobName: msJob.JobName[:],
	}

	return pJob
}

func ProtoToMetaJob(protoJob *Job) *struct {
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       metascheduler.JobDefinition
	Valid            bool
	Cost             metascheduler.JobCost
	Time             metascheduler.JobTime
	JobName          [32]byte
	HasCancelRequest bool
} {
	metaJob := &struct {
		JobId            [32]byte
		Status           uint8
		CustomerAddr     common.Address
		ProviderAddr     common.Address
		Definition       metascheduler.JobDefinition
		Valid            bool
		Cost             metascheduler.JobCost
		Time             metascheduler.JobTime
		JobName          [32]byte
		HasCancelRequest bool
	}{
		Status:       uint8(protoJob.Status),
		CustomerAddr: common.HexToAddress(protoJob.CustomerAddr),
		ProviderAddr: common.HexToAddress(protoJob.ProviderAddr),
		Definition: metascheduler.JobDefinition{
			GpuPerTask:        protoJob.Definition.GpuPerTask,
			MemPerCpu:         protoJob.Definition.MemPerCpu,
			CpuPerTask:        protoJob.Definition.CpuPerTask,
			Ntasks:            protoJob.Definition.Ntasks,
			BatchLocationHash: protoJob.Definition.BatchLocationHash,
			StorageType:       uint8(protoJob.Definition.StorageType),
		},
		Valid: protoJob.Valid,
		Cost: metascheduler.JobCost{
			MaxCost:                   new(big.Int).SetBytes(protoJob.Cost.MaxCost.Bytes),
			FinalCost:                 new(big.Int).SetBytes(protoJob.Cost.FinalCost.Bytes),
			PendingTopUp:              new(big.Int).SetBytes(protoJob.Cost.PendingTopUp.Bytes),
			DelegateSpendingAuthority: protoJob.Cost.DelegateSpendingAuthority,
		},
		Time: metascheduler.JobTime{
			Start:                  new(big.Int).SetBytes(protoJob.Time.Start.Bytes),
			End:                    new(big.Int).SetBytes(protoJob.Time.End.Bytes),
			CancelRequestTimestamp: new(big.Int).SetBytes(protoJob.Time.CancelRequestTimestamp.Bytes),
			BlockNumberStateChange: new(big.Int).SetBytes(protoJob.Time.BlockNumberStateChange.Bytes),
		},
		HasCancelRequest: protoJob.GetHasCancelRequest(),
	}

	copy(metaJob.JobId[:], protoJob.JobId)
	copy(metaJob.JobName[:], protoJob.JobName)
	return metaJob
}
