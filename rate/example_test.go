package rate_test

import (
	"context"
	"fmt"
	"github.com/vikash-paf/goutils/rate"
	"time"
)

func ExampleTokenBucket_Allow() {
	tb := rate.NewTokenBucket(1, time.Minute)
	fmt.Println(tb.Allow()) // Consumes the only token
	fmt.Println(tb.Allow()) // No tokens left
	// Output:
	// true
	// false
}

func ExampleTokenBucket_Wait() {
	tb := rate.NewTokenBucket(1, 50*time.Millisecond)
	_ = tb.Allow() // Consume first token

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := tb.Wait(ctx) // Blocks until refilled
	fmt.Println(err == nil)
	// Output: true
}
