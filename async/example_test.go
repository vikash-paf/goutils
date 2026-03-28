package async_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/vikash-paf/goutils/async"
	"time"
)

func ExampleMapAsync() {
	items := []int{1, 2, 3, 4, 5}
	results := async.MapAsync(items, func(n int) int {
		return n * 2
	}, 2)
	fmt.Println(results)
	// Output: [2 4 6 8 10]
}

func ExampleRetry() {
	count := 0
	err := async.Retry(3, 10*time.Millisecond, func() error {
		count++
		if count < 3 {
			return errors.New("temporary error")
		}
		return nil
	})
	fmt.Println(err)
	// Output: <nil>
}

func ExampleRetryWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := async.RetryWithContext(ctx, 5, 20*time.Millisecond, func() error {
		return errors.New("always fail")
	})
	fmt.Println(err != nil)
	// Output: true
}
