// Package result provides a functional approach to error handling.
package result

// Result encapsulates either a successful value of type T or an error.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a successful Result.
func Ok[T any](v T) Result[T] {
	return Result[T]{value: v}
}

// Err creates a failed Result.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// IsOk returns true if the result is successful.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the result contains an error.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Unwrap returns the value if successful, otherwise it panics.
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

// UnwrapOr returns the value if successful, otherwise returns the provided default.
func (r Result[T]) UnwrapOr(def T) T {
	if r.err != nil {
		return def
	}
	return r.value
}

// Error returns the error if present, or nil.
func (r Result[T]) Error() error {
	return r.err
}

// Value returns the value if present, or the zero value.
func (r Result[T]) Value() T {
	return r.value
}

// Map applies a function to the value if successful.
func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	if r.err != nil {
		return Err[U](r.err)
	}
	return Ok(fn(r.value))
}

// AndThen applies a function that returns a Result to the value if successful.
func AndThen[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	if r.err != nil {
		return Err[U](r.err)
	}
	return fn(r.value)
}

// Match executes onOk if the result is successful, otherwise executes onErr.
func (r Result[T]) Match(onOk func(T), onErr func(error)) {
	if r.err != nil {
		onErr(r.err)
		return
	}
	onOk(r.value)
}
