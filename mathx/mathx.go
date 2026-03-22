// Package mathx provides generic math operations for numeric types,
// such as Sum, Average, and Clamp.
//
// Usage:
//
//	total := mathx.Sum([]int{1, 2, 3})
//	clamped := mathx.Clamp(150, 0, 100)
package mathx

import "cmp"

// Integer represents all signed and unsigned integer types.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float represents all floating-point types.
type Float interface {
	~float32 | ~float64
}

// Number represents all numeric types.
type Number interface {
	Integer | Float
}

// Sum calculates the sum of a slice of numbers.
func Sum[T Number](items []T) T {
	var sum T
	for _, v := range items {
		sum += v
	}
	return sum
}

// Average calculates the arithmetic mean of a slice of numbers.
// It returns 0 if the slice is empty.
func Average[T Number](items []T) float64 {
	if len(items) == 0 {
		return 0
	}
	var sum float64
	for _, v := range items {
		sum += float64(v)
	}
	return sum / float64(len(items))
}

// MinBy finds the minimum element in a slice based on a selector function.
// It returns a pointer to the element, or nil if the slice is empty.
func MinBy[T any, U cmp.Ordered](items []T, selector func(T) U) *T {
	if len(items) == 0 {
		return nil
	}
	minItem := items[0]
	minVal := selector(minItem)

	for i := 1; i < len(items); i++ {
		val := selector(items[i])
		if val < minVal {
			minVal = val
			minItem = items[i]
		}
	}
	return &minItem
}

// MaxBy finds the maximum element in a slice based on a selector function.
// It returns a pointer to the element, or nil if the slice is empty.
func MaxBy[T any, U cmp.Ordered](items []T, selector func(T) U) *T {
	if len(items) == 0 {
		return nil
	}
	maxItem := items[0]
	maxVal := selector(maxItem)

	for i := 1; i < len(items); i++ {
		val := selector(items[i])
		if val > maxVal {
			maxVal = val
			maxItem = items[i]
		}
	}
	return &maxItem
}

// Clamp restricts a value to be within a specific range [min, max].
func Clamp[T cmp.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
