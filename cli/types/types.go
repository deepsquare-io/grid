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

// Package types provides the main types of the library.
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

// Job is the object stored in the smart-contract for accounting.
type Job *metaschedulerabi.Job

// Label is a key-value object used for filtering and annotating clusters.
type Label metaschedulerabi.Label

// Affinity is a key-value object with an operator for filtering clusters.
type Affinity metaschedulerabi.Affinity

// SubmitJobOption is used to apply default and optional parameters for submitting a job.
type SubmitJobOption func(*SubmitJobOptions)

// SubmitJobOptions is the object containing optional parameters for submitting a job.
type SubmitJobOptions struct {
	Uses       []Label
	Affinities []Affinity
}

// WithUse adds strict key-value filters to the job, which filters the available clusters.
func WithUse(labels ...Label) SubmitJobOption {
	return func(sjo *SubmitJobOptions) {
		sjo.Uses = labels
	}
}

// WithAffinity adds key-value filters with operators to the job, which filters the available clusters.
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
	Next(ctx context.Context) (ok bool)
	// Fetches the previous job.
	Prev(ctx context.Context) (ok bool)
	// Get the current job.
	Current() Job
	// Get the current error.
	Error() error
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

// ProviderDetail contains all the specs and statuses of a Provider.
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

// NewJobRequest is an event that happens when a user submit a job.
type NewJobRequest *metaschedulerabi.MetaSchedulerNewJobRequestEvent

// JobTransition is an event that happens when the status of a job changes.
type JobTransition *metaschedulerabi.MetaSchedulerJobTransitionEvent

// Transfer is an event that happens when there is a ERC20 transaction.
type Transfer *metaschedulerabi.IERC20Transfer

// Approval is an event that happens when an user sets a new allowance.
type Approval *metaschedulerabi.IERC20Approval

// SubscriptionOptions contains the channels used to pass events.
type SubscriptionOptions struct {
	NewJobRequestChan chan<- NewJobRequest
	JobTransitionChan chan<- JobTransition
	TransferChan      chan<- Transfer
	ApprovalChan      chan<- Approval
}

// SubscriptionOption applies default and optional parameters to the SubscribeEvents method.
type SubscriptionOption func(*SubscriptionOptions)

// FilterTransfer allows taking the Transfer events from the subscription.
func FilterTransfer(filtered chan<- Transfer) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.TransferChan = filtered
	}
}

// FilterApproval allows taking the Approval events from the subscription.
func FilterApproval(filtered chan<- Approval) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.ApprovalChan = filtered
	}
}

// FilterNewJobRequest allows taking the NewJobRequest events from the subscription.
func FilterNewJobRequest(
	filtered chan<- NewJobRequest,
) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.NewJobRequestChan = filtered
	}
}

// FilterJobTransition allows taking the JobTransition events from the subscription.
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
