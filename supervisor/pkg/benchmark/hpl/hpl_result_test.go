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

package hpl_test

import (
	"strings"
	"testing"
	"time"

	_ "embed"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/hpl"
	"github.com/stretchr/testify/require"
)

//go:embed fixtures/3n2gpu16cpuucxrdma.log
var fixture3n2gpu16cpuucxrdma string

//go:embed fixtures/3n2gpu16cpuucxtcp.log
var fixture3n2gpu16cpuucxtcp string

//go:embed fixtures/1n2gpu16cpu.log
var fixture1n2gpu16cpu string

func TestFindMaxGflopsResult(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected *hpl.Result
	}{
		{
			name: "Valid test",
			data: `HPL_AI WRC01 95000 64 2 2 29.67 1.927e+04 5.71402 2 1.616e+04
HPL_AI WRC01 95000 128 2 2 15.67 3.647e+04 5.72738 2 2.671e+04
HPL_AI WRC01 95000 224 2 2 19.00 3.009e+04 5.74521 2 2.310e+04
HPL_AI WRC01 95000 256 2 2 17.30 3.304e+04 5.71711 2 2.483e+04
HPL_AI WRC01 95000 384 2 2 14.75 3.876e+04 5.77248 2 2.785e+04
HPL_AI WRC01 95000 512 2 2 14.93 3.828e+04 5.76942 2 2.761e+04`,
			expected: &hpl.Result{
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
			name: "Real test 1",
			data: fixture1n2gpu16cpu,
			expected: &hpl.Result{
				ProblemSize:       97000,
				NB:                1024,
				P:                 2,
				Q:                 1,
				Time:              time.Duration(6300000000),
				Gflops:            96510,
				Refine:            7.1469,
				Iterations:        2,
				GflopsWRefinement: 45230,
			},
		},
		{
			name: "Real test 2",
			data: fixture3n2gpu16cpuucxrdma,
			expected: &hpl.Result{
				ProblemSize:       168000,
				NB:                1024,
				P:                 3,
				Q:                 2,
				Time:              time.Duration(9340000000),
				Gflops:            338300,
				Refine:            15.12079,
				Iterations:        5,
				GflopsWRefinement: 129200,
			},
		},
		{
			name: "Real test 3",
			data: fixture3n2gpu16cpuucxtcp,
			expected: &hpl.Result{
				ProblemSize:       166000,
				NB:                1024,
				P:                 3,
				Q:                 2,
				Time:              time.Duration(22840000000),
				Gflops:            133500,
				Refine:            14.50631,
				Iterations:        5,
				GflopsWRefinement: 81660,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputReader := strings.NewReader(tt.data)
			got, err := hpl.FindMaxGflopsResult(hpl.NewReader(inputReader))
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
			got, err := hpl.ComputeAvgGflopsResult(hpl.NewReader(inputReader))
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}
