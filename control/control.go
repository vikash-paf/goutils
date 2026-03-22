// Package control provides utility functions for flow control and error handling,
// such as a ternary operator (If) and the Must pattern.
//
// Usage:
//
//	val := control.If(condition, "true", "false")
//	db := control.Must(sql.Open("postgres", connStr))
package control

// If is a generic ternary operator. It returns trueVal if the condition is true,
// otherwise it returns falseVal.
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Must panics if the error is not nil, otherwise it returns the value.
// It is useful for standardizing the "panic on error" initialization pattern.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Coalesce returns the first non-zero value from the provided arguments.
// If all values are the zero value of type T, it returns the zero value.
func Coalesce[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

// Try executes a function that returns a value and an error.
// If the function returns an error, Try returns the fallback value instead.
// If the function succeeds, it returns the function's result.
func Try[T any](fallback T, fn func() (T, error)) T {
	val, err := fn()
	if err != nil {
		return fallback
	}
	return val
}
