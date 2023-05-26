package lock

import (
	"fmt"
	"sync"
)

type ResourceManager struct {
	locks map[string]*sync.Mutex
	mutex sync.Mutex
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		locks: make(map[string]*sync.Mutex),
	}
}

func (rm *ResourceManager) Lock(name string) {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	lock, ok := rm.locks[name]
	if !ok {
		lock = &sync.Mutex{}
		rm.locks[name] = lock
	}
	lock.Lock()
}

func (rm *ResourceManager) Unlock(name string) {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	lock, ok := rm.locks[name]
	if !ok {
		panic(fmt.Sprintf("Unlock called on non-existing lock: %s", name))
	}
	lock.Unlock()
}
