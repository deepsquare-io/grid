//go:generate go run generate.go

package graph

import (
	"github.com/deepsquare-io/grid/sbatch-service/logger"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/deepsquare-io/grid/sbatch-service/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Debug       bool
	Storage     storage.Storage
	JobRenderer *renderer.JobRenderer
}

func NewResolver(storage storage.Storage, renderer *renderer.JobRenderer) *Resolver {
	if storage == nil {
		logger.I.Panic("storage is nil")
	}
	if renderer == nil {
		logger.I.Panic("renderer is nil")
	}
	return &Resolver{
		Storage:     storage,
		JobRenderer: renderer,
	}
}
