// Package set provides a generic, native-feeling Set data structure.
package set

// Set represents a collection of unique elements.
type Set[T comparable] map[T]struct{}

// New creates a new Set with the given items.
func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// FromSlice creates a new Set from a slice of items.
func FromSlice[T comparable](items []T) Set[T] {
	return New(items...)
}

// Add inserts an element into the set.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Remove deletes an element from the set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Contains returns true if the set contains the given element.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Values returns all elements in the set as a slice.
// The order of the returned elements is not guaranteed.
func (s Set[T]) Values() []T {
	result := make([]T, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}

// Union returns a new Set containing all unique elements from both sets.
func Union[T comparable](s1, s2 Set[T]) Set[T] {
	result := make(Set[T], len(s1)+len(s2))
	for v := range s1 {
		result.Add(v)
	}
	for v := range s2 {
		result.Add(v)
	}
	return result
}

// Intersection returns a new Set containing elements that exist in both sets.
func Intersection[T comparable](s1, s2 Set[T]) Set[T] {
	result := make(Set[T])
	// Iterate over the smaller set for performance
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for v := range s1 {
		if s2.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns a new Set containing elements in s1 that are not in s2.
func Difference[T comparable](s1, s2 Set[T]) Set[T] {
	result := make(Set[T])
	for v := range s1 {
		if !s2.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// SymmetricDifference returns a new Set containing elements that are in either s1 or s2, but not both.
func SymmetricDifference[T comparable](s1, s2 Set[T]) Set[T] {
	result := make(Set[T])
	for v := range s1 {
		if !s2.Contains(v) {
			result.Add(v)
		}
	}
	for v := range s2 {
		if !s1.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// IsSubset returns true if all elements in the set exist in the other set.
func (s Set[T]) IsSubset(other Set[T]) bool {
	for v := range s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if the set contains all elements from the other set.
func (s Set[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}
