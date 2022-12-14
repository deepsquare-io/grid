package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanResources = model.Resources{
	Tasks:       1,
	CpusPerTask: 1,
	MemPerCPU:   1,
	GpusPerTask: 0,
}

func cleanStepWithRun(command string) *model.Step {
	return &model.Step{
		Name: "test",
		Run: &model.StepRun{
			Resources: &cleanResources,
			X11:       utils.Ptr(true),
			Env: []*model.EnvVar{
				{
					Key:   "test",
					Value: "value",
				},
			},
			Image:   utils.Ptr("image"),
			Command: command,
		},
	}
}

func TestRenderStepRun(t *testing.T) {
	tests := []struct {
		input         model.Step
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: *cleanStepWithRun("hostname"),
			expected: `MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'STORAGE_PATH=/deepsquare','test'='value' \
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
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Env:       cleanStepWithRun("").Run.Env,
					Resources: &cleanResources,
					Command:   "hostname",
				},
			},
			expected: `
/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'hostname'`,
			title: "Positive test without image",
		},
		{
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Env:       cleanStepWithRun("").Run.Env,
					Resources: &cleanResources,
					Command: `hostname
echo "test"`,
				},
			},
			expected: `
/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'hostname
echo "test"'`,
			title: "Positive test with multiline command",
		},
		{
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Resources: &cleanResources,
					Image:     utils.Ptr("something???wrong"),
					Command:   "hostname",
				},
			},
			isError:       true,
			errorContains: []string{"Image", "valid_container_image_url"},
			title:         "Negative test with invalid image",
		},
		{
			input: model.Step{
				Name: "test",
				Run: &model.StepRun{
					Resources: &model.Resources{
						Tasks:       0,
						CpusPerTask: 1,
						MemPerCPU:   1,
						GpusPerTask: 0,
					},
					Command: "hostname",
				},
			},
			isError:       true,
			errorContains: []string{"Tasks", "gte"},
			title:         "Negative test with invalid resources",
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
				shellcheck(t, actual)
			}
		})
	}
}

func TestRenderStepFor(t *testing.T) {
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
				For: &model.StepFor{
					Parallel: true,
					Items:    []string{"a", "b", "c"},
					Steps: []*model.Step{
						cleanStepWithRun("echo $item"),
						cleanStepWithRun("echo $item"),
					},
				},
			},
			expected: `doFor() {
export item="$1"
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo $item'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo $item'
}
pids=()
items=('a' 'b' 'c' )
for item in "${items[@]}"; do
  doFor "$item" &
  pids+=("$!")
done
for pid in "${pids[@]}"; do
  wait "$pid"
done`,
			title: "Positive test with items",
		},
		{
			input: model.Step{
				Name: "test",
				For: &model.StepFor{
					Parallel: true,
					Range: &model.ForRange{
						Begin:     0,
						End:       -10,
						Increment: -2,
					},
					Steps: []*model.Step{
						cleanStepWithRun("echo $index"),
						cleanStepWithRun("echo $index"),
					},
				},
			},
			expected: `doFor() {
export index="$1"
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo $index'
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro"
/usr/bin/srun --job-name='test' \
  --export=ALL,"$(loadDeepsquareEnv)",'STORAGE_PATH=/deepsquare','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-mounts="${MOUNTS}" \
  --container-image='image' \
  /bin/sh -c 'echo $index'
}
pids=()
for index in $(seq 0 -2 -10); do
  doFor "$index" &
  pids+=("$!")
done
for pid in "${pids[@]}"; do
  wait "$pid"
done`,
			title: "Positive test with range",
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
				shellcheck(t, actual)
			}
		})
	}
}
