// Package job describes the types and methods to interact with the jobs.
package job

import (
	"context"
	"math/big"

	"github.com/deepsquare-io/grid/cli/sbatch"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
)

// SubmitJobOption is used to apply default and optional parameters for submitting a job.
type SubmitJobOption func(*SubmitJobOptions)

// SubmitJobOptions is the object containing optional parameters for submitting a job.
type SubmitJobOptions struct {
	Uses       []types.Label
	Affinities []types.Affinity
}

// WithUse adds strict key-value filters to the job, which filters the available clusters.
func WithUse(labels ...types.Label) SubmitJobOption {
	return func(sjo *SubmitJobOptions) {
		sjo.Uses = labels
	}
}

// WithAffinity adds key-value filters with operators to the job, which filters the available clusters.
func WithAffinity(affinities ...types.Affinity) SubmitJobOption {
	return func(sjo *SubmitJobOptions) {
		sjo.Affinities = affinities
	}
}

// Scheduler schedules and cancels jobs.
type Scheduler interface {
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

// LazyIterator iterates on a lazy list of jobs.
//
// When calling Next or Prev, a request will be sent to the data source.
type LazyIterator interface {
	// Get the current job.
	Current() types.Job
	// Get the current error.
	Error() error
	// Get the current index.
	Index() int
	// Get the current size.
	Size() int

	// Set the error.
	SetError(err error)
	// Increment the index.
	IncrementIndex()
	// Set the current job.
	SetJob(job types.Job)
	// Get the next index.
	GetNextID() [32]byte
}

// Fetcher fetches jobs.
type Fetcher interface {
	// Get a job.
	GetJob(ctx context.Context, id [32]byte) (types.Job, error)
	// Get a iterator of jobs. If there is no job, nil is returned.
	GetJobs(ctx context.Context) (*Iterator, error)

	// Fetches the next job.
	Next(ctx context.Context, it LazyIterator) (ok bool)
}

// MetaScheduledIDsFetcher fetches meta-scheduled jobs ids.
//
// This contacts directly the meta-scheduler without the need to fetch all the jobs.
type MetaScheduledIDsFetcher interface {
	GetMetaScheduledJobIDs(ctx context.Context) ([][32]byte, error)
}

// ByProviderFetcher fetches the jobs meta-scheduled or running on the provider.
type ByProviderFetcher interface {
	GetJobsByProvider(ctx context.Context, providerAddress common.Address) ([]types.Job, error)
}

// Iterator is an iterator of jobs.
type Iterator struct {
	array  [][32]byte
	length int
	index  int
	job    types.Job
	err    error
}

// NewIterator creates a new iterator.
func NewIterator(array [][32]byte) *Iterator {
	return &Iterator{
		array:  array,
		length: len(array),
		index:  -1,
		job:    nil,
	}
}

// GetNextID returns the next job.
func (it *Iterator) GetNextID() [32]byte {
	return it.array[it.index+1]
}

// IncrementIndex increments the index.
func (it *Iterator) IncrementIndex() {
	it.index++
}

// Index returns the current index.
func (it *Iterator) Index() int {
	return it.index
}

// SetError sets the error.
func (it *Iterator) SetError(err error) {
	it.err = err
}

// SetJob sets the current job.
func (it *Iterator) SetJob(job types.Job) {
	it.job = job
}

// Size returns the size of the array.
func (it *Iterator) Size() int {
	return it.length
}

// Current returns the current job.
func (it *Iterator) Current() types.Job {
	return it.job
}

// Error returns the current error.
func (it *Iterator) Error() error {
	return it.err
}
