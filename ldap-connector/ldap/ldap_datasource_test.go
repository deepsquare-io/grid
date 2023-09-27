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
