package slice

import "math/rand"

// Find searches for the first element in the slice that satisfies the predicate.
// It returns a pointer to the element if found, or nil if not found.
func Find[T any](items []T, predicate func(T) bool) *T {
	for _, item := range items {
		if predicate(item) {
			val := item
			return &val
		}
	}
	return nil
}

// FindIndex searches for the first element in the slice that satisfies the predicate
// and returns its index. It returns -1 if no element is found.
func FindIndex[T any](items []T, predicate func(T) bool) int {
	for i, item := range items {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// Some returns true if at least one element in the slice satisfies the predicate.
func Some[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Every returns true if all elements in the slice satisfy the predicate.
// It returns true for empty slices (vacuous truth).
func Every[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Reverse reverses the elements of the slice in place.
func Reverse[T any](items []T) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

// Shuffle randomizes the order of elements in the slice in place.
func Shuffle[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}

// Partition splits the slice into two slices: one with elements that satisfy the predicate,
// and one with elements that do not.
func Partition[T any](items []T, predicate func(T) bool) (passed []T, failed []T) {
	for _, item := range items {
		if predicate(item) {
			passed = append(passed, item)
		} else {
			failed = append(failed, item)
		}
	}
	if passed == nil {
		passed = []T{}
	}
	if failed == nil {
		failed = []T{}
	}
	return passed, failed
}

// DiffState compares two slices and returns the items that were added and removed.
func DiffState[T comparable](old, new []T) (added []T, removed []T) {
	oldSet := make(map[T]struct{}, len(old))
	newSet := make(map[T]struct{}, len(new))
	
	for _, v := range old {
		oldSet[v] = struct{}{}
	}
	for _, v := range new {
		newSet[v] = struct{}{}
	}

	for _, v := range old {
		if _, ok := newSet[v]; !ok {
			removed = append(removed, v)
		}
	}
	for _, v := range new {
		if _, ok := oldSet[v]; !ok {
			added = append(added, v)
		}
	}
	
	if added == nil {
		added = []T{}
	}
	if removed == nil {
		removed = []T{}
	}
	return added, removed
}

// CountBy returns the number of elements in the slice that satisfy the predicate.
func CountBy[T any](items []T, predicate func(T) bool) int {
	count := 0
	for _, item := range items {
		if predicate(item) {
			count++
		}
	}
	return count
}
