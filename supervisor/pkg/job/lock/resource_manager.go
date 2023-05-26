package lock

import (
	"fmt"
	"sync"
)

type ResourceManager struct {
	locks map[string]*sync.Mutex
	mutex sync.Mutex
}

func NewNamedLock() *ResourceManager {
	return &ResourceManager{
		locks: make(map[string]*sync.Mutex),
	}
}

func (nl *ResourceManager) Lock(name string) {
	nl.mutex.Lock()
	defer nl.mutex.Unlock()

	lock, ok := nl.locks[name]
	if !ok {
		lock = &sync.Mutex{}
		nl.locks[name] = lock
	}
	lock.Lock()
}

func (nl *ResourceManager) Unlock(name string) {
	nl.mutex.Lock()
	defer nl.mutex.Unlock()

	lock, ok := nl.locks[name]
	if !ok {
		panic(fmt.Sprintf("Unlock called on non-existing lock: %s", name))
	}
	lock.Unlock()
}
