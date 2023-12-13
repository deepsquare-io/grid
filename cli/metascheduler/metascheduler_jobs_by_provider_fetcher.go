package metascheduler

import (
	"bytes"
	"context"

	"github.com/deepsquare-io/grid/cli/types"
	"github.com/ethereum/go-ethereum/common"
)

type runningJobsByProviderFetcher struct {
	types.MetaScheduledJobsIdsFetcher
	types.JobFetcher
}

func NewRunningJobsByProviderFetcher(
	oracle types.MetaScheduledJobsIdsFetcher,
	fetcher types.JobFetcher,
) types.JobsByProviderFetcher {
	return &runningJobsByProviderFetcher{
		MetaScheduledJobsIdsFetcher: oracle,
		JobFetcher:                  fetcher,
	}
}

func (f *runningJobsByProviderFetcher) GetJobsByProvider(
	ctx context.Context,
	providerAddress common.Address,
) ([]types.Job, error) {
	// Find jobs
	jobIDs, err := f.GetMetaScheduledJobIDs(ctx)
	if err != nil {
		return nil, err
	}

	jobs := make([]types.Job, 0, len(jobIDs))

	for _, jobID := range jobIDs {
		job, err := f.GetJob(ctx, jobID)
		if err != nil {
			return nil, err
		}

		if bytes.Equal(job.ProviderAddr[:], providerAddress[:]) {
			jobs = append(jobs, job)
		}
	}

	return jobs, nil
}
