// Copyright (C) 2024 DeepSquare Association
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

	"github.com/deepsquare-io/grid/sbatch-service/renderer"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/stretchr/testify/require"
)

func TestFormatImageURL(t *testing.T) {
	tests := []struct {
		registry         *string
		image            string
		apptainer        *bool
		deepsquareHosted *bool

		title    string
		expected string
	}{
		{
			registry:         utils.Ptr("registry"),
			image:            "image",
			apptainer:        utils.Ptr(false),
			deepsquareHosted: utils.Ptr(false),
			title:            "Positive test: enroot",
			expected:         "docker://registry#image",
		},
		{
			image:            "image",
			apptainer:        utils.Ptr(false),
			deepsquareHosted: utils.Ptr(false),
			title:            "Positive test: enroot",
			expected:         "docker://image",
		},
		{
			registry:         utils.Ptr("registry"),
			image:            "/",
			apptainer:        utils.Ptr(false),
			deepsquareHosted: utils.Ptr(false),
			title:            "Positive test: absolute path",
			expected:         "/",
		},
		{
			registry:         utils.Ptr("registry"),
			image:            "/../../..",
			apptainer:        utils.Ptr(false),
			deepsquareHosted: utils.Ptr(false),
			title:            "Positive test: absolute path traversal attack",
			expected:         "/",
		},
		{
			registry:         utils.Ptr("registry"),
			image:            "image",
			apptainer:        utils.Ptr(true),
			deepsquareHosted: utils.Ptr(false),
			title:            "Positive test: apptainer",
			expected:         "docker://registry/image",
		},
		{
			image:            "image",
			apptainer:        utils.Ptr(true),
			deepsquareHosted: utils.Ptr(false),
			title:            "Positive test: apptainer",
			expected:         "docker://image",
		},
		{
			registry:         utils.Ptr("registry"),
			image:            "image",
			apptainer:        utils.Ptr(true),
			deepsquareHosted: utils.Ptr(true),
			title:            "Positive test: deepsquare hosted",
			expected:         "/opt/software/registry/image",
		},
		{
			image:            "image",
			apptainer:        utils.Ptr(true),
			deepsquareHosted: utils.Ptr(true),
			title:            "Positive test: deepsquare hosted",
			expected:         "/opt/software/image",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			res := renderer.FormatImageURL(tt.registry, tt.image, tt.apptainer, tt.deepsquareHosted)

			// Assert
			require.Equal(t, tt.expected, res)
		})
	}
}
