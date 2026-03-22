package poolx

import "sync"

// TypedPool is a purely type-safe wrapper around Go's native sync.Pool.
// It removes the boilerplate requirement to constantly cast `any` (interface{})
// strings or pointers every time memory is retrieved from the pool.
type TypedPool[T any] struct {
	pool sync.Pool
}

// NewTypedPool instantiates a dedicated memory pool for exactly one data type.
// The `newFunc` is called when the pool is empty to create a new object.
func NewTypedPool[T any](newFunc func() T) *TypedPool[T] {
	return &TypedPool[T]{
		pool: sync.Pool{
			New: func() any {
				return newFunc()
			},
		},
	}
}

// Get retrieves an object from the pool.
// It returns a value of type T.
func (p *TypedPool[T]) Get() T {
	return p.pool.Get().(T) // Guaranteed to be T since Put() exclusively takes T
}

// Put adds an object back to the pool for reuse.
func (p *TypedPool[T]) Put(x T) {
	p.pool.Put(x)
}
