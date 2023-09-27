// Copyright (C) 2023 DeepSquare Association
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

package lock

import (
	"fmt"
	"sync"

	"github.com/deepsquare-io/grid/supervisor/logger"
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
