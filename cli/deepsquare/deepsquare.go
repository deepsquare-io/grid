package deepsquare

import (
	"context"
	"math/big"

	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/deepsquare/generated/abi/metascheduler"
	loggerv1alpha1 "github.com/deepsquare-io/the-grid/cli/deepsquare/generated/logger/v1alpha1"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// LogStream is a readable stream of logs.
type LogStream loggerv1alpha1.LoggerAPI_ReadClient

// Job represents a job object in the smart-contract.
type Job struct {
	//lint:ignore ST1003 no need to write ID instead of as we map to the metascheduler structure
	JobId            [32]byte
	Status           uint8
	CustomerAddr     common.Address
	ProviderAddr     common.Address
	Definition       metaschedulerabi.JobDefinition
	Valid            bool
	Cost             metaschedulerabi.JobCost
	Time             metaschedulerabi.JobTime
	JobName          [32]byte
	HasCancelRequest bool
}

// Logger fetches the logs of a job.
type Logger interface {
	// Watch the logs of a job
	WatchLogs(ctx context.Context, jobName string) (LogStream, error)
}

// JobScheduler schedules and cancels jobs.
type JobScheduler interface {
	// Submit a batch script to the batch service and metascheduler.
	SubmitJob(
		ctx context.Context,
		content string,
		definition metaschedulerabi.JobDefinition,
		lockedAmount *big.Int,
		jobName [32]byte,
	) ([32]byte, error)
	// Cancel a job.
	CancelJob(ctx context.Context, id [32]byte) error
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
	Current() *Job
}

// JobFetcher fetches jobs.
type JobFetcher interface {
	// Get a job.
	GetJob(ctx context.Context, id [32]byte) (*Job, error)
	// Get a iterator of jobs. If there is no job, nil is returned.
	GetJobs(ctx context.Context) (JobLazyIterator, error)
}

// JobWatcher watches smart-contract events.
type JobWatcher interface {
	// Subscribe to metascheduler events.
	SubscribeEvents(ctx context.Context, ch chan<- types.Log) (ethereum.Subscription, error)
	// Filter new job requests events.
	FilterNewJobRequests(
		ch <-chan types.Log,
	) <-chan *metaschedulerabi.MetaSchedulerNewJobRequestEvent
	// Filter job transition events.
	FilterJobTransition(
		ch <-chan types.Log,
	) <-chan *metaschedulerabi.MetaSchedulerJobTransitionEvent
}

// AllowanceManager set the allowed quantity of credit for smart-contract interactions.
type AllowanceManager interface {
	// Set the allowance for smart-contract interactions.
	SetAllowance(ctx context.Context, amount *big.Int) error

	// ClearAllowance is an alias to SetAllowance 0.
	ClearAllowance(ctx context.Context) error

	// Get the current allowance toward the contract.
	GetAllowance(ctx context.Context) (*big.Int, error)
}
