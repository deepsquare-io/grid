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

package ior_test

import (
	"strings"
	"testing"

	_ "embed"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/ior"
	"github.com/stretchr/testify/require"
)

//go:embed fixtures/result.csv
var fixture string

func TestComputeAvgReadWrite(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		expectedR *ior.Result
		expectedW *ior.Result
	}{
		{
			name: "Real test",
			data: fixture,
			expectedR: &ior.Result{
				Access:        "read",
				Bandwidth:     111.97344000000001,
				IOPS:          111.97398000000001,
				Latency:       0.41022,
				BlockSize:     1.048576e+06,
				TransferSize:  1024,
				OpenDuration:  0.049980000000000004,
				WrRdDuration:  438.95918,
				CloseDuration: 238.14438,
				TotalDuration: 438.96126000000004,
				Tasks:         0x30,
				Iteration:     5,
			},
			expectedW: &ior.Result{
				Access:        "write",
				Bandwidth:     50.12118,
				IOPS:          50.12128,
				Latency:       0.8774200000000001,
				BlockSize:     1.048576e+06,
				TransferSize:  1024,
				OpenDuration:  1.05266,
				WrRdDuration:  980.74072,
				CloseDuration: 263.16834,
				TotalDuration: 980.7422399999999,
				Tasks:         0x30,
				Iteration:     5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputReader := strings.NewReader(tt.data)
			r, w, err := ior.ComputeAvgReadWrite(ior.NewReader(inputReader))
			require.NoError(t, err)
			require.Equal(t, tt.expectedR, r)
			require.Equal(t, tt.expectedW, w)
		})
	}
}
