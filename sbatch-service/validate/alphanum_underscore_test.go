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

package validate_test

import (
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/validate"
	"github.com/stretchr/testify/require"
)

func TestAlphaNumUnderscoreValidator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "image_test",
			expected: true,
		},
		{
			input:    "IMAGE_test",
			expected: true,
		},
		{
			input:    "image/test",
			expected: false,
		},
		{
			input:    "image-test",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res := validate.AlphaNumUnderscoreValidator(tt.input)
			require.Equal(t, tt.expected, res)
		})
	}
}
