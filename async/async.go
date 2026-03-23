// Package async provides generic utility functions for concurrent programming,
// such as parallel mapping and retry logic with backoff.
//
// Usage:
//
//	results := async.MapAsync(urls, func(url string) string {
//	    return fetch(url)
//	}, 5)
package async

import (
	"context"
	"sync"
	"time"
)

// MapAsync applies a function asynchronously to each element of a slice and returns the results.
// It uses a maxConcurrency limit to prevent overwhelming system resources.
// The order of results is guaranteed to match the order of the input slice.
func MapAsync[T, U any](items []T, iteratee func(T) U, maxConcurrency int) []U {
	if items == nil {
		return nil
	}

	maxConcurrency = max(1, maxConcurrency)

	result := make([]U, len(items))
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxConcurrency)

	for i, item := range items {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire semaphore

		go func(index int, val T) {
			defer wg.Done()
			result[index] = iteratee(val)
			<-semaphore // Release semaphore
		}(i, item)
	}

	wg.Wait()
	return result
}

// Retry attempts to execute a function repeatedly until it succeeds or the max attempts are reached.
// It waits for the specified delay between attempts.
func Retry(attempts int, delay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		if i < attempts-1 {
			time.Sleep(delay)
		}
	}
	return err
}

// RetryWithContext is like Retry but respects context cancellation.
func RetryWithContext(ctx context.Context, attempts int, delay time.Duration, fn func() error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}

		if i < attempts-1 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
			}
		}
	}
	return err
}
