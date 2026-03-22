package syncx

import (
	"sync"
)

// KeyMutex provides a set of mutexes identified by a comparable key.
// It allows for granular locking of resources.
type KeyMutex[K comparable] struct {
	locks map[K]*sync.Mutex
	mu    sync.Mutex
}

// NewKeyMutex creates a new KeyMutex.
func NewKeyMutex[K comparable]() *KeyMutex[K] {
	return &KeyMutex[K]{
		locks: make(map[K]*sync.Mutex),
	}
}

// Lock acquires the mutex associated with the given key.
func (km *KeyMutex[K]) Lock(key K) {
	km.mu.Lock()
	lock, ok := km.locks[key]
	if !ok {
		lock = &sync.Mutex{}
		km.locks[key] = lock
	}
	km.mu.Unlock()

	lock.Lock()
}

// Unlock releases the mutex associated with the given key.
func (km *KeyMutex[K]) Unlock(key K) {
	km.mu.Lock()
	lock, ok := km.locks[key]
	km.mu.Unlock()

	if !ok {
		panic("unlock of uninitialized key mutex")
	}
	lock.Unlock()
}

// Note: To prevent memory leaks, a production KeyMutex would need a reference counting
// mechanism to delete mutexes that are no longer in use. For simplicity in this utility
// and general small-scale key sets, this implementation works well.
