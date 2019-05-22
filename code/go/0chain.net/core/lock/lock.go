package lock

import "sync"

var lockPool = make(map[string]*sync.Mutex, 0)
var lockMutex = &sync.Mutex{}

func GetMutex(tablename string, key string) *sync.Mutex {
	lockKey := tablename + ":" + key
	lockMutex.Lock()
	defer lockMutex.Unlock()
	if eLock, ok := lockPool[lockKey]; ok {
		return eLock
	}
	lockPool[lockKey] = &sync.Mutex{}
	return lockPool[lockKey]
}
