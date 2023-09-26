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
