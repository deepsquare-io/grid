package lock

import (
	"fmt"
	"sync"

	"github.com/deepsquare-io/the-grid/supervisor/logger"
	"go.uber.org/zap"
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
	logger.I.Debug("locking...", zap.String("name", name))
	rm.mutex.Lock()
	lock, ok := rm.locks[name]
	if !ok {
		lock = &sync.Mutex{}
		rm.locks[name] = lock
	}
	rm.mutex.Unlock()
	lock.Lock()
	logger.I.Debug("locked", zap.String("name", name))
}

func (rm *ResourceManager) Unlock(name string) {
	logger.I.Debug("unlocking...", zap.String("name", name))
	rm.mutex.Lock()
	lock, ok := rm.locks[name]
	if !ok {
		panic(fmt.Sprintf("Unlock called on non-existing lock: %s", name))
	}
	rm.mutex.Unlock()
	lock.Unlock()
	logger.I.Debug("unlocked", zap.String("name", name))
}
