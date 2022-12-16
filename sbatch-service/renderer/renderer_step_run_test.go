package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanStepRunResources = model.StepRunResources{
	Tasks:       utils.Ptr(1),
	CpusPerTask: utils.Ptr(1),
	MemPerCPU:   utils.Ptr(1),
	GpusPerTask: utils.Ptr(0),
}

func cleanStepRun(command string) *model.StepRun {
	return &model.StepRun{
		Resources: &cleanStepRunResources,
		X11:       utils.Ptr(true),
		Env: []*model.EnvVar{
			{
				Key:   "test",
				Value: "value",
			},
		},
		Image:   utils.Ptr("image"),
		Command: command,
	}
}

func TestRenderStepRun(t *testing.T) {
	tests := []struct {
		input         model.StepRun
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: *cleanStepRun("hostname"),
			expected: `MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with image",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				Command:   "hostname",
			},
			expected: `/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'hostname'`,
			title: "Positive test without image",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				Command: `hostname
/usr/bin/echo "test"`,
			},
			expected: `/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'hostname
/usr/bin/echo "test"'`,
			title: "Positive test with multiline command",
		},
		{
			input: model.StepRun{
				Resources: &cleanStepRunResources,
				Image:     utils.Ptr("something???wrong"),
				Command:   "hostname",
			},
			isError:       true,
			errorContains: []string{"Image", "valid_container_image_url"},
			title:         "Negative test with invalid image",
		},
		{
			input: model.StepRun{
				Resources: &model.StepRunResources{
					Tasks:       utils.Ptr(0),
					CpusPerTask: utils.Ptr(1),
					MemPerCPU:   utils.Ptr(1),
					GpusPerTask: utils.Ptr(0),
				},
				Command: "hostname",
			},
			isError:       true,
			errorContains: []string{"Tasks", "gte"},
			title:         "Negative test with invalid resources",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderStepRun(&cleanJob, cleanStepWithRun(&tt.input))

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
				shellcheck(t, actual)
			}
		})
	}
}
