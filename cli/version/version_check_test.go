package version_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/deepsquare-io/grid/cli/version"
	"github.com/stretchr/testify/require"
)

func TestCheckLatest(t *testing.T) {
	latest, err := version.CheckLatest(context.Background())
	require.NoError(t, err)
	fmt.Println(latest)
}
