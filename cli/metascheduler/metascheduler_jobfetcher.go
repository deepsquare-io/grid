package metascheduler

import (
	"context"
	"fmt"

	"github.com/deepsquare-io/the-grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/the-grid/cli/types/abi/metascheduler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type jobFetcher struct {
	*RPCClientSet
	*metaschedulerabi.IJobRepository
}

func (c *jobFetcher) GetJob(ctx context.Context, id [32]byte) (types.Job, error) {
	job, err := c.Get(&bind.CallOpts{
		Context: ctx,
	}, id)
	if err != nil {
		return nil, WrapError(err)
	}
	return &job, nil
}

type jobIterator struct {
	*jobFetcher
	array  [][32]byte
	length int
	index  int
	job    types.Job
}

func (it *jobIterator) Next(
	ctx context.Context,
) (next types.JobLazyIterator, ok bool, err error) {
	if it.index+1 >= it.length {
		return nil, false, nil
	}
	job, err := it.GetJob(ctx, it.array[it.index+1])
	if err != nil {
		return nil, false, fmt.Errorf("failed to get job: %w", err)
	}

	return &jobIterator{
		jobFetcher: it.jobFetcher,
		array:      it.array,
		length:     it.length,
		index:      it.index + 1,
		job:        job,
	}, true, nil
}

func (it *jobIterator) Prev(
	ctx context.Context,
) (prev types.JobLazyIterator, ok bool, err error) {
	if it.index-1 < 0 {
		return nil, false, nil
	}
	job, err := it.GetJob(ctx, it.array[it.index-1])
	if err != nil {
		return nil, false, fmt.Errorf("failed to get job: %w", err)
	}

	return &jobIterator{
		jobFetcher: it.jobFetcher,
		array:      it.array,
		length:     it.length,
		index:      it.index - 1,
		job:        job,
	}, true, nil
}

func (it *jobIterator) Current() types.Job {
	return it.job
}

func (c *jobFetcher) GetJobs(ctx context.Context) (types.JobLazyIterator, error) {
	jobIDs, err := c.GetByCustomer(&bind.CallOpts{
		Context: ctx,
	}, c.from())
	if err != nil {
		return nil, WrapError(err)
	}
	if len(jobIDs) == 0 {
		return nil, nil
	}
	// Reverse array, from new to old
	for i, j := 0, len(jobIDs)-1; i < j; i, j = i+1, j-1 {
		jobIDs[i], jobIDs[j] = jobIDs[j], jobIDs[i]
	}
	// Initialize first data
	job, err := c.GetJob(ctx, jobIDs[0])
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %w", err)
	}
	return &jobIterator{
		jobFetcher: c,
		array:      jobIDs,
		length:     len(jobIDs),
		index:      0,
		job:        job,
	}, nil
}
