package auth

import (
	"errors"
	"sync"
)

var (
	ErrUserNotExists = errors.New("user does not exist")
	ErrUserExists    = errors.New("user already exists")
)

type User struct {
	Address string
	Nonce   []byte
}

type MemStorage struct {
	lock  sync.RWMutex
	users map[string]User
}

func (m *MemStorage) CreateIfNotExists(u User) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, exists := m.users[u.Address]; exists {
		return ErrUserExists
	}
	m.users[u.Address] = u
	return nil
}

func (m *MemStorage) Get(address string) (User, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	u, exists := m.users[address]
	if !exists {
		return u, ErrUserNotExists
	}
	return u, nil
}

func (m *MemStorage) Update(user User) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.users[user.Address] = user
	return nil
}

func NewMemStorage() *MemStorage {
	ans := MemStorage{
		users: make(map[string]User),
	}
	return &ans
}
