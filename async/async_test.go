package async

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestMapAsync(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := MapAsync(input, func(val int) int {
		time.Sleep(10 * time.Millisecond) // Simulate work
		return val * 2
	}, 2)

	want := []int{2, 4, 6, 8, 10}
	if len(got) != len(want) {
		t.Fatalf("MapAsync() returned %d items, want %d", len(got), len(want))
	}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("MapAsync()[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestRetry(t *testing.T) {
	attempts := 0
	err := Retry(3, 10*time.Millisecond, func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil
	})

	if err != nil {
		t.Errorf("Retry() failed after 3 attempts")
	}

	attempts = 0
	err = Retry(2, 10*time.Millisecond, func() error {
		attempts++
		return errors.New("permanent error")
	})

	if err == nil {
		t.Errorf("Retry() should have returned an error")
	}
}

func TestRetryWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err := RetryWithContext(ctx, 3, 100*time.Millisecond, func() error {
		return errors.New("should cancel before retry")
	})

	if err != context.Canceled {
		t.Errorf("RetryWithContext() returned %v, want context.Canceled", err)
	}
}

func ExampleMapAsync() {
	input := []int{1, 2, 3}
	results := MapAsync(input, func(val int) int {
		return val * val
	}, 2) // Max 2 concurrent workers

	fmt.Println(results)
	// Output: [1 4 9]
}
