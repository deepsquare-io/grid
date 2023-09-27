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

package health

import (
	"context"
	"time"

	healthv1 "github.com/deepsquare-io/grid/grid-logger/gen/go/grpc/health/v1"
	"github.com/deepsquare-io/grid/grid-logger/logger"
	"go.uber.org/zap"
)

type health struct {
	healthv1.UnimplementedHealthServer
}

func New() *health {
	return &health{}
}

func (h *health) Check(
	ctx context.Context,
	req *healthv1.HealthCheckRequest,
) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *health) Watch(
	req *healthv1.HealthCheckRequest,
	stream healthv1.Health_WatchServer,
) error {
	ctx := stream.Context()
	ticker := time.NewTicker(10 * time.Second)

	go func(ctx context.Context) {
		for {
			if err := stream.Send(&healthv1.HealthCheckResponse{
				Status: healthv1.HealthCheckResponse_SERVING,
			}); err != nil {
				logger.I.Error("healthcheck send failed", zap.Error(err))
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
		}
	}(ctx)

	return nil
}
