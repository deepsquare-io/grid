package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanApptainerStepRun(command string) *model.StepRun {
	return &model.StepRun{
		Resources: &cleanStepRunResources,
		Container: &model.ContainerRun{
			Image:     "image",
			Registry:  utils.Ptr("registry"),
			Username:  utils.Ptr("username"),
			Password:  utils.Ptr("password"),
			X11:       utils.Ptr(true),
			Apptainer: utils.Ptr(true),
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

func TestRenderApptainer(t *testing.T) {
	tests := []struct {
		input         model.StepRun
		isError       bool
		errorContains []string
		expected      string
		title         string
	}{
		{
			input: *cleanApptainerStepRun("hostname"),
			expected: `/usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  'docker://registry/image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with apptainer image",
		},
		{
			input: func() model.StepRun {
				r := *cleanApptainerStepRun("hostname")
				r.MapRoot = utils.Ptr(true)
				return r
			}(),
			expected: `/usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  --fakeroot \
  'docker://registry/image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with maproot",
		},
		{
			input: func() model.StepRun {
				r := *cleanApptainerStepRun("hostname")
				r.Container.Image = "/test/my.sqshfs"
				return r
			}(),
			expected: `/usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
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
			expected: `/usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  '/opt/software/registry/image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with deepsquare-hosted image",
		},
		{
			input: func() model.StepRun {
				r := *cleanApptainerStepRun("hostname")
				r.WorkDir = utils.Ptr("/dir")
				return r
			}(),
			expected: `/usr/bin/apptainer --silent exec \
  --disable-cache \
  --no-home \
  --nv \
  --pwd '/dir' \
  'docker://registry/image' \
  /bin/sh -c 'hostname'`,
			title: "Positive test with workdir",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			actual, err := renderer.RenderApptainerCommand(&tt.input)

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
