package resilience_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/vikash-paf/goutils/resilience"
	"time"
)

func ExampleCircuitBreaker() {
	cb := resilience.NewCircuitBreaker(2, 100*time.Millisecond)
	
	// First two failures
	cb.Execute(func() error { return errors.New("fail") })
	cb.Execute(func() error { return errors.New("fail") })
	
	// Third call should be rejected or return the failure
	err := cb.Execute(func() error { return nil })
	fmt.Println(err == resilience.ErrCircuitOpen)
	// Output: true
}

func ExampleRetry() {
	ctx := context.Background()
	config := resilience.DefaultRetryConfig
	config.MaxRetries = 2
	config.BaseDelay = 10 * time.Millisecond
	config.Jitter = false

	count := 0
	err := resilience.Retry(ctx, config, func(ctx context.Context) error {
		count++
		if count < 2 {
			return errors.New("temporary error")
		}
		return nil
	})
	fmt.Println(err)
	// Output: <nil>
}
