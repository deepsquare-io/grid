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
