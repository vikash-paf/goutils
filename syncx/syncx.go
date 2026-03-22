// Package syncx provides generic concurrency primitives like Debounce, Throttle, and Batcher.
package syncx

import (
	"sync"
	"time"
)

// Debounce creates a debounced function that delays invoking the provided function
// until after the specified duration has elapsed since the last time the debounced function was invoked.
func Debounce(interval time.Duration, fn func()) func() {
	var mu sync.Mutex
	var timer *time.Timer

	return func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(interval, fn)
	}
}

// Throttle creates a throttled function that only invokes the provided function
// at most once per every specified duration.
func Throttle(interval time.Duration, fn func()) func() {
	var mu sync.Mutex
	var lastExec time.Time

	return func() {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if now.Sub(lastExec) >= interval {
			lastExec = now
			// Execute in a goroutine to avoid blocking the caller
			go fn()
		}
	}
}

// Batcher collects items and processes them in batches.
// A batch is processed either when the batch size is reached or when the timeout expires.
type Batcher[T any] struct {
	batchSize int
	timeout   time.Duration
	processor func([]T)

	items  []T
	mu     sync.Mutex
	timer  *time.Timer
	closed bool
}

// NewBatcher creates a new Batcher.
func NewBatcher[T any](size int, timeout time.Duration, processor func([]T)) *Batcher[T] {
	if size <= 0 {
		panic("Batcher size must be > 0")
	}
	return &Batcher[T]{
		batchSize: size,
		timeout:   timeout,
		processor: processor,
	}
}

// Add appends an item to the current batch.
func (b *Batcher[T]) Add(item T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	b.items = append(b.items, item)

	if len(b.items) >= b.batchSize {
		b.flush()
	} else if len(b.items) == 1 {
		// First item in a new batch, start the timer
		if b.timer != nil {
			b.timer.Stop()
		}
		b.timer = time.AfterFunc(b.timeout, func() {
			b.mu.Lock()
			defer b.mu.Unlock()
			b.flush()
		})
	}
}

// flush processes the current batch and resets the internal state.
// It assumes the caller holds the lock.
func (b *Batcher[T]) flush() {
	if len(b.items) > 0 {
		batch := b.items
		b.items = nil // Reset for next batch

		if b.timer != nil {
			b.timer.Stop()
		}

		// Process asynchronously so we don't hold the lock or block the timer
		go b.processor(batch)
	}
}

// Close flushes any remaining items and prevents further additions.
func (b *Batcher[T]) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.closed = true
	b.flush()
}
