//go:build unit

package try_test

import (
	"context"
	"testing"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/pkg/utils/try"
	"github.com/stretchr/testify/require"
)

func TestDoWithContextTimeout(t *testing.T) {
	err := try.DoWithContextTimeout(context.Background(), 5, time.Millisecond, 1*time.Second, func(ctx context.Context, try int) error {
		time.Sleep(2 * time.Second)
		return nil
	})

	require.Equal(t, context.DeadlineExceeded, err)
}
