// Copyright (C) 2023 DeepSquare Asociation
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

package validator_test

import (
	"testing"

	"github.com/deepsquare-io/grid/cli/internal/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsNumber(t *testing.T) {
	tests := []struct {
		input         string
		isError       bool
		errorContains []string
	}{
		{
			input: "1",
		},
		{
			input: "1.0",
		},
		{
			input: "1e10",
		},
		{
			input:   "",
			isError: true,
		},
		{
			input:   "a",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			err := validator.IsNumber(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsMap(t *testing.T) {
	tests := []struct {
		input         string
		isError       bool
		errorContains []string
	}{
		{
			input: "a=b",
		},
		{
			input: "a=b,c=d",
		},
		{
			input: "",
		},
		{
			input:   "a",
			isError: true,
		},
		{
			input:   "a=b,c",
			isError: true,
		},
		{
			input:   "os=linux,arch=amd64,=invalid",
			isError: true,
		},
		{
			input:   "os=linux,arch=amd64,key_with_@=value",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			err := validator.IsMap(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}
