// Copyright (C) 2023 DeepSquare Association
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

package renderer_test

import (
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
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
  --contain \
  --writable-tmpfs \
  --no-home \
  --nv \
  --pwd "/deepsquare" \
  "$IMAGE_PATH" \
  /bin/sh -c 'hostname'`,
			title: "Positive test with apptainer image",
		},
		{
			input: func() model.StepRun {
				r := *cleanApptainerStepRun("hostname")
				r.Container.Image = "/test/my.sqshfs"
				return r
			}(),
			expected: `/usr/bin/apptainer --silent exec \
  --disable-cache \
  --contain \
  --writable-tmpfs \
  --no-home \
  --nv \
  --pwd "/deepsquare" \
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
  --contain \
  --writable-tmpfs \
  --no-home \
  --nv \
  --pwd "/deepsquare" \
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
  --contain \
  --writable-tmpfs \
  --no-home \
  --nv \
  --pwd '/dir' \
  "$IMAGE_PATH" \
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
