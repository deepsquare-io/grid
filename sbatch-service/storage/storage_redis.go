package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	*redis.Client
}

func NewRedisStorage(c *redis.Client) Storage {
	if c == nil {
		panic("redis client is nil")
	}
	return &redisStorage{
		Client: c,
	}
}

func (s *redisStorage) Set(
	ctx context.Context,
	key string,
	value string,
	expiration time.Duration,
) error {
	_, err := s.Client.Set(ctx, key, value, expiration).Result()
	return err
}

func (s *redisStorage) Get(ctx context.Context, key string) (string, error) {
	resp, err := s.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrNotFound
		}
		return "", err
	}
	return resp, nil
}
