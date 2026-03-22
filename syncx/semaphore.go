package syncx

import "context"

// Semaphore provides a way to limit concurrent access to a resource.
// It is unweighted, meaning each Acquire takes exactly 1 slot.
type Semaphore chan struct{}

// NewSemaphore creates a Semaphore with a maximum of n concurrent slots.
func NewSemaphore(n int) Semaphore {
	if n <= 0 {
		panic("semaphore size must be greater than 0")
	}
	return make(chan struct{}, n)
}

// Acquire blocks until a slot is available or the context is canceled.
// It returns nil if a slot was successfully acquired, or ctx.Err() if canceled.
func (s Semaphore) Acquire(ctx context.Context) error {
	select {
	case s <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TryAcquire attempts to acquire a slot without blocking.
// Returns true if successfully acquired, false otherwise.
func (s Semaphore) TryAcquire() bool {
	select {
	case s <- struct{}{}:
		return true
	default:
		return false
	}
}

// Release frees up a slot in the semaphore.
// It will panic if the semaphore is released more times than it was acquired.
func (s Semaphore) Release() {
	select {
	case <-s:
	default:
		panic("semaphore release without acquire")
	}
}
