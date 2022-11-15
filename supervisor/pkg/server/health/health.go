package health

import (
	"context"
	"time"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/grpc/health/v1"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/slurm"
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
				stream.Send(&healthv1.HealthCheckResponse{
					Status: healthv1.HealthCheckResponse_NOT_SERVING,
				})
			} else {
				stream.Send(&healthv1.HealthCheckResponse{
					Status: healthv1.HealthCheckResponse_SERVING,
				})
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
