package metascheduler_test

import (
	"testing"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"
	"github.com/stretchr/testify/require"
)

func TestProcessLabels(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"my_invalid-label123!", "my_invalid-label123"},
		{"validLabel123", "validlabel123"},
		{".a@b#c$d%e^f&g.f", ".a-b-c-d-e-f-g.f"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := metascheduler.ProcessLabel(metaschedulerabi.Label{
				Key:   "",
				Value: tc.input,
			})
			require.Equal(t, tc.expected, result.Value)
		})
	}
}
