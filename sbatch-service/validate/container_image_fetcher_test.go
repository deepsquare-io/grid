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
	"errors"
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/validate"
	"github.com/stretchr/testify/require"
)

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
			registry: "registry-1.deepsquare.run",
			image:    "library/ubuntu:latest",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.image, func(t *testing.T) {
			_, res := validate.DefaultImageFetcher.FetchContainerImage(
				tt.username,
				tt.password,
				tt.registry,
				tt.image,
			)
			require.True(t, errors.Is(res, tt.expected))
		})
	}
}
