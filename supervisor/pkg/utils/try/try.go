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

package try

import (
	"context"
	"time"

	"github.com/deepsquare-io/grid/supervisor/logger"
	"go.uber.org/zap"
)

func Do(
	tries int,
	delay time.Duration,
	fn func(try int) error,
) (err error) {
	if tries <= 0 {
		logger.I.Panic("tries is 0 or negative", zap.Int("tries", tries))
	}
	for try := 0; try < tries; try++ {
		err = fn(try)
		if err == nil {
			return nil
		}
		logger.I.Warn("try failed", zap.Error(err), zap.Int("try", try), zap.Int("maxTries", tries))
		time.Sleep(delay)
	}
	logger.I.Warn("failed all tries", zap.Error(err))
	return err
}

func DoWithContextTimeout(
	parent context.Context,
	tries int,
	delay time.Duration,
	timeout time.Duration,
	fn func(ctx context.Context, try int) error,
) (err error) {
	if tries <= 0 {
		logger.I.Panic("tries is 0 or negative", zap.Int("tries", tries))
	}

	for try := 0; try < tries; try++ {
		err = func() error {
			ctx, cancel := context.WithTimeout(parent, timeout)
			defer cancel()

			errChan := make(chan error, 1)
			go func(try int) {
				defer close(errChan)
				errChan <- fn(ctx, try)
			}(try)

			select {
			case err = <-errChan:
				if err != nil {
					logger.I.Warn(
						"try failed",
						zap.Error(err),
						zap.Int("try", try),
						zap.Int("maxTries", tries),
					)
				}
				if err == nil {
					return nil
				}
			case <-ctx.Done():
				err = ctx.Err()
				logger.I.Warn(
					"try failed",
					zap.Error(err),
					zap.Int("try", try),
					zap.Int("maxTries", tries),
				)
			}
			return err
		}()

		if err == nil {
			return nil
		}

		time.Sleep(delay)
	}
	logger.I.Warn("failed all tries", zap.Error(err))
	return err
}

func DoWithResult[T interface{}](
	tries int,
	delay time.Duration,
	fn func(try int) (T, error),
) (result T, err error) {
	if tries <= 0 {
		logger.I.Panic("tries is 0 or negative", zap.Int("tries", tries))
	}
	for try := 0; try < tries; try++ {
		result, err = fn(try)
		if err == nil {
			return result, nil
		}
		logger.I.Warn("try failed", zap.Error(err), zap.Int("try", try))
		time.Sleep(delay)
	}
	logger.I.Warn("failed all tries", zap.Error(err))
	return result, err
}
