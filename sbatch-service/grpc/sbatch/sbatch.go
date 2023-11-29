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
	"bytes"
	"context"
	"errors"

	metaschedulerabi "github.com/deepsquare-io/grid/sbatch-service/abi/metascheduler"
	"github.com/deepsquare-io/grid/sbatch-service/auth"
	sbatchapiv1alpha1 "github.com/deepsquare-io/grid/sbatch-service/gen/go/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/grid/sbatch-service/logger"
	"github.com/deepsquare-io/grid/sbatch-service/storage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type API struct {
	sbatchapiv1alpha1.UnimplementedSBatchAPIServer
	auth           *auth.Auth
	jobs           *metaschedulerabi.IJobRepository
	storage        storage.Storage
	loggerEndpoint string
}

func NewAPI(
	auth *auth.Auth, storage storage.Storage,
	jobs *metaschedulerabi.IJobRepository,
	loggerEndpoint string,
) *API {
	if auth == nil {
		logger.I.Panic("auth is nil")
	}
	if storage == nil {
		logger.I.Panic("storage is nil")
	}
	return &API{
		storage:        storage,
		jobs:           jobs,
		auth:           auth,
		loggerEndpoint: loggerEndpoint,
	}
}

func (a *API) GetSBatch(
	ctx context.Context,
	req *sbatchapiv1alpha1.GetSBatchRequest,
) (*sbatchapiv1alpha1.GetSBatchResponse, error) {
	logger.I.Info(
		"get",
		zap.String("batchLocationHash", req.BatchLocationHash),
		zap.String("gridLoggerURL", a.loggerEndpoint),
	)

	// Check origin
	if err := a.auth.Verify(ctx, hexutil.Encode(req.ProviderAddress), req.Challenge, req.SignedChallenge); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// Check existing job
	var idB [32]byte
	idS := hexutil.MustDecode(req.JobId)
	copy(idB[:], idS)
	job, err := a.jobs.Get(&bind.CallOpts{Context: ctx}, idB)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !bytes.EqualFold(job.CustomerAddr[:], req.CustomerAddress[:]) {
		return nil, status.Error(codes.Unauthenticated, "customer is not the same")
	}

	if !bytes.EqualFold(job.ProviderAddr[:], req.ProviderAddress[:]) {
		return nil, status.Error(codes.Unauthenticated, "provider is not the same")
	}

	resp, err := a.storage.Get(ctx, req.BatchLocationHash)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &sbatchapiv1alpha1.GetSBatchResponse{Sbatch: resp, GridLoggerUrl: a.loggerEndpoint}, nil
}
