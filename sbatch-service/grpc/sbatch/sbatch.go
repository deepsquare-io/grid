package sbatch

import (
	"context"

	sbatchapiv1alpha1 "github.com/deepsquare-io/the-grid/sbatch-service/gen/go/sbatchapi/v1alpha1"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type API struct {
	sbatchapiv1alpha1.UnimplementedSBatchAPIServer
	RedisClient *redis.Client
}

func NewAPI(r *redis.Client) *API {
	if r == nil {
		logger.I.Panic("redis is nil")
	}
	return &API{
		RedisClient: r,
	}
}

func (a *API) GetSBatch(ctx context.Context, req *sbatchapiv1alpha1.GetSBatchRequest) (*sbatchapiv1alpha1.GetSBatchResponse, error) {
	logger.I.Info("get", zap.String("batchLocationHash", req.BatchLocationHash))
	resp, err := a.RedisClient.Get(ctx, req.BatchLocationHash).Result()
	if err != nil {
		return nil, err
	}
	return &sbatchapiv1alpha1.GetSBatchResponse{Sbatch: resp}, nil
}
