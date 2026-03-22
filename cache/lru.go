// Package cache provides generic caching primitives.
package cache

import (
	"container/list"
	"sync"
)

// LRU is a generic, thread-safe Least Recently Used (LRU) cache.
type LRU[K comparable, V any] struct {
	capacity int
	mu       sync.RWMutex
	ll       *list.List
	cache    map[K]*list.Element
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewLRU creates a new LRU cache with the specified capacity.
func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	if capacity <= 0 {
		panic("LRU cache capacity must be > 0")
	}
	return &LRU[K, V]{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[K]*list.Element),
	}
}

// Set adds a value to the cache or updates an existing one.
func (c *LRU[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Update existing
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		element.Value.(*entry[K, V]).value = value
		return
	}

	// Evict if full
	if c.ll.Len() >= c.capacity {
		back := c.ll.Back()
		if back != nil {
			c.ll.Remove(back)
			delete(c.cache, back.Value.(*entry[K, V]).key)
		}
	}

	// Add new
	ele := c.ll.PushFront(&entry[K, V]{key, value})
	c.cache[key] = ele
}

// Get looks up a key's value from the cache.
func (c *LRU[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry[K, V]).value, true
	}

	var zero V
	return zero, false
}

// Contains checks if the cache has a key without updating its recency.
func (c *LRU[K, V]) Contains(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.cache[key]
	return ok
}

// Remove deletes an item from the cache.
func (c *LRU[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.ll.Remove(ele)
		delete(c.cache, key)
	}
}

// Clear removes all items from the cache.
func (c *LRU[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ll.Init()
	c.cache = make(map[K]*list.Element)
}

// Len returns the number of items currently in the cache.
func (c *LRU[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.ll.Len()
}
