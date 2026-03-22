package resilience

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetry_SuccessImmediate(t *testing.T) {
	attempts := 0
	err := Retry(context.Background(), DefaultRetryConfig, func(ctx context.Context) error {
		attempts++
		return nil
	})

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if attempts != 1 {
		t.Errorf("expected 1 attempt, got %d", attempts)
	}
}

func TestRetry_SuccessAfterFailures(t *testing.T) {
	config := RetryConfig{MaxRetries: 3, BaseDelay: 1 * time.Millisecond, MaxDelay: 10 * time.Millisecond, Jitter: false}
	attempts := 0

	err := Retry(context.Background(), config, func(ctx context.Context) error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil
	})

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}

func TestRetry_MaxRetriesExceeded(t *testing.T) {
	config := RetryConfig{MaxRetries: 2, BaseDelay: 1 * time.Millisecond, MaxDelay: 10 * time.Millisecond, Jitter: false}
	attempts := 0

	errString := "permanent error"
	err := Retry(context.Background(), config, func(ctx context.Context) error {
		attempts++
		return errors.New(errString)
	})

	if err == nil || err.Error() != errString {
		t.Errorf("expected %q, got %v", errString, err)
	}
	// 1 initial try + 2 retries = 3 attempts total.
	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}

func TestRetry_ContextCancellation(t *testing.T) {
	config := RetryConfig{MaxRetries: 5, BaseDelay: 50 * time.Millisecond, MaxDelay: 100 * time.Millisecond, Jitter: false}
	ctx, cancel := context.WithCancel(context.Background())
	attempts := 0

	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	err := Retry(ctx, config, func(ctx context.Context) error {
		attempts++
		return errors.New("failing")
	})

	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
	// It should cancel before the second attempt (waiting 50ms)
	if attempts != 1 {
		t.Errorf("expected 1 attempt, got %d", attempts)
	}
}

func ExampleRetry() {
	ctx := context.Background()

	var attempts int
	operation := func(ctx context.Context) error {
		attempts++
		if attempts < 2 {
			return errors.New("network timeout")
		}
		fmt.Println("Success on attempt:", attempts)
		return nil
	}

	err := Retry(ctx, DefaultRetryConfig, operation)
	if err != nil {
		fmt.Println("Operation failed:", err)
	}
	// Output: Success on attempt: 2
}
