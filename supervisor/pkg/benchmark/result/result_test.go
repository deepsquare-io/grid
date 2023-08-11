package result_test

import (
	"strings"
	"testing"
	"time"

	_ "embed"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/benchmark/result"
	"github.com/stretchr/testify/require"
)

//go:embed fixtures/2n2gpu16cpuucxrdma.log
var fixture string

func TestFindMaxGflopsResult(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected *result.Result
	}{
		{
			name: "Valid test",
			data: `HPL_AI WRC01 95000 64 2 2 29.67 1.927e+04 5.71402 2 1.616e+04
HPL_AI WRC01 95000 128 2 2 15.67 3.647e+04 5.72738 2 2.671e+04
HPL_AI WRC01 95000 224 2 2 19.00 3.009e+04 5.74521 2 2.310e+04
HPL_AI WRC01 95000 256 2 2 17.30 3.304e+04 5.71711 2 2.483e+04
HPL_AI WRC01 95000 384 2 2 14.75 3.876e+04 5.77248 2 2.785e+04
HPL_AI WRC01 95000 512 2 2 14.93 3.828e+04 5.76942 2 2.761e+04`,
			expected: &result.Result{
				ProblemSize:       95000,
				NB:                384,
				P:                 2,
				Q:                 2,
				Time:              time.Duration(14.75 * float64(time.Second)),
				Gflops:            3.876e4,
				Refine:            5.77248,
				Iterations:        2,
				GflopsWRefinement: 2.785e4,
			},
		},
		{
			name: "Real test",
			data: fixture,
			expected: &result.Result{
				ProblemSize:       0x29bf8,
				NB:                0x80,
				P:                 2,
				Q:                 3,
				Time:              time.Duration(46810000000),
				Gflops:            71220,
				Refine:            6.64575,
				Iterations:        2,
				GflopsWRefinement: 62370,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputReader := strings.NewReader(tt.data)
			got, err := result.FindMaxGflopsResult(result.NewReader(inputReader))
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}

func TestComputeAvgGflopsResult(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected float64
	}{
		{
			name: "Valid test",
			data: `HPL_AI WRC01 95000 64 2 2 29.67 1.927e+04 5.71402 2 1.616e+04
HPL_AI WRC01 95000 128 2 2 15.67 3.647e+04 5.72738 2 2.671e+04
HPL_AI WRC01 95000 224 2 2 19.00 3.009e+04 5.74521 2 2.310e+04
HPL_AI WRC01 95000 256 2 2 17.30 3.304e+04 5.71711 2 2.483e+04
HPL_AI WRC01 95000 384 2 2 14.75 3.876e+04 5.77248 2 2.785e+04
HPL_AI WRC01 95000 512 2 2 14.93 3.828e+04 5.76942 2 2.761e+04`,
			expected: float64(32651.666666666668),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputReader := strings.NewReader(tt.data)
			got, err := result.ComputeAvgGflopsResult(result.NewReader(inputReader))
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}
