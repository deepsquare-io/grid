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

//go:build integration

package metascheduler_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/deepsquare-io/grid/cli/metascheduler"
	"github.com/deepsquare-io/grid/cli/types"
	"github.com/stretchr/testify/suite"
)

type OracleTestSuite struct {
	suite.Suite

	impl types.MetaScheduledJobsIdsFetcher
}

func (suite *OracleTestSuite) BeforeTest(_, _ string) {
	suite.impl = metascheduler.NewOracle(
		metascheduler.DefaultOracleURL,
		metascheduler.OracleOptions{},
	)
}

func (suite *OracleTestSuite) TestGetRunningJobIDs() {
	jobs, err := suite.impl.GetMetaScheduledJobIDs(context.Background())

	suite.Require().NoError(err)
	fmt.Println(jobs)
}

func TestOracle(t *testing.T) {
	suite.Run(t, &OracleTestSuite{})
}
