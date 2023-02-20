package validate_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
	"github.com/stretchr/testify/require"
)

func TestEnvVarNameValidator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "test",
			expected: true,
		},
		{
			input:    "test test",
			expected: false,
		},
		{
			input:    "test'test",
			expected: false,
		},
		{
			input:    "test\"test",
			expected: false,
		},
		{
			input:    "test\\test",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res := validate.EnvVarNameValidator(tt.input)
			require.Equal(t, tt.expected, res)
		})
	}
}
