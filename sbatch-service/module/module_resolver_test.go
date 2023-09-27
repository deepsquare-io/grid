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

package module_test

import (
	"context"
	"testing"
	"time"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/module"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolve(t *testing.T) {
	tests := []struct {
		repository    string
		ref           string
		expectedName  string
		isError       bool
		errorContains []string
		title         string
	}{
		{
			repository:   "github.com/deepsquare-io/workflow-module-example",
			expectedName: "Hello World",
			title:        "Positive test: no ref",
		},
		{
			repository:   "github.com/deepsquare-io/workflow-module-example",
			ref:          "6abe5d5980b46bcad807b571a3feeae86600e204",
			expectedName: "Hello World",
			title:        "Positive test: commit hash",
		},
		{
			repository:   "github.com/deepsquare-io/workflow-module-example",
			ref:          "6abe5d5",
			expectedName: "Hello World",
			title:        "Positive test: short commit hash",
		},
		{
			repository:   "github.com/deepsquare-io/workflow-module-example",
			ref:          "v0.1.0",
			expectedName: "Hello World",
			title:        "Positive test: tag",
		},
		{
			repository:   "github.com/deepsquare-io/workflow-module-example/other-module-example",
			expectedName: "Another Hello World",
			title:        "Positive test: submodule",
		},
		{
			repository:   "github.com/deepsquare-io/workflow-module-example/other-module-example",
			ref:          "v1.0.0",
			expectedName: "Another Hello World",
			title:        "Positive test: submodule",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			mod, err := module.Resolve(ctx, &model.Job{
				Resources: &model.JobResources{
					MemPerCPU: 100,
				},
			}, &model.Step{}, tt.repository, tt.ref)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedName, mod.Name)
			}
		})
	}
}
