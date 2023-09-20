package types

import (
	"context"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/internal/abi/metascheduler"
	loggerv1alpha1 "github.com/deepsquare-io/the-grid/cli/internal/logger/v1alpha1"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
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

// JobScheduler schedules and cancels jobs.
type JobScheduler interface {
	// Submit a batch script to the batch service and metascheduler.
	SubmitJob(
		ctx context.Context,
		job *sbatch.Job,
		uses []metaschedulerabi.Label,
		lockedAmount *big.Int,
		jobName [32]byte,
	) ([32]byte, error)
	// Cancel a job.
	CancelJob(ctx context.Context, jobID [32]byte) error
	// TopUp a job.
	TopUpJob(ctx context.Context, jobID [32]byte, amount *big.Int) error
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
	Current() *metaschedulerabi.Job
}

// JobFetcher fetches jobs.
type JobFetcher interface {
	// Get a job.
	GetJob(ctx context.Context, id [32]byte) (*metaschedulerabi.Job, error)
	// Get a iterator of jobs. If there is no job, nil is returned.
	GetJobs(ctx context.Context) (JobLazyIterator, error)
}

// // JobFilterer watches smart-contract events.
// type JobFilterer interface {
// 	// Filter new job requests events.
// 	FilterNewJobRequests(
// 		ch <-chan types.Log,
// 	) (filtered <-chan *metaschedulerabi.MetaSchedulerNewJobRequestEvent, rest <-chan types.Log)
// 	// Filter job transition events.
// 	FilterJobTransition(
// 		ch <-chan types.Log,
// 	) (filtered <-chan *metaschedulerabi.MetaSchedulerJobTransitionEvent, rest <-chan types.Log)
// }

// CreditManager handles the credits of the user.
type CreditManager interface {
	// Balance fetches the current balance of credits.
	Balance(ctx context.Context) (*big.Int, error)

	// Transfer tranfers credits from one address to another.
	Transfer(ctx context.Context, to common.Address, amount *big.Int) error

	// ReduceToBalance reduces a channel of transfers into balance.
	ReduceToBalance(
		ctx context.Context,
		transfers <-chan *metaschedulerabi.IERC20Transfer,
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
		approvals <-chan *metaschedulerabi.IERC20Approval,
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
	GetProviders(ctx context.Context) (providers []ProviderDetail, err error)
}

type SubscriptionOptions struct {
	NewJobRequestChan chan<- *metaschedulerabi.MetaSchedulerNewJobRequestEvent
	JobTransitionChan chan<- *metaschedulerabi.MetaSchedulerJobTransitionEvent
	TransferChan      chan<- *metaschedulerabi.IERC20Transfer
	ApprovalChan      chan<- *metaschedulerabi.IERC20Approval
}

type SubscriptionOption func(*SubscriptionOptions)

func FilterTransfer(filtered chan<- *metaschedulerabi.IERC20Transfer) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.TransferChan = filtered
	}
}

func FilterApproval(filtered chan<- *metaschedulerabi.IERC20Approval) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.ApprovalChan = filtered
	}
}

func FilterNewJobRequest(
	filtered chan<- *metaschedulerabi.MetaSchedulerNewJobRequestEvent,
) SubscriptionOption {
	return func(so *SubscriptionOptions) {
		so.NewJobRequestChan = filtered
	}
}

func FilterJobTransition(
	filtered chan<- *metaschedulerabi.MetaSchedulerJobTransitionEvent,
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
