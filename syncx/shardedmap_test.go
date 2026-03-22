package syncx

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestShardedMap(t *testing.T) {
	sm := NewShardedMap[int](4)

	sm.Set("one", 1)
	sm.Set("two", 2)
	sm.Set("three", 3)

	if sm.Len() != 3 {
		t.Errorf("expected 3 items, got %d", sm.Len())
	}

	val, ok := sm.Get("two")
	if !ok || val != 2 {
		t.Errorf("expected 2, got %v", val)
	}

	sm.Delete("one")
	if sm.Len() != 2 {
		t.Errorf("expected 2 items after deletion, got %d", sm.Len())
	}

	if _, ok := sm.Get("one"); ok {
		t.Error("expected 'one' to be deleted")
	}

	if _, ok := sm.Get("nonexistent"); ok {
		t.Error("expected nonexistent key to return false")
	}
}

func TestShardedMap_Concurrency(t *testing.T) {
	sm := NewShardedMap[string](16)
	var wg sync.WaitGroup

	numWriters := 100
	writesPerWorker := 100

	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()
			for j := 0; j < writesPerWorker; j++ {
				key := strconv.Itoa(writerID) + "-" + strconv.Itoa(j)
				sm.Set(key, "data")
			}
		}(i)
	}

	wg.Wait()

	expectedLength := numWriters * writesPerWorker
	if sm.Len() != expectedLength {
		t.Errorf("expected %d elements, got %d", expectedLength, sm.Len())
	}
}

func ExampleShardedMap() {
	// Initialize a map partitioned into 32 locks natively
	cache := NewShardedMap[int](32)

	cache.Set("alice", 100)
	cache.Set("bob", 200)

	val, ok := cache.Get("alice")
	if ok {
		fmt.Printf("Alice has %d points.\n", val)
	}

	fmt.Println("Total items:", cache.Len())

	// Output:
	// Alice has 100 points.
	// Total items: 2
}
