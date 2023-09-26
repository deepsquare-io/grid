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

package sbatch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "embed"

	"github.com/deepsquare-io/grid/cli/internal/graphql"
	"github.com/deepsquare-io/grid/cli/sbatch"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

//go:embed fixture.yaml
var fixture []byte

type ServiceTestSuite struct {
	suite.Suite
	server *httptest.Server
	impl   sbatch.Service
}

func (suite *ServiceTestSuite) TestSubmit() {
	// Arrange
	suite.server = httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dat, err := io.ReadAll(r.Body)
			suite.Require().NoError(err)
			fmt.Println(string(dat))

			err = json.NewEncoder(w).Encode(graphql.Response[struct {
				Submit string `json:"submit"`
			}]{
				Data: struct {
					Submit string `json:"submit"`
				}{
					Submit: "OK",
				},
			})
			suite.Require().NoError(err)
		}),
	)
	defer suite.server.Close()
	suite.impl = sbatch.NewService(suite.server.Client(), suite.server.URL)

	// Act
	var job sbatch.Job
	err := yaml.Unmarshal(fixture, &job)
	suite.Require().NoError(err)
	respData, err := suite.impl.Submit(context.Background(), &job)

	// Assert
	suite.Require().NoError(err)
	suite.Require().Equal("OK", respData)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, &ServiceTestSuite{})
}
