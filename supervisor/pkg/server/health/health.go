package health

import (
	"context"
	"time"

	healthv1 "github.com/deepsquare-io/grid/supervisor/generated/grpc/health/v1"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"go.uber.org/zap"
)

type Server struct {
	healthv1.UnimplementedHealthServer
}

func New() *Server {
	return &Server{}
}

func (h *Server) Check(
	ctx context.Context,
	req *healthv1.HealthCheckRequest,
) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *Server) Watch(
	req *healthv1.HealthCheckRequest,
	stream healthv1.Health_WatchServer,
) error {
	ctx := stream.Context()
	ticker := time.NewTicker(10 * time.Second)

	go func() {
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
	}()

	return nil
}
