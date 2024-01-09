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

package validate_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/validate"
	"github.com/stretchr/testify/require"
)

func TestContainerURLValidator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "image/test",
			expected: true,
		},
		{
			input:    "/image/test",
			expected: true,
		},
		{
			input:    "registry-1.docker.io#image/test",
			expected: false,
		},
		{
			input:    "user@registry-1.docker.io#image/test",
			expected: false,
		},
		{
			input:    "user#registry-1.docker.io@image/test",
			expected: false,
		},
		{
			input:    "user#registry-1.docker.io@image/test",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res := validate.ContainerURLValidator(tt.input)
			require.Equal(t, tt.expected, res)
		})
	}
}

func TestCheckContainerImage(t *testing.T) {
	tests := []struct {
		username string
		password string
		registry string
		image    string
		expected error
	}{
		{
			username: "",
			password: "",
			registry: "registry-1.docker.io",
			image:    "curlimages/curl:latest",
			expected: nil,
		},
		{
			username: "",
			password: "",
			registry: "",
			image:    "library/mariadb:latest",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.image, func(t *testing.T) {
			res := validate.CheckContainerImage(
				tt.username,
				tt.password,
				tt.registry,
				tt.image,
			)
			fmt.Println(res)
			require.True(t, errors.Is(res, tt.expected))
		})
	}
}
