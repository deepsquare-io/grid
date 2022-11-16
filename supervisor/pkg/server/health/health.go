package health

import (
	"context"
	"time"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/grpc/health/v1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
	"go.uber.org/zap"
)

type health struct {
	healthv1.UnimplementedHealthServer
	slurm *slurm.Service
}

func New(slurm *slurm.Service) *health {
	return &health{
		slurm: slurm,
	}
}

func (h *health) Check(ctx context.Context, req *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	if err := h.slurm.HealthCheck(ctx); err != nil {
		return &healthv1.HealthCheckResponse{
			Status: healthv1.HealthCheckResponse_NOT_SERVING,
		}, nil
	}
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
			if err := h.slurm.HealthCheck(ctx); err != nil {
				if err := stream.Send(&healthv1.HealthCheckResponse{
					Status: healthv1.HealthCheckResponse_NOT_SERVING,
				}); err != nil {
					logger.I.Error("healthcheck send failed", zap.Error(err))
				}
			} else {
				if err := stream.Send(&healthv1.HealthCheckResponse{
					Status: healthv1.HealthCheckResponse_SERVING,
				}); err != nil {
					logger.I.Error("healthcheck send failed", zap.Error(err))
				}
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
