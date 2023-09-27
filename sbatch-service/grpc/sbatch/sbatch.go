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
	"errors"

	sbatchapiv1alpha1 "github.com/deepsquare-io/grid/sbatch-service/gen/go/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/grid/sbatch-service/logger"
	"github.com/deepsquare-io/grid/sbatch-service/storage"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type API struct {
	sbatchapiv1alpha1.UnimplementedSBatchAPIServer
	storage        storage.Storage
	loggerEndpoint string
}

func NewAPI(storage storage.Storage, loggerEndpoint string) *API {
	if storage == nil {
		logger.I.Panic("storage is nil")
	}
	return &API{
		storage:        storage,
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
	resp, err := a.storage.Get(ctx, req.BatchLocationHash)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}
	return &sbatchapiv1alpha1.GetSBatchResponse{Sbatch: resp, GridLoggerUrl: a.loggerEndpoint}, nil
}
