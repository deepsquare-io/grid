// Copyright (C) 2024 DeepSquare Association
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

package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ Storage = (*RedisStorage)(nil)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(c *redis.Client) *RedisStorage {
	if c == nil {
		panic("redis client is nil")
	}
	return &RedisStorage{
		client: c,
	}
}

func (s *RedisStorage) Set(
	ctx context.Context,
	key string,
	value string,
	expiration time.Duration,
) error {
	_, err := s.client.Set(ctx, key, value, expiration).Result()
	return err
}

func (s *RedisStorage) Get(ctx context.Context, key string) (string, error) {
	resp, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrNotFound
		}
		return "", err
	}
	return resp, nil
}
