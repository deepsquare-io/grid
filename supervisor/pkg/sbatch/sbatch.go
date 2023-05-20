package sbatch

import (
	"context"

	sbatchv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/generated/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Client interface {
	// HealthCheck sends a simple commands to check if the service is alive.
	HealthCheck(ctx context.Context) error
	// Fetch a job batch content.
	Fetch(ctx context.Context, hash string) (string, error)
}

func NewClient(
	endpoint string,
	opts ...grpc.DialOption,
) Client {
	return &client{
		endpoint: endpoint,
		opts:     opts,
	}
}

type client struct {
	endpoint string
	opts     []grpc.DialOption
}

func (s *client) dial(
	ctx context.Context,
) (sbatchv1alpha1.SBatchAPIClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx, s.endpoint, s.opts...)
	if err != nil {
		return nil, nil, err
	}
	return sbatchv1alpha1.NewSBatchAPIClient(conn), conn, nil
}

func (s *client) HealthCheck(ctx context.Context) error {
	logger.I.Info("Healthcheck sbatch service")
	_, conn, err := s.dial(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	return nil
}

func (s *client) Fetch(ctx context.Context, hash string) (string, error) {
	logger.I.Info("Fetch sbatch", zap.String("hash", hash))
	client, conn, err := s.dial(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = conn.Close()
	}()
	resp, err := client.GetSBatch(ctx, &sbatchv1alpha1.GetSBatchRequest{
		BatchLocationHash: hash,
	})
	return resp.GetSbatch(), err
}
