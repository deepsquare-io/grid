package storage

import (
	"context"
	"errors"
	"time"
)

var ErrNotFound = errors.New("no entry exists under this name")

type Storage interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
