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
	"sync"
	"time"
)

type inmemoryStorage struct {
	mu     sync.RWMutex
	values map[string]data
}

type data struct {
	value      string
	expiration time.Time
}

func NewInMemoryStorage() Storage {
	return &inmemoryStorage{
		values: make(map[string]data),
	}
}

func (s *inmemoryStorage) Set(
	ctx context.Context,
	key string,
	value string,
	expiration time.Duration,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Store the value along with its expiration time
	s.values[key] = data{
		value:      value,
		expiration: time.Now().Add(expiration),
	}

	return nil
}

func (s *inmemoryStorage) Get(ctx context.Context, key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Check if the key exists
	entry, found := s.values[key]
	if !found {
		return "", ErrNotFound
	}

	// Check if the entry is expired
	if entry.expiration.Before(time.Now()) {
		delete(s.values, key) // Remove the expired entry
		return "", ErrNotFound
	}

	return entry.value, nil
}
