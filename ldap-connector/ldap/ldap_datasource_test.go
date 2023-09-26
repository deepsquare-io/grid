//go:build integration

package ldap_test

import (
	"context"
	"os"
	"testing"

	"github.com/deepsquare-io/grid/ldap-connector/config"
	"github.com/deepsquare-io/grid/ldap-connector/ldap"
	"github.com/deepsquare-io/grid/ldap-connector/logger"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type DataSourceTestSuite struct {
	suite.Suite
	url      string
	bindDN   string
	password string
	config   config.Config
	insecure bool
	caFile   string
	impl     *ldap.DataSource
}

func (suite *DataSourceTestSuite) BeforeTest(suiteName, testName string) {
	suite.impl = ldap.New(
		suite.url,
		suite.bindDN,
		suite.password,
		suite.config,
		suite.insecure,
		suite.caFile,
	)
}

func (suite *DataSourceTestSuite) TestCreateUser() {
	ctx := context.Background()
	err := suite.impl.CreateUser(ctx, "test-user")

	suite.NoError(err)
}

func (suite *DataSourceTestSuite) TestAddUserToGroup() {
	ctx := context.Background()
	err := suite.impl.AddUserToGroup(ctx, "test-user")

	suite.NoError(err)
}

func TestDataSourceTestSuite(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		// Skip test if not defined
		logger.I.Error("Error loading .env.test file", zap.Error(err))
	} else {
		conf, err := config.ParseFile(os.Getenv("CONFIG_PATH"))
		if err != nil {
			logger.I.Fatal("conf failed", zap.Error(err))
		}
		suite.Run(t, &DataSourceTestSuite{
			url:      os.Getenv("LDAP_URL"),
			bindDN:   os.Getenv("LDAP_BIND_DN"),
			password: os.Getenv("LDAP_BIND_PASSWORD"),
			insecure: os.Getenv("LDAP_INSECURE") == "true",
			config:   conf,
		})
	}
}
