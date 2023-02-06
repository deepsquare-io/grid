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
		Container: &model.ContainerRun{
			Image:    "image",
			Registry: utils.Ptr("registry"),
			Username: utils.Ptr("username"),
			Password: utils.Ptr("password"),
			X11:      utils.Ptr(true),
			Mounts: []*model.Mount{
				{
					HostDir:      "/host",
					ContainerDir: "/container",
					Options:      "ro",
				},
			},
		},
		Env: []*model.EnvVar{
			{
				Key:   "test",
				Value: "value",
			},
		},
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
			expected: `/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine "registry" login "username" password "password"
EOFnetrc
MOUNTS="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)",'key'='test'\''test','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image='registry#image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with image",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.Container.Apptainer = utils.Ptr(true)
				return r
			}(),
			expected: `export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
export APPTAINER_DOCKER_USERNAME='username'
export APPTAINER_DOCKER_PASSWORD='password'
# shellcheck disable=SC2097,SC2098
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)",'key'='test'\''test','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /usr/bin/apptainer --silent exec \
  --nv \
  'docker://registry/image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with apptainer image",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.Container.Apptainer = utils.Ptr(true)
				r.Container.Image = "/test/my.sqshfs"
				return r
			}(),
			expected: `export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
export APPTAINER_DOCKER_USERNAME='username'
export APPTAINER_DOCKER_PASSWORD='password'
# shellcheck disable=SC2097,SC2098
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)",'key'='test'\''test','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /usr/bin/apptainer --silent exec \
  --nv \
  "$STORAGE_PATH"'/test/my.sqshfs' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with apptainer absolute path image",
		},
		{
			input: func() model.StepRun {
				r := *cleanStepRun("hostname")
				r.Container.DeepsquareHosted = utils.Ptr(true)
				return r
			}(),
			expected: `export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,/tmp/.X11-unix:/tmp/.X11-unix:ro",'/host':'/container':'ro'
export APPTAINER_DOCKER_USERNAME='username'
export APPTAINER_DOCKER_PASSWORD='password'
# shellcheck disable=SC2097,SC2098
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' /usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)",'key'='test'\''test','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --gpu-bind=none \
  /usr/bin/apptainer --silent exec \
  --nv \
  '/opt/software/registry/image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with deepsquare-hosted image",
		},
		{
			input: model.StepRun{
				Env:       cleanStepRun("").Env,
				Resources: &cleanStepRunResources,
				Command: `hostname
/usr/bin/echo "test"`,
			},
			expected: `/usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)",'key'='test'\''test','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  /bin/sh -c 'hostname
/usr/bin/echo "test"'`,
			title: "Positive test with multiline command",
		},
		{
			input: model.StepRun{
				Env:               cleanStepRun("").Env,
				Resources:         &cleanStepRunResources,
				Command:           "hostname",
				DisableCPUBinding: utils.Ptr(true),
			},
			expected: `/usr/bin/srun --job-name='test' \
  --export=ALL"$(loadDeepsquareEnv)",'key'='test'\''test','test'='value' \
  --cpus-per-task=1 \
  --mem-per-cpu=1M \
  --gpus-per-task=0 \
  --ntasks=1 \
  --cpu-bind=none \
  /bin/sh -c 'hostname'`,
			title: "Positive test with disable cpu-bind",
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
				require.NoError(t, renderer.Shellcheck(actual))
			}
		})
	}
}
