package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderStep(t *testing.T) {
	tests := []struct {
		input         model.Step
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Resources: &model.Resources{
						Tasks:       1,
						CpusPerTask: 1,
						MemPerCPU:   1,
						GpusPerTask: 0,
					},
					Image:   utils.Ptr("image"),
					Command: "hostname",
				},
			},
			expected: `srun --job-name='test' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-image='image' \
  sh -c 'hostname'
`,
			title: "Positive test with image",
		},
		{
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Resources: &model.Resources{
						Tasks:       1,
						CpusPerTask: 1,
						MemPerCPU:   1,
						GpusPerTask: 0,
					},
					Command: "hostname",
				},
			},
			expected: `srun --job-name='test' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  sh -c 'hostname'
`,
			title: "Positive test without image",
		},
		{
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Resources: &model.Resources{
						Tasks:       1,
						CpusPerTask: 1,
						MemPerCPU:   1,
						GpusPerTask: 0,
					},
					Command: `hostname
echo "test"`,
				},
			},
			expected: `srun --job-name='test' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  sh -c 'hostname
echo "test"'
`,
			title: "Positive test with multiline command",
		},
		{
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Resources: &model.Resources{
						Tasks:       1,
						CpusPerTask: 1,
						MemPerCPU:   1,
						GpusPerTask: 0,
					},
					Image:   utils.Ptr("something???wrong"),
					Command: "hostname",
				},
			},
			isError:       true,
			errorContains: []string{"valid_container_image_url"},
			title:         "Negative test with invalid image",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderStep(&tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
			}
		})
	}
}
