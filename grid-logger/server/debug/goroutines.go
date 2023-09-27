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

package debug

import (
	"context"
	"runtime"
	"time"

	"github.com/deepsquare-io/grid/grid-logger/logger"
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
