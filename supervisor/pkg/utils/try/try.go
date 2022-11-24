package try

import (
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
	return result, err
}
