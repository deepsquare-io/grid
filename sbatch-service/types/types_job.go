// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
