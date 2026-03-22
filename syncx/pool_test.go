package syncx

import (
	"testing"
)

func TestPool(t *testing.T) {
	pool := NewPool[int, int](3, func(job int) int {
		return job * 2
	})

	for i := 1; i <= 5; i++ {
		pool.Submit(i)
	}
	pool.Close()

	sum := 0
	count := 0
	for res := range pool.Results() {
		sum += res
		count++
	}

	if count != 5 {
		t.Errorf("Expected 5 results, got %d", count)
	}

	// 2 + 4 + 6 + 8 + 10 = 30
	if sum != 30 {
		t.Errorf("Expected sum 30, got %d", sum)
	}
}

func TestPoolShutdown(t *testing.T) {
	pool := NewPool[int, int](2, func(job int) int {
		return job
	})

	pool.Submit(1)
	pool.Shutdown()
	pool.Submit(2) // Shouldn't block indefinitely

	count := 0
	for range pool.Results() {
		count++
	}

	if count > 1 {
		t.Errorf("Shutdown didn't stop jobs fast enough, processed %d", count)
	}
}
