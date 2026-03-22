package cache

import (
	"container/list"
	"sync"
)

// LFU is a generic, thread-safe Least Frequently Used (LFU) cache.
// When the cache is full, the least frequently used item is evicted.
// If multiple items share the same minimum frequency, the least recently used among them is evicted.
type LFU[K comparable, V any] struct {
	capacity int
	mu       sync.RWMutex
	cache    map[K]*list.Element
	freqs    map[int]*list.List
	minFreq  int
}

type lfuItem[K comparable, V any] struct {
	key   K
	value V
	freq  int
}

// NewLFU creates a new LFU cache with the specified capacity.
func NewLFU[K comparable, V any](capacity int) *LFU[K, V] {
	if capacity <= 0 {
		panic("LFU cache capacity must be > 0")
	}
	return &LFU[K, V]{
		capacity: capacity,
		cache:    make(map[K]*list.Element),
		freqs:    make(map[int]*list.List),
	}
}

func (c *LFU[K, V]) increaseFreq(element *list.Element) {
	item := element.Value.(*lfuItem[K, V])

	// Remove from current freq list
	oldList := c.freqs[item.freq]
	oldList.Remove(element)

	// Check if minFreq needs to be updated
	if item.freq == c.minFreq && oldList.Len() == 0 {
		c.minFreq++
	}

	item.freq++

	// Add to new freq list
	newList, ok := c.freqs[item.freq]
	if !ok {
		newList = list.New()
		c.freqs[item.freq] = newList
	}

	newElement := newList.PushFront(item)
	// Update cache pointer to the new element
	c.cache[item.key] = newElement
}

// Set adds a value to the cache or updates an existing one.
func (c *LFU[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.cache[key]; ok {
		element.Value.(*lfuItem[K, V]).value = value
		c.increaseFreq(element)
		return
	}

	if len(c.cache) >= c.capacity {
		minList := c.freqs[c.minFreq]
		back := minList.Back()
		if back != nil {
			minList.Remove(back)
			delete(c.cache, back.Value.(*lfuItem[K, V]).key)
		}
	}

	item := &lfuItem[K, V]{key: key, value: value, freq: 1}

	l, ok := c.freqs[1]
	if !ok {
		l = list.New()
		c.freqs[1] = l
	}
	element := l.PushFront(item)

	c.cache[key] = element
	c.minFreq = 1
}

// Get looks up a key's value from the cache.
func (c *LFU[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.cache[key]; ok {
		c.increaseFreq(element)
		return element.Value.(*lfuItem[K, V]).value, true
	}

	var zero V
	return zero, false
}

// Contains checks if the cache has a key without updating its frequency.
func (c *LFU[K, V]) Contains(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.cache[key]
	return ok
}

// Remove deletes an item from the cache.
func (c *LFU[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.cache[key]; ok {
		item := element.Value.(*lfuItem[K, V])
		l := c.freqs[item.freq]
		l.Remove(element)
		delete(c.cache, key)
	}
}

// Clear removes all items from the cache.
func (c *LFU[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[K]*list.Element)
	c.freqs = make(map[int]*list.List)
	c.minFreq = 0
}

// Len returns the number of items currently in the cache.
func (c *LFU[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.cache)
}
