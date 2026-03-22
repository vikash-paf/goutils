package cache

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	c := NewLFU[int, int](2)

	c.Set(1, 1)
	c.Set(2, 2)

	if val, ok := c.Get(1); !ok || val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	c.Set(3, 3) // Evicts 2 because 1 has freq 2, and 2 has freq 1

	if _, ok := c.Get(2); ok {
		t.Error("expected 2 to be evicted")
	}

	if val, ok := c.Get(3); !ok || val != 3 {
		t.Errorf("expected 3, got %v", val)
	}

	c.Set(4, 4) // Evicts 3 because it has freq 1 (same as 4 initially, but 3 was just added so it's most recent... wait, 1 has freq 2. 3 has 1. We evict 3).

	if _, ok := c.Get(1); !ok {
		t.Error("expected 1 to remain")
	}

	if _, ok := c.Get(3); ok {
		t.Error("expected 3 to be evicted")
	}

	if val, ok := c.Get(4); !ok || val != 4 {
		t.Errorf("expected 4, got %v", val)
	}
}

func ExampleLFU() {
	cache := NewLFU[string, string](2)
	cache.Set("a", "alpha")
	cache.Set("b", "beta")

	// Increase 'a' frequency
	cache.Get("a")

	// 'b' is evicted because it was accessed less frequently
	cache.Set("c", "gamma")

	if _, ok := cache.Get("b"); !ok {
		fmt.Println("b was evicted")
	}

	// Output:
	// b was evicted
}
