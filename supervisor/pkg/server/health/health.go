package health

import (
	"context"
	"time"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/grpc/health/v1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

type health struct {
	healthv1.UnimplementedHealthServer
}

func New() *health {
	return &health{}
}

func (h *health) Check(ctx context.Context, req *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *health) Watch(
	req *healthv1.HealthCheckRequest,
	stream healthv1.Health_WatchServer,
) error {
	ctx := context.Background()
	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)

	go func(ctx context.Context) {
		for {
			if err := stream.Send(&healthv1.HealthCheckResponse{
				Status: healthv1.HealthCheckResponse_SERVING,
			}); err != nil {
				logger.I.Error("healthcheck send failed", zap.Error(err))
			}

			select {
			case <-done:
				return
			case <-ticker.C:
			}
		}
	}(ctx)

	return nil
}
