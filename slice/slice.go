// Package slice provides generic utility functions for manipulating slices.
package slice

// Map iterates over a slice and applies a function to each element,
// returning a new slice containing the results.
func Map[T, U any](items []T, iteratee func(T) U) []U {
	if items == nil {
		return nil
	}
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = iteratee(item)
	}
	return result
}
