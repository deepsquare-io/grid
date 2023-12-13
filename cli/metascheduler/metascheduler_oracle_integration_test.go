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
