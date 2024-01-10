// Copyright (C) 2024 DeepSquare Association
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

package metascheduler

import (
	"context"
	"fmt"

	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
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
	types.JobFetcher
	array  [][32]byte
	length int
	index  int
	job    types.Job
	err    error
}

func (it *jobIterator) Next(
	ctx context.Context,
) (ok bool) {
	if it.index+1 >= it.length {
		return false
	}
	job, err := it.GetJob(ctx, it.array[it.index+1])
	if err != nil {
		it.err = fmt.Errorf("failed to get job: %w", err)
		return false
	}

	it.index++
	it.job = job

	return true
}

func (it *jobIterator) Prev(
	ctx context.Context,
) (ok bool) {
	if it.index-1 < 0 {
		return false
	}
	job, err := it.GetJob(ctx, it.array[it.index-1])
	if err != nil {
		it.err = fmt.Errorf("failed to get job: %w", err)
		return false
	}

	it.index--
	it.job = job

	return true
}

func (it *jobIterator) Current() types.Job {
	return it.job
}

func (it *jobIterator) Error() error {
	return it.err
}

func (c *jobFetcher) GetJobs(ctx context.Context) (types.JobLazyIterator, error) {
	jobIDs, err := c.GetByCustomer(&bind.CallOpts{
		Context: ctx,
	}, c.from())
	if err != nil {
		return nil, WrapError(err)
	}
	// Reverse array, from new to old
	for i, j := 0, len(jobIDs)-1; i < j; i, j = i+1, j-1 {
		jobIDs[i], jobIDs[j] = jobIDs[j], jobIDs[i]
	}
	return &jobIterator{
		JobFetcher: c,
		array:      jobIDs,
		length:     len(jobIDs),
		index:      -1,
		job:        nil,
	}, nil
}
