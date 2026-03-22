// Package iterx provides high-level functional utilities for Go 1.23+ iterators.
// It allows for transforming, filtering, and aggregating iterators with a clean API.
//
// Usage:
//
//	nums := []int{1, 2, 3}
//	seq := iterx.FromSlice(nums)
//	doubled := iterx.Map(seq, func(n int) int { return n * 2 })
//	slice := iterx.ToSlice(doubled)
package iterx

import "iter"

// FromSlice converts a slice into an iterator sequence.
func FromSlice[T any](slice []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range slice {
			if !yield(v) {
				return
			}
		}
	}
}

// ToSlice converts an iterator sequence into a slice.
func ToSlice[T any](seq iter.Seq[T]) []T {
	var slice []T
	for v := range seq {
		slice = append(slice, v)
	}
	return slice
}

// Map applies a function to each element of an iterator sequence,
// returning a new iterator sequence with the transformed elements.
func Map[T, U any](seq iter.Seq[T], fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

// Filter returns a new iterator sequence containing only the elements
// that pass the predicate function.
func Filter[T any](seq iter.Seq[T], fn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if fn(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Reduce reduces an iterator sequence to a single value by iteratively
// combining each element with an accumulator value.
func Reduce[T, R any](seq iter.Seq[T], fn func(R, T) R, initial R) R {
	result := initial
	for v := range seq {
		result = fn(result, v)
	}
	return result
}

// Take returns a new iterator sequence containing only the first n elements.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n <= 0 {
			return
		}
		count := 0
		for v := range seq {
			if !yield(v) {
				return
			}
			count++
			if count >= n {
				return
			}
		}
	}
}
