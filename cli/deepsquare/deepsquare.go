package deepsquare

import (
	"context"
	"math/big"

	loggerv1alpha1 "github.com/deepsquare-io/the-grid/cli/deepsquare/generated/logger/v1alpha1"
)

// LogStream is a readable stream of logs.
type LogStream loggerv1alpha1.LoggerAPI_ReadClient

// Logger fetches the logs of a job.
type Logger interface {
	// Watch the logs of a job
	Watch(ctx context.Context, jobName string) (LogStream, error)
}

// JobSubmitter submits jobs.
type JobSubmitter interface {
	// Submit a batch script to the batch service.
	Submit(ctx context.Context, content string) (string, error)
}

// JobFetcher fetches jobs.
type JobFetcher interface {
	// Get a job.
	GetOne(ctx context.Context)
	// Get multiple jobs.
	GetSome(ctx context.Context)
}

// JobWatcher watches smart-contract events.
type JobWatcher interface {
	// Watch new job requests events.
	WatchNewJobRequests(ctx context.Context)
	// Watch job transition events.
	WatchJobTransition(ctx context.Context)
}

// AllowanceManager set the allowed quantity of credit for smart-contract interactions.
type AllowanceManager interface {
	// Set the allowance for smart-contract interactions.
	SetAllowance(ctx context.Context, amount *big.Int)

	// ClearAllowance is an alias to SetAllowance 0.
	ClearAllowance(ctx context.Context)
}
