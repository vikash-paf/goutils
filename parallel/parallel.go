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
