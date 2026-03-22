// Package algo provides generic implementation of common algorithms.
package algo

import "cmp"

// BinarySearch searches for a target value in a sorted slice of items.
// It uses a selector function to extract the value from each item for comparison.
// It returns the index of the item if found, otherwise it returns -1.
func BinarySearch[T any, U cmp.Ordered](items []T, target U, selector func(T) U) int {
	low := 0
	high := len(items) - 1

	for low <= high {
		mid := low + (high-low)/2
		val := selector(items[mid])

		if val == target {
			return mid
		} else if val < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
