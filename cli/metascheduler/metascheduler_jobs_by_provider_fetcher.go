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
	"bytes"
	"context"

	"github.com/deepsquare-io/grid/cli/types"
	"github.com/deepsquare-io/grid/cli/types/job"
	"github.com/ethereum/go-ethereum/common"
)

type runningJobsByProviderFetcher struct {
	job.MetaScheduledIdsFetcher
	job.Fetcher
}

// NewJobsByProviderFetcher instanciates a JobsByProviderFetcher.
func NewJobsByProviderFetcher(
	oracle job.MetaScheduledIdsFetcher,
	fetcher job.Fetcher,
) job.ByProviderFetcher {
	return &runningJobsByProviderFetcher{
		MetaScheduledIdsFetcher: oracle,
		Fetcher:                 fetcher,
	}
}

func (f *runningJobsByProviderFetcher) GetJobsByProvider(
	ctx context.Context,
	providerAddress common.Address,
) ([]types.Job, error) {
	// Find jobs
	jobIDs, err := f.GetMetaScheduledJobIDs(ctx)
	if err != nil {
		return []types.Job{}, err
	}

	jobs := make([]types.Job, 0, len(jobIDs))

	for _, jobID := range jobIDs {
		job, err := f.GetJob(ctx, jobID)
		if err != nil {
			return []types.Job{}, err
		}

		if bytes.Equal(job.ProviderAddr[:], providerAddress[:]) {
			jobs = append(jobs, job)
		}
	}

	return jobs, nil
}
