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

package metascheduler_test

import (
	"context"
	"testing"

	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/mocks/mockjob"
	"github.com/deepsquare-io/grid/cli/types"
	metaschedulerabi "github.com/deepsquare-io/grid/cli/types/abi/metascheduler"
	"github.com/deepsquare-io/grid/cli/types/job"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MetaScheduledJobsByProviderFetcherTestSuite struct {
	suite.Suite

	fetcher                    *mockjob.Fetcher
	metascheduledJobIDsFetcher *mockjob.MetaScheduledIdsFetcher
	impl                       job.ByProviderFetcher
}

func (suite *MetaScheduledJobsByProviderFetcherTestSuite) BeforeTest(_, _ string) {
	suite.fetcher = mockjob.NewFetcher(suite.T())
	suite.metascheduledJobIDsFetcher = mockjob.NewMetaScheduledIdsFetcher(suite.T())
	suite.impl = metascheduler.NewJobsByProviderFetcher(
		suite.metascheduledJobIDsFetcher,
		suite.fetcher,
	)
}

func (suite *MetaScheduledJobsByProviderFetcherTestSuite) TestGetRunningJobsByProvider() {
	// Arrange
	providerAddress := common.Address{1}
	jobID := metascheduler.JobIDFromHex(
		"0x00000000000000000000000000000000000000000000000000000000000001b4",
	)
	providerAddress2 := common.Address{2}
	jobID2 := metascheduler.JobIDFromHex(
		"0x00000000000000000000000000000000000000000000000000000000000001b4",
	)
	suite.metascheduledJobIDsFetcher.EXPECT().GetMetaScheduledJobIDs(mock.Anything).Return(
		[][32]byte{jobID, jobID2},
		nil,
	)

	suite.fetcher.EXPECT().GetJob(mock.Anything, jobID).Return(&metaschedulerabi.Job{
		JobId:        jobID,
		ProviderAddr: providerAddress,
	}, nil).Once()
	suite.fetcher.EXPECT().GetJob(mock.Anything, jobID).Return(&metaschedulerabi.Job{
		JobId:        jobID2,
		ProviderAddr: providerAddress2,
	}, nil).Once()

	// Act
	jobs, err := suite.impl.GetJobsByProvider(context.Background(), providerAddress)

	// Assert
	suite.Require().NoError(err)
	suite.Require().Equal([]types.Job{
		{
			JobId:        jobID,
			ProviderAddr: providerAddress,
		},
	}, jobs)
}

func TestRunningJobsByProviderFetcher(t *testing.T) {
	suite.Run(t, &MetaScheduledJobsByProviderFetcherTestSuite{})
}
