package syncx

import (
	"context"
	"sync"
)

// Promise represents a concurrent operation that will eventually produce a value or an error.
// It allows for awaiting a result without manually managing channels.
type Promise[T any] struct {
	val  T
	err  error
	done chan struct{}
	once sync.Once
}

// NewPromise creates a new Promise.
func NewPromise[T any]() *Promise[T] {
	return &Promise[T]{
		done: make(chan struct{}),
	}
}

// Resolve fulfills the promise with a value.
func (p *Promise[T]) Resolve(val T) {
	p.once.Do(func() {
		p.val = val
		close(p.done)
	})
}

// Reject fulfills the promise with an error.
func (p *Promise[T]) Reject(err error) {
	p.once.Do(func() {
		p.err = err
		close(p.done)
	})
}

// Await blocks until the promise is fulfilled and returns the value and error.
func (p *Promise[T]) Await(ctx context.Context) (T, error) {
	select {
	case <-p.done:
		return p.val, p.err
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	}
}
