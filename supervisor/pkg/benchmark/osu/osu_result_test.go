package osu_test

import (
	"strings"
	"testing"

	_ "embed"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/osu"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed fixtures/alltoall.log
	fixturesAllToAll string
	//go:embed fixtures/pt2pt-bibw.log
	fixturesPt2PtBiBW string
	//go:embed fixtures/pt2pt-latency.log
	fixturesPt2PtLatency string
)

func TestParseOSULog(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected float64
	}{
		{
			name:     "Real test 1",
			data:     fixturesAllToAll,
			expected: 1884.39,
		},
		{
			name:     "Real test 2",
			data:     fixturesPt2PtBiBW,
			expected: 2314.11,
		},
		{
			name:     "Real test 3",
			data:     fixturesPt2PtLatency,
			expected: 3635.66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := osu.ParseOSULog(strings.NewReader(tt.data))
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}
