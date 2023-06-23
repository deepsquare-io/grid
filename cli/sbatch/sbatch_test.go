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

	"github.com/deepsquare-io/the-grid/cli/internal/graphql"
	"github.com/deepsquare-io/the-grid/cli/sbatch"
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
