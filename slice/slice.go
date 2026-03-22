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

// Filter iterates over a slice, returning a new slice with all elements
// that pass the predicate function.
func Filter[T any](items []T, predicate func(T) bool) []T {
	if items == nil {
		return nil
	}
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	// Return empty slice instead of nil if items was empty or none passed,
	// but strictly nil if nil passed. Let's stick to returning an empty slice
	// when original isn't nil, for better composition.
	if result == nil {
		return []T{}
	}
	return result
}

// Reduce reduces a slice to a single value by iteratively combining each
// element of the slice with an accumulator value.
func Reduce[T, R any](items []T, accumulator func(R, T) R, initial R) R {
	result := initial
	for _, item := range items {
		result = accumulator(result, item)
	}
	return result
}

// Unique returns a new slice containing only the unique elements of the given slice.
// The order of elements is preserved.
func Unique[T comparable](items []T) []T {
	if items == nil {
		return nil
	}
	result := make([]T, 0, len(items))
	seen := make(map[T]struct{})
	for _, item := range items {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// UniqueBy returns a new slice containing only the unique elements of the given slice,
// based on the value returned by the iteratee function. The order of elements is preserved.
func UniqueBy[T any, U comparable](items []T, iteratee func(T) U) []T {
	if items == nil {
		return nil
	}
	result := make([]T, 0, len(items))
	seen := make(map[U]struct{})
	for _, item := range items {
		key := iteratee(item)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Chunk splits a slice into an array of smaller slices, each of the given size.
// If the slice cannot be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](items []T, size int) [][]T {
	if items == nil {
		return nil
	}
	if size <= 0 {
		panic("size must be greater than 0")
	}

	chunks := make([][]T, 0, (len(items)+size-1)/size)
	for i := 0; i < len(items); i += size {
		end := i + size
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}
	return chunks
}

// GroupBy groups the elements of a slice according to the key returned by the iteratee function.
func GroupBy[T any, K comparable](items []T, keyFunc func(T) K) map[K][]T {
	if items == nil {
		return nil
	}
	result := make(map[K][]T)
	for _, item := range items {
		key := keyFunc(item)
		result[key] = append(result[key], item)
	}
	return result
}
