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

package sbatch

import (
	"context"
	"crypto/ecdsa"
	"time"

	sbatchv1alpha1 "github.com/deepsquare-io/grid/supervisor/generated/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/grid/supervisor/logger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	Fetch(
		ctx context.Context,
		hash string,
		customerAddress common.Address,
		jobID [32]byte,
	) (FetchResponse, error)
}

func NewClient(
	endpoint string,
	pk *ecdsa.PrivateKey,
	opts ...grpc.DialOption,
) Client {
	return &client{
		endpoint: endpoint,
		opts:     opts,
		pk:       pk,
		pub:      &pk.PublicKey,
	}
}

type client struct {
	endpoint string
	opts     []grpc.DialOption
	pk       *ecdsa.PrivateKey
	pub      *ecdsa.PublicKey
}

func (s *client) dial(
	ctx context.Context,
) (sbatchv1alpha1.SBatchAPIClient, sbatchv1alpha1.AuthAPIClient, *grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, s.endpoint, s.opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	return sbatchv1alpha1.NewSBatchAPIClient(conn), sbatchv1alpha1.NewAuthAPIClient(conn), conn, nil
}

func (s *client) HealthCheck(ctx context.Context) error {
	logger.I.Info("Healthcheck sbatch service")
	_, _, conn, err := s.dial(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	return nil
}

func (s *client) Fetch(
	ctx context.Context,
	hash string,
	customerAddress common.Address,
	jobID [32]byte,
) (FetchResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	logger.I.Info("Fetch sbatch", zap.String("hash", hash))
	client, auth, conn, err := s.dial(ctx)
	if err != nil {
		return FetchResponse{}, err
	}
	defer func() {
		_ = conn.Close()
	}()

	challenge, err := auth.Challenge(ctx, &sbatchv1alpha1.ChallengeRequest{})
	if err != nil {
		return FetchResponse{}, err
	}
	signed, err := crypto.Sign(challenge.GetChallenge(), s.pk)
	if err != nil {
		return FetchResponse{}, err
	}

	addr := crypto.PubkeyToAddress(*s.pub)
	resp, err := client.GetSBatch(ctx, &sbatchv1alpha1.GetSBatchRequest{
		BatchLocationHash: hash,
		SignedChallenge:   signed,
		Challenge:         challenge.GetChallenge(),
		ProviderAddress:   addr[:],
		CustomerAddress:   customerAddress[:],
		JobId:             hexutil.Encode(jobID[:]),
	})
	return FetchResponse{
		SBatch:        resp.GetSbatch(),
		GridLoggerURL: resp.GetGridLoggerUrl(),
	}, err
}
