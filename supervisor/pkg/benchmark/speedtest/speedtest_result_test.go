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

package speedtest_test

import (
	"testing"

	_ "embed"

	"github.com/deepsquare-io/grid/supervisor/pkg/benchmark/speedtest"
	"github.com/stretchr/testify/require"
)

//go:embed fixtures/result.json
var fixture []byte

func TestUnmarshal(t *testing.T) {
	r, err := speedtest.UnmarshalResult(fixture)
	require.NoError(t, err)
	require.EqualValues(t, 1135906059, r.Download.Bandwidth)

	_, err = r.Marshal()
	require.NoError(t, err)
}
