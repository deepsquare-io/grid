// Copyright (C) 2023 DeepSquare Association
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

package metascheduler

import (
	"context"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"go.uber.org/zap"
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

func FromStructToJob(s metaschedulerabi.Job) *Job {
	return &Job{
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

type ProviderJobIterator struct {
	Job *Job
	*metaschedulerabi.MetaSchedulerClaimJobEventIterator
	client          MetaScheduler
	providerAddress common.Address
}

func (it *ProviderJobIterator) Next(ctx context.Context) bool {
	for it.MetaSchedulerClaimJobEventIterator.Next() {
		if it.Event.ProviderAddr != it.providerAddress {
			continue
		}

		job, err := it.client.GetJob(ctx, it.Event.JobId)
		if err != nil {
			logger.I.Error("GetJob failed", zap.Error(err))
			return false
		}
		if job.ProviderAddr != it.providerAddress {
			continue
		}

		it.Job = job
		return true
	}

	return false
}

type setJobStatusOptions struct {
	err      error
	exitCode int64
}

func applySetJobStatusOptions(opts []SetJobStatusOption) *setJobStatusOptions {
	o := &setJobStatusOptions{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type SetJobStatusOption func(*setJobStatusOptions)

func SetJobStatusWithError(err error) SetJobStatusOption {
	return func(sjso *setJobStatusOptions) {
		sjso.err = err
	}
}

func SetJobStatusWithExitCode(exitCode int64) SetJobStatusOption {
	return func(sjso *setJobStatusOptions) {
		sjso.exitCode = exitCode
	}
}

type MetaScheduler interface {
	IsRequestNewJobEnabled(ctx context.Context) (bool, error)
	// Claim a job for scheduling.
	Claim(ctx context.Context) error
	// Claim cancelling calls.
	ClaimCancelling(ctx context.Context) error
	// Claim top up calls.
	ClaimTopUp(ctx context.Context) error
	// Refuse a job for metascheduling.
	RefuseJob(ctx context.Context, jobID [32]byte) error
	// WatchEvents observes the incoming ClaimNextTopUpJobEvent, ClaimNextCancellingJobEvent and ClaimJobEvent.
	WatchEvents(
		ctx context.Context,
		claimNextTopUpJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent,
		claimNextCancellingJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent,
		claimJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent,
	) (event.Subscription, error)
	// GetProviderAddress fetches the provider public address.
	GetProviderAddress() common.Address
	// GetOldInfo fetches the last registration provider information.
	GetOldInfo(ctx context.Context) (*metaschedulerabi.Provider, error)
	// GetJobStatus fetches the job status.
	GetJobStatus(ctx context.Context, jobID [32]byte) (JobStatus, error)
	SetJobStatus(
		ctx context.Context,
		jobID [32]byte,
		status JobStatus,
		jobDurationMinute uint64,
		opts ...SetJobStatusOption,
	) error
	Register(
		ctx context.Context,
		hardware metaschedulerabi.ProviderHardware,
		prices metaschedulerabi.ProviderPrices,
		labels []metaschedulerabi.Label,
	) error
	// GetJob fetches a job.
	GetJob(ctx context.Context, jobID [32]byte) (*Job, error)
	// GetJobs fetches the jobs of the provider.
	GetJobs(ctx context.Context) (*ProviderJobIterator, error)
}
