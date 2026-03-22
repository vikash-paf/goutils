// Package ptr provides generic utility functions for pointer manipulation.
package ptr

// Of returns a pointer to the given value.
func Of[T any](v T) *T {
	return &v
}

// Val safely dereferences a pointer. It returns the zero value of the type
// if the pointer is nil.
func Val[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// ValOrDefault safely dereferences a pointer. It returns the provided default
// value if the pointer is nil.
func ValOrDefault[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// Equal safely compares the values referenced by two pointers.
// It returns true if both are nil, or if both point to equal values.
func Equal[T comparable](p1, p2 *T) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}
	return *p1 == *p2
}
