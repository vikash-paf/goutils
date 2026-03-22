package resilience

import (
	"context"
	"math/rand"
	"time"
)

// RetryConfig holds the configuration for the Retry function.
type RetryConfig struct {
	MaxRetries int
	BaseDelay  time.Duration
	MaxDelay   time.Duration
	Jitter     bool // If true, adds up to 50% random jitter to the delay
}

// DefaultRetryConfig provides sensible defaults for a retry strategy.
var DefaultRetryConfig = RetryConfig{
	MaxRetries: 3,
	BaseDelay:  100 * time.Millisecond,
	MaxDelay:   2 * time.Second,
	Jitter:     true,
}

// Retry executes the provided function, retrying it according to the RetryConfig
// if it returns an error. The operation will exit early if the context is canceled.
func Retry(ctx context.Context, config RetryConfig, operation func(context.Context) error) error {
	var err error
	delay := config.BaseDelay

	for i := 0; i <= config.MaxRetries; i++ {
		err = operation(ctx)
		if err == nil {
			return nil
		}

		if i == config.MaxRetries {
			break
		}

		currentDelay := delay
		if config.Jitter {
			// Add up to 50% jitter to prevent thundering herds
			jitterAmount := time.Duration(rand.Int63n(int64(currentDelay)/2 + 1))
			currentDelay += jitterAmount
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(currentDelay):
		}

		// Exponential backoff
		delay *= 2
		if delay > config.MaxDelay {
			delay = config.MaxDelay
		}
	}

	return err
}
