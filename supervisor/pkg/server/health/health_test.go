//go:build unit

package health_test

import (
	"context"
	"testing"
	"time"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/grpc/health/v1"
	"github.com/deepsquare-io/the-grid/supervisor/mocks"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/server/health"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HealthTestSuite struct {
	suite.Suite
	impl *health.Server
}

func (suite *HealthTestSuite) BeforeTest(suiteName, testName string) {
	suite.impl = health.New()
}

func (suite *HealthTestSuite) TestCheck() {
	ctx := context.Background()

	// Act
	resp, err := suite.impl.Check(ctx, &healthv1.HealthCheckRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(&healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_SERVING,
	}, resp)
}

func (suite *HealthTestSuite) TestWatch() {
	ctx, cancel := context.WithCancel(context.Background())

	mockStream := mocks.NewHealth_WatchServer(suite.T())
	mockStream.On("Send", mock.Anything).Return(nil)
	mockStream.On("Context").Return(ctx)

	// Act
	err := suite.impl.Watch(&healthv1.HealthCheckRequest{}, mockStream)
	suite.Require().NoError(err)

	time.Sleep(time.Second)

	cancel()

	mockStream.AssertExpectations(suite.T())
}

func TestHealthTestSuite(t *testing.T) {
	suite.Run(t, &HealthTestSuite{})
}