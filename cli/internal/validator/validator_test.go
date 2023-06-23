package validator_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/cli/internal/validator"
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
