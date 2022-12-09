package renderer_test

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func shellcheck(t *testing.T, script string) {
	_, err := exec.LookPath("shellcheck")
	if err != nil {
		logger.I.Warn("shellcheck is disabled, test is not complete")
		return
	}
	if err := os.WriteFile("test.sh", []byte(script), 0o777); err != nil {
		logger.I.Panic("failed to write", zap.Error(err))
	}
	out, err := exec.Command("shellcheck", "-S", "warning", "-s", "bash", "test.sh").CombinedOutput()
	if err != nil {
		logger.I.Error(string(out))
		require.NoError(t, errors.New("shellcheck failed"))
	}

	_ = os.Remove("test.sh")
}

var cleanJob = model.Job{
	Env: []*model.EnvVar{
		{
			Key:   "key",
			Value: "test'test",
		},
	},
	Steps: []*model.Step{
		cleanStepWithRun("echo 'hello world'"),
	},
}

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

func TestRenderJob(t *testing.T) {
	tests := []struct {
		input         model.Job
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: cleanJob,
			expected: `set -e

export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
mkdir -p "$STORAGE_PATH"
chmod -R 700 "$STORAGE_PATH"
chown -R "$UID:cluster-users" "$STORAGE_PATH"
export 'key'='test'\''test'
srun --job-name='test' \
  --export=ALL,'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-image='image' \
  sh -c 'echo '\''hello world'\'''
`,
			title: "Positive test 'hello world'",
		},
		{
			input: model.Job{
				Env: []*model.EnvVar{
					{
						Key:   "aze'aze",
						Value: "test'test",
					},
				},
			},
			isError:       true,
			errorContains: []string{"valid_envvar_name", "Key"},
			title:         "Negative test invalid env var name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderJob(&tt.input)

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
			expected: `srun --job-name='test' \
  --export=ALL,'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  --container-image='image' \
  sh -c 'hostname'`,
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
			expected: `srun --job-name='test' \
  --export=ALL,'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  sh -c 'hostname'`,
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
			expected: `srun --job-name='test' \
  --export=ALL,'test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1 \
  --gpus-per-task=0 \
  --ntasks=1 \
  sh -c 'hostname
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
  srun --job-name='test' \
    --export=ALL,'test'='value' \
    --cpus-per-task=1 \
    --mem-per-cpu=1 \
    --gpus-per-task=0 \
    --ntasks=1 \
    --container-image='image' \
    sh -c 'echo $item'
  srun --job-name='test' \
    --export=ALL,'test'='value' \
    --cpus-per-task=1 \
    --mem-per-cpu=1 \
    --gpus-per-task=0 \
    --ntasks=1 \
    --container-image='image' \
    sh -c 'echo $item'
}
items=('a' 'b' 'c' )
for item in "${items[@]}"; do
doFor "$item" &
done
wait`,
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
  srun --job-name='test' \
    --export=ALL,'test'='value' \
    --cpus-per-task=1 \
    --mem-per-cpu=1 \
    --gpus-per-task=0 \
    --ntasks=1 \
    --container-image='image' \
    sh -c 'echo $index'
  srun --job-name='test' \
    --export=ALL,'test'='value' \
    --cpus-per-task=1 \
    --mem-per-cpu=1 \
    --gpus-per-task=0 \
    --ntasks=1 \
    --container-image='image' \
    sh -c 'echo $index'
}
for index in $(seq 0 -2 -10); do
doFor "$index" &
done
wait`,
			title: "Positive test with items",
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
