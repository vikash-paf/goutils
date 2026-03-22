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
