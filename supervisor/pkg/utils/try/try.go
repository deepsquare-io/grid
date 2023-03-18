package try

import (
	"context"
	"time"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
)

func Do(
	fn func() error,
	tries int,
	delay time.Duration,
) (err error) {
	for try := 0; try < tries; try++ {
		err = fn()
		if err == nil {
			break
		}
		logger.I.Warn("try failed", zap.Error(err), zap.Int("try", try))
		time.Sleep(delay)
	}
	if err != nil {
		logger.I.Warn("failed all tries", zap.Error(err))
	}
	return err
}

func DoWithContextTimeout(
	parent context.Context,
	fn func() error,
	tries int,
	delay time.Duration,
	timeout time.Duration,
) (err error) {
	for try := 0; try < tries; try++ {
		ctx, cancel := context.WithTimeout(parent, timeout)
		defer cancel()
		errChan := make(chan error)

		go func() {
			errChan <- fn()
		}()

		select {
		case err = <-errChan:
			if err != nil {
				logger.I.Warn("try failed", zap.Error(err), zap.Int("try", try))
			}
			if err == nil {
				return nil
			}
		case <-ctx.Done():
			err = ctx.Err()
			logger.I.Warn("try failed", zap.Error(err), zap.Int("try", try))
		}
		time.Sleep(delay)
	}
	if err != nil {
		logger.I.Warn("failed all tries", zap.Error(err))
	}
	return err
}

func DoWithResult[T interface{}](
	fn func() (T, error),
	tries int,
	delay time.Duration,
) (result T, err error) {
	for try := 0; try < tries; try++ {
		result, err = fn()
		if err == nil {
			break
		}
		logger.I.Warn("try failed", zap.Error(err), zap.Int("try", try))
		time.Sleep(delay)
	}
	if err != nil {
		logger.I.Warn("failed all tries", zap.Error(err))
	}
	return result, err
}
