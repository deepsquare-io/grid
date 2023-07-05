package module_test

import (
	"context"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/module"
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
			repository:   "github.com/deepsquare-io/workflow-module-example/other-module-example@v1.0.0",
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
