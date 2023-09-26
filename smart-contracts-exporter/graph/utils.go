package graph

import (
	"github.com/deepsquare-io/grid/smart-contracts-exporter/contracts/metascheduler"
	"github.com/deepsquare-io/grid/smart-contracts-exporter/graph/model"
	"github.com/deepsquare-io/grid/smart-contracts-exporter/graph/scalar"
	"github.com/ethereum/go-ethereum/common"
)

func mapMetaschedulerJobToGraphJob(j *metascheduler.Job) model.Job {
	jobDefinition := &model.JobDefinition{
		GpuPerTask: int(j.Definition.GpusPerTask),
		MemPerCPU:  int(j.Definition.MemPerCpu),
		CPUPerTask: int(j.Definition.CpusPerTask),
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
		Cost:             jobCost,
		Time:             jobTime,
		JobName:          string(j.JobName[:]),
		HasCancelRequest: j.HasCancelRequest,
	}
}

func ConvertJobs(input <-chan *metascheduler.Job) <-chan *model.Job {
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
