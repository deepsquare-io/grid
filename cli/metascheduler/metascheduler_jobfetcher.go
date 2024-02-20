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
	"github.com/deepsquare-io/grid/cli/types/job"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var _ job.Fetcher = (*JobFetcher)(nil)

// JobFetcher is a fetcher for jobs.
type JobFetcher struct {
	*RPCClientSet
	*metaschedulerabi.IJobRepository
}

// GetJob returns a job by its ID.
func (c *JobFetcher) GetJob(ctx context.Context, id [32]byte) (types.Job, error) {
	job, err := c.Get(&bind.CallOpts{
		Context: ctx,
	}, id)
	if err != nil {
		return nil, WrapError(err)
	}
	return &job, nil
}

// Next returns the next job in the iterator.
func (c *JobFetcher) Next(ctx context.Context, it job.LazyIterator) (ok bool) {
	if it.Index()+1 >= it.Size() {
		return false
	}
	job, err := c.GetJob(ctx, it.GetNextID())
	if err != nil {
		it.SetError(fmt.Errorf("failed to get job: %w", err))
		return false
	}

	it.IncrementIndex()
	it.SetJob(job)

	return true
}

// GetJobs returns all jobs of the user.
func (c *JobFetcher) GetJobs(ctx context.Context) (*job.Iterator, error) {
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
	return job.NewIterator(jobIDs), nil
}
