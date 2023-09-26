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
