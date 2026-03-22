package cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := NewLRU[int, string](2)

	cache.Set(1, "one")
	cache.Set(2, "two")

	if val, ok := cache.Get(1); !ok || val != "one" {
		t.Errorf("Get(1) = %v, want one", val)
	}

	// 1 was accessed, so 2 should be the LRU. Adding 3 should evict 2.
	cache.Set(3, "three")

	if _, ok := cache.Get(2); ok {
		t.Error("2 should have been evicted")
	}

	if cache.Len() != 2 {
		t.Errorf("Len() = %d, want 2", cache.Len())
	}

	cache.Remove(3)
	if cache.Contains(3) {
		t.Error("3 should have been removed")
	}

	cache.Clear()
	if cache.Len() != 0 {
		t.Error("Cache should be empty after clear")
	}
}

func TestLRUConcurrency(t *testing.T) {
	cache := NewLRU[int, int](100)
	var wg sync.WaitGroup

	// Concurrently Set and Get
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			cache.Set(key%10, key)
			cache.Get(key % 10)
		}(i)
	}

	wg.Wait()
	if cache.Len() > 10 {
		t.Errorf("Cache exceeded capacity")
	}
}

func ExampleNewLRU() {
	c := NewLRU[string, int](2)
	c.Set("apple", 1)
	c.Set("banana", 2)
	c.Set("cherry", 3) // Evicts "apple"

	if val, ok := c.Get("apple"); ok {
		fmt.Println("Found apple:", val)
	} else {
		fmt.Println("Apple evicted")
	}
	// Output: Apple evicted
}
