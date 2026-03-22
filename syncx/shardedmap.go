package syncx

import (
	"hash/maphash"
	"sync"
)

// ShardedMap is a high-performance concurrent map utilizing bucketing/sharding.
// By partitioning the map into shards, lock contention is massively reduced compared
// to a single sync.RWMutex lock securing the entire map structure.
// Specifically optimized for string keys due to rapid underlying hashing.
type ShardedMap[V any] struct {
	shards    []*shard[V]
	numShards uint64
	seed      maphash.Seed
}

// shard holds the intrinsic Go map and the mutex securing it.
type shard[V any] struct {
	sync.RWMutex
	m map[string]V
}

// NewShardedMap initializes a concurrent map broken into 'numShards' mutual exclusion buckets.
// If numShards <= 0, it defaults to an optimal 32.
func NewShardedMap[V any](numShards int) *ShardedMap[V] {
	if numShards <= 0 {
		numShards = 32
	}
	sm := &ShardedMap[V]{
		shards:    make([]*shard[V], numShards),
		numShards: uint64(numShards),
		seed:      maphash.MakeSeed(),
	}
	for i := range sm.shards {
		sm.shards[i] = &shard[V]{m: make(map[string]V)}
	}
	return sm
}

// getShard returns the shard responsible for the given key.
func (m *ShardedMap[V]) getShard(key string) *shard[V] {
	hashed := maphash.String(m.seed, key)
	return m.shards[hashed%m.numShards]
}

// Set adds or updates a value in the map.
func (m *ShardedMap[V]) Set(key string, val V) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.m[key] = val
}

// Get retrieves a value from the map. Returns false if the key doesn't exist.
func (m *ShardedMap[V]) Get(key string) (V, bool) {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	val, ok := shard.m[key]
	return val, ok
}

// Delete removes a key from the map.
func (m *ShardedMap[V]) Delete(key string) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	delete(shard.m, key)
}

// Len returns the total number of elements in the map.
// It iterates over all shards and locks them individually.
func (m *ShardedMap[V]) Len() int {
	count := 0
	for _, shard := range m.shards {
		shard.RLock()
		count += len(shard.m)
		shard.RUnlock()
	}
	return count
}
