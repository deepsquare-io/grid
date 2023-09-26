// Copyright (C) 2023 DeepSquare Asociation
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
	"context"
	"math/big"

	loggerv1alpha1 "github.com/deepsquare-io/grid/cli/internal/logger/v1alpha1"
	"github.com/deepsquare-io/grid/cli/sbatch"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// LogStream is a readable stream of logs.
type LogStream loggerv1alpha1.LoggerAPI_ReadClient

// Logger fetches the logs of a job.
type Logger interface {
	// Watch the logs of a job
	WatchLogs(ctx context.Context, jobID [32]byte) (LogStream, error)
}

type Job *metaschedulerabi.Job
type Label metaschedulerabi.Label
type Affinity metaschedulerabi.Affinity

type SubmitJobOption func(*SubmitJobOptions)
type SubmitJobOptions struct {
	Uses       []Label
	Affinities []Affinity
}

func WithUse(labels ...Label) SubmitJobOption {
	return func(sjo *SubmitJobOptions) {
		sjo.Uses = labels
	}
}

func WithAffinity(affinities ...Affinity) SubmitJobOption {
	return func(sjo *SubmitJobOptions) {
		sjo.Affinities = affinities
	}
}

// JobScheduler schedules and cancels jobs.
type JobScheduler interface {
	// Submit a batch script to the batch service and metascheduler.
	SubmitJob(
		ctx context.Context,
		job *sbatch.Job,
		lockedAmount *big.Int,
		jobName [32]byte,
		opts ...SubmitJobOption,
	) ([32]byte, error)
	// Cancel a job.
	CancelJob(ctx context.Context, jobID [32]byte) error
	// TopUp a job.
	TopUpJob(ctx context.Context, jobID [32]byte, amount *big.Int) error
	// Panic a job.
	PanicJob(ctx context.Context, jobID [32]byte, reason string) error
}

// JobLazyIterator iterates on a lazy list of jobs.
//
// When calling Next or Prev, a request will be sent to the data source.
type JobLazyIterator interface {
	// Fetches the next job.
	Next(ctx context.Context) (next JobLazyIterator, ok bool, err error)
	// Fetches the previous job.
	Prev(ctx context.Context) (prev JobLazyIterator, ok bool, err error)
	// Get the current job.
	Current() Job
}

// JobFetcher fetches jobs.
type JobFetcher interface {
	// Get a job.
	GetJob(ctx context.Context, id [32]byte) (Job, error)
	// Get a iterator of jobs. If there is no job, nil is returned.
	GetJobs(ctx context.Context) (JobLazyIterator, error)
}

// CreditManager handles the credits of the user.
type CreditManager interface {
	// Balance fetches the current balance of credits.
	Balance(ctx context.Context) (*big.Int, error)

	// Balance fetches the current balance of credits.
	BalanceOf(ctx context.Context, address common.Address) (*big.Int, error)

	// Transfer tranfers credits from one address to another.
	Transfer(ctx context.Context, to common.Address, amount *big.Int) error

	// ReduceToBalance reduces a channel of transfers into balance.
	ReduceToBalance(
		ctx context.Context,
		transfers <-chan Transfer,
	) (<-chan *big.Int, error)
}

// AllowanceManager set the allowed quantity of credit for smart-contract interactions.
type AllowanceManager interface {
	// Set the allowance for smart-contract interactions.
	SetAllowance(ctx context.Context, amount *big.Int) error

	// ClearAllowance is an alias to SetAllowance 0.
	ClearAllowance(ctx context.Context) error

	// Get the current allowance toward the contract.
	GetAllowance(ctx context.Context) (*big.Int, error)

	// ReduceToAllowance reduces a channel of approval into allowance.
	ReduceToAllowance(
		ctx context.Context,
		approvals <-chan Approval,
	) (<-chan *big.Int, error)
}

type ProviderDetail struct {
	metaschedulerabi.Provider
	IsWaitingForApproval bool
	IsValidForScheduling bool
	JobCount             uint64
}

// ProviderManager manages admin operation of providers
type ProviderManager interface {
	Approve(ctx context.Context, provider common.Address) error
	Remove(ctx context.Context, provider common.Address) error
	GetProvider(ctx context.Context, address common.Address) (provider ProviderDetail, err error)
	GetProviders(ctx context.Context) (providers []ProviderDetail, err error)
}

type NewJobRequest *metaschedulerabi.MetaSchedulerNewJobRequestEvent
type JobTransition *metaschedulerabi.MetaSchedulerJobTransitionEvent
type Transfer *metaschedulerabi.IERC20Transfer
type Approval *metaschedulerabi.IERC20Approval

type SubscriptionOptions struct {
	NewJobRequestChan chan<- NewJobRequest
	JobTransitionChan chan<- JobTransition
	TransferChan      chan<- Transfer
	ApprovalChan      chan<- Approval
}

type SubscriptionOption func(*SubscriptionOptions)

func FilterTransfer(filtered chan<- Transfer) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.TransferChan = filtered
	}
}

func FilterApproval(filtered chan<- Approval) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.ApprovalChan = filtered
	}
}

func FilterNewJobRequest(
	filtered chan<- NewJobRequest,
) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.NewJobRequestChan = filtered
	}
}

func FilterJobTransition(
	filtered chan<- JobTransition,
) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.JobTransitionChan = filtered
	}
}

// EventSubscriber watches smart-contract events.
type EventSubscriber interface {
	// Subscribe to metascheduler events.
	SubscribeEvents(
		ctx context.Context,
		opts ...SubscriptionOption,
	) (ethereum.Subscription, error)
}
