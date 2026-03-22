// Package opt provides a generic Optional type to represent values that may or may not exist.
// It is inspired by Rust's Option type and encourages explicit error/null handling.
//
// Usage:
//
//	val := opt.Some(42)
//	if val.IsSome() {
//	    fmt.Println(val.Unwrap())
//	}
package opt

// Opt represents an optional value of type T.
type Opt[T any] struct {
	value T
	valid bool
}

// Some creates an Opt containing the provided value.
func Some[T any](v T) Opt[T] {
	return Opt[T]{
		value: v,
		valid: true,
	}
}

// None creates an empty Opt of type T.
func None[T any]() Opt[T] {
	return Opt[T]{
		valid: false,
	}
}

// IsSome returns true if the Opt contains a value.
func (o Opt[T]) IsSome() bool {
	return o.valid
}

// IsNone returns true if the Opt is empty.
func (o Opt[T]) IsNone() bool {
	return !o.valid
}

// Unwrap returns the contained value. It panics if the Opt is empty.
func (o Opt[T]) Unwrap() T {
	if !o.valid {
		panic("called Unwrap on an empty Opt")
	}
	return o.value
}

// UnwrapOr returns the contained value if present, otherwise returns the provided default value.
func (o Opt[T]) UnwrapOr(def T) T {
	if !o.valid {
		return def
	}
	return o.value
}

// Map applies a function to the contained value if present, returning a new Opt.
// If the original Opt is empty, it returns an empty Opt of the new type.
func Map[T, U any](o Opt[T], fn func(T) U) Opt[U] {
	if o.valid {
		return Some(fn(o.value))
	}
	return None[U]()
}
