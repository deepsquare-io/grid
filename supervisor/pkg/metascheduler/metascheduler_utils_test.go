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

package metascheduler_test

import (
	"testing"

	metaschedulerabi "github.com/deepsquare-io/grid/supervisor/generated/abi/metascheduler"
	"github.com/deepsquare-io/grid/supervisor/pkg/metascheduler"
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
