package validate_test

import (
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
