package syncx

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	sem := NewSemaphore(2)

	// Acquire 2 slots
	if err := sem.Acquire(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := sem.Acquire(context.Background()); err != nil {
		t.Fatal(err)
	}

	// Try acquiring 3rd should fail non-blocking
	if sem.TryAcquire() {
		t.Fatal("TryAcquire should fail when semaphore is full")
	}

	sem.Release()

	// 3rd should now succeed
	if !sem.TryAcquire() {
		t.Fatal("TryAcquire should succeed after a release")
	}
}

func TestSemaphore_ContextCancellation(t *testing.T) {
	sem := NewSemaphore(1)
	if err := sem.Acquire(context.Background()); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := sem.Acquire(ctx)
	if err != context.DeadlineExceeded {
		t.Fatalf("expected DeadlineExceeded, got %v", err)
	}
}

func TestSemaphore_Concurrency(t *testing.T) {
	sem := NewSemaphore(3)
	var wg sync.WaitGroup
	var active int32

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem.Acquire(context.Background())
			current := atomic.AddInt32(&active, 1)
			if current > 3 {
				t.Errorf("too many active routines: %d", current)
			}
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt32(&active, -1)
			sem.Release()
		}()
	}
	wg.Wait()
}

func ExampleSemaphore() {
	sem := NewSemaphore(2) // Max 2 concurrent workers

	for i := 0; i < 3; i++ {
		sem.Acquire(context.Background())
		go func(id int) {
			defer sem.Release()
			// Simulate work
			// Do heavy lifting here
		}(i)
	}
	fmt.Println("Done")
	// Output: Done
}
