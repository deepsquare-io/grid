//go:generate go run generate.go

package graph

import (
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/redis/go-redis/v9"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Debug       bool
	RedisClient *redis.Client
	JobRenderer *renderer.JobRenderer
}

func NewResolver(redis *redis.Client, renderer *renderer.JobRenderer) *Resolver {
	if redis == nil {
		logger.I.Panic("redis is nil")
	}
	if renderer == nil {
		logger.I.Panic("renderer is nil")
	}
	return &Resolver{
		RedisClient: redis,
		JobRenderer: renderer,
	}
}
