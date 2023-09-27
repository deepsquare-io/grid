// Copyright (C) 2023 DeepSquare Association
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

//go:build unit

package health_test

import (
	"context"
	"testing"
	"time"

	healthv1 "github.com/deepsquare-io/grid/supervisor/generated/grpc/health/v1"
	"github.com/deepsquare-io/grid/supervisor/mocks/mockhealthv1"
	"github.com/deepsquare-io/grid/supervisor/pkg/server/health"
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

	mockStream := mockhealthv1.NewHealth_WatchServer(suite.T())
	mockStream.EXPECT().Send(mock.Anything).Return(nil)
	mockStream.EXPECT().Context().Return(ctx)

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
