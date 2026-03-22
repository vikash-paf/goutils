package syncx

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	var count int32
	var wg sync.WaitGroup

	fn := func() {
		atomic.AddInt32(&count, 1)
		wg.Done()
	}

	debounced := Debounce(50*time.Millisecond, fn)

	wg.Add(1)
	// Call rapidly, should only execute once after 50ms
	for range 5 {
		debounced()
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Debounce executed %d times, want 1", count)
	}
}

func TestThrottle(t *testing.T) {
	var count int32

	fn := func() {
		atomic.AddInt32(&count, 1)
	}

	throttled := Throttle(50*time.Millisecond, fn)

	for i := 0; i < 10; i++ {
		throttled()
		time.Sleep(10 * time.Millisecond)
	}

	// Should execute roughly 2 times (10 * 10ms = 100ms / 50ms = 2)
	c := atomic.LoadInt32(&count)
	if c < 1 || c > 3 {
		t.Errorf("Throttle executed %d times, expected ~2", c)
	}
}

func TestBatcher(t *testing.T) {
	var processedBatches int32
	var wg sync.WaitGroup

	b := NewBatcher(3, 50*time.Millisecond, func(items []int) {
		atomic.AddInt32(&processedBatches, 1)
		wg.Done()
	})

	wg.Add(2) // Expecting 2 batches: one full (size 3), one partial triggered by timeout/close (size 2)

	// Add 5 items
	for i := 0; i < 5; i++ {
		b.Add(i)
	}

	b.Close() // Force flush of remaining 2 items

	wg.Wait()

	if atomic.LoadInt32(&processedBatches) != 2 {
		t.Errorf("Batcher processed %d batches, want 2", processedBatches)
	}
}

func ExampleBatcher() {
	var wg sync.WaitGroup
	wg.Add(1)

	b := NewBatcher(2, 100*time.Millisecond, func(items []string) {
		fmt.Printf("Batch processed: %v\n", items)
		wg.Done()
	})

	b.Add("Alice")
	b.Add("Bob") // Batch size 2 reached, triggers flush immediately

	wg.Wait()
	// Output: Batch processed: [Alice Bob]
}
