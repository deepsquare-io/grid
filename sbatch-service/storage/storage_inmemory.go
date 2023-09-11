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
