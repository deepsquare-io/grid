package types

import (
	metaschedulerabi "github.com/deepsquare-io/grid/sbatch-service/abi/metascheduler"
	"github.com/ethereum/go-ethereum/common"
)

type Job struct {
	JobID            [32]byte
	Status           JobStatus
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       metaschedulerabi.JobDefinition
	Cost             metaschedulerabi.JobCost
	Time             metaschedulerabi.JobTime
	JobName          [32]byte
	HasCancelRequest bool
	LastError        string
}

func FromStructToJob(s metaschedulerabi.Job) Job {
	return Job{
		JobID:            s.JobId,
		Status:           JobStatus(s.Status),
		CustomerAddr:     s.CustomerAddr,
		ProviderAddr:     s.ProviderAddr,
		Definition:       s.Definition,
		Cost:             s.Cost,
		Time:             s.Time,
		JobName:          s.JobName,
		HasCancelRequest: s.HasCancelRequest,
		LastError:        s.LastError,
	}
}
