// Package parallel provides utilities for executing tasks in parallel
// and collecting their results. It uses sync.WaitGroup for efficient
// concurrency management.
//
// Usage:
//
//	results := parallel.Map(inputs, func(v int) int { return v * 2 })
package parallel

import "sync"

// Map applies a function to each element in a slice in parallel and returns
// a new slice with the results.
func Map[T, U any](items []T, fn func(T) U) []U {
	if len(items) == 0 {
		return nil
	}

	results := make([]U, len(items))
	var wg sync.WaitGroup
	wg.Add(len(items))

	for i, item := range items {
		go func(i int, item T) {
			defer wg.Done()
			results[i] = fn(item)
		}(i, item)
	}

	wg.Wait()
	return results
}

// ForEach executes a function for each element in a slice in parallel.
func ForEach[T any](items []T, fn func(T)) {
	if len(items) == 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(items))

	for _, item := range items {
		go func(item T) {
			defer wg.Done()
			fn(item)
		}(item)
	}

	wg.Wait()
}

// MapBatched applies a function to each element in a slice in parallel,
// but limits the number of concurrent executions to the specified batch size.
func MapBatched[T, U any](items []T, batchSize int, fn func(T) U) []U {
	if len(items) == 0 {
		return nil
	}
	if batchSize <= 0 {
		return Map(items, fn)
	}

	results := make([]U, len(items))
	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}

		batchItems := items[i:end]
		var wg sync.WaitGroup
		wg.Add(len(batchItems))

		for j, item := range batchItems {
			go func(index int, val T) {
				defer wg.Done()
				results[index] = fn(val)
			}(i+j, item)
		}
		wg.Wait()
	}
	return results
}

// ForEachBatched executes a function for each element in a slice in parallel,
// but limits the number of concurrent executions to the specified batch size.
func ForEachBatched[T any](items []T, batchSize int, fn func(T)) {
	if len(items) == 0 {
		return
	}
	if batchSize <= 0 {
		ForEach(items, fn)
		return
	}

	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}

		batchItems := items[i:end]
		var wg sync.WaitGroup
		wg.Add(len(batchItems))

		for _, item := range batchItems {
			go func(val T) {
				defer wg.Done()
				fn(val)
			}(item)
		}
		wg.Wait()
	}
}
