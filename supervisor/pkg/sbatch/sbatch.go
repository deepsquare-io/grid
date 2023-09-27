// Copyright (C) 2023 DeepSquare Asociation
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

package sbatch

import (
	"context"
	"time"

	sbatchv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type FetchResponse struct {
	SBatch        string
	GridLoggerURL string
}

type Client interface {
	// HealthCheck sends a simple commands to check if the service is alive.
	HealthCheck(ctx context.Context) error
	// Fetch a job batch content.
	Fetch(ctx context.Context, hash string) (FetchResponse, error)
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
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
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

func (s *client) Fetch(ctx context.Context, hash string) (FetchResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	logger.I.Info("Fetch sbatch", zap.String("hash", hash))
	client, conn, err := s.dial(ctx)
	if err != nil {
		return FetchResponse{}, err
	}
	defer func() {
		_ = conn.Close()
	}()
	resp, err := client.GetSBatch(ctx, &sbatchv1alpha1.GetSBatchRequest{
		BatchLocationHash: hash,
	})
	return FetchResponse{
		SBatch:        resp.GetSbatch(),
		GridLoggerURL: resp.GetGridLoggerUrl(),
	}, err
}
