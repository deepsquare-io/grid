package graph

import (
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/contracts/metascheduler"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/graph/model"
	"github.com/deepsquare-io/the-grid/smart-contracts-exporter/graph/scalar"
	"github.com/ethereum/go-ethereum/common"
)

func mapMetaschedulerJobToGraphJob(j *struct {
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
}) model.Job {
	jobDefinition := &model.JobDefinition{
		GpuPerTask: int(j.Definition.GpuPerTask),
		MemPerCPU:  int(j.Definition.MemPerCpu),
		CPUPerTask: int(j.Definition.CpuPerTask),
		Ntasks:     int(j.Definition.Ntasks),
	}

	jobCost := &model.JobCost{
		MaxCost:                   scalar.BigInt{Int: j.Cost.MaxCost},
		FinalCost:                 scalar.BigInt{Int: j.Cost.FinalCost},
		PendingTopUp:              scalar.BigInt{Int: j.Cost.PendingTopUp},
		DelegateSpendingAuthority: j.Cost.DelegateSpendingAuthority,
	}

	jobTime := &model.JobTime{
		Start:                  scalar.BigInt{Int: j.Time.Start},
		End:                    scalar.BigInt{Int: j.Time.End},
		CancelRequestTimestamp: scalar.BigInt{Int: j.Time.CancelRequestTimestamp},
		BlockNumberStateChange: scalar.BigInt{Int: j.Time.BlockNumberStateChange},
	}

	return model.Job{
		JobID:            common.Bytes2Hex(j.JobId[:]),
		Status:           int(j.Status),
		CustomerAddr:     j.CustomerAddr.Hex(),
		ProviderAddr:     j.ProviderAddr.Hex(),
		Definition:       jobDefinition,
		Valid:            j.Valid,
		Cost:             jobCost,
		Time:             jobTime,
		JobName:          string(j.JobName[:]),
		HasCancelRequest: j.HasCancelRequest,
	}
}

func ConvertJobs(input <-chan *struct {
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
}) <-chan *model.Job {
	output := make(chan *model.Job)

	go func() {
		defer close(output)
		for j := range input {
			convertedJob := mapMetaschedulerJobToGraphJob(j)
			output <- &convertedJob
		}
	}()

	return output
}