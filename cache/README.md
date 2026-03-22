# `cache`

The `cache` package provides memory-optimized, thread-safe caching primitives.

## LRU Cache
A generic, thread-safe Least Recently Used (LRU) Cache utilizing `sync.RWMutex` and a doubly-linked list for $O(1)$ operations.

### Usage
```go
// Create an LRU cache that holds up to 100 string->int pairs
c := cache.NewLRU[string, int](100)

c.Set("alice", 30)
c.Set("bob", 25)

if age, ok := c.Get("alice"); ok {
    fmt.Println("Alice's age:", age)
}

c.Remove("bob")
fmt.Println("Cache size:", c.Len())
```

## LFU Cache
A thread-safe Least Frequently Used (LFU) cache that evicts items accessed the fewest number of times. When making space, if multiple items share the minimum frequency, the least recently used among them is evicted. Operations are strictly $O(1)$.

### Usage
```go
// Create an LFU cache that holds up to 100 string->int pairs
c := cache.NewLFU[string, int](100)

c.Set("alice", 30)
c.Set("bob", 25)

// "alice" is accessed again, raising its frequency
c.Get("alice")

// Inserting a new key evicts "bob", because it has a lower frequency than "alice"
c.Set("charlie", 28)
```
