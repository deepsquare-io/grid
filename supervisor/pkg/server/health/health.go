package health

import (
	"context"

	healthv1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/grpc/health/v1"
)

type health struct {
	healthv1.UnimplementedHealthServer
}

func New() *health {
	return &health{}
}

func (*health) Check(ctx context.Context, req *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_SERVING,
	}, nil
}

func (*health) Watch(
	req *healthv1.HealthCheckRequest,
	stream healthv1.Health_WatchServer,
) error {
	return stream.Send(&healthv1.HealthCheckResponse{
		Status: healthv1.HealthCheckResponse_SERVING,
	})
}
