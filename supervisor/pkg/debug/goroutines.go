package debug

import (
	"context"
	"runtime"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

func WatchGoRoutines(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	for {
		count := runtime.NumGoroutine()
		logger.I.Debug("goroutines running", zap.Int("n", count))
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
