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
