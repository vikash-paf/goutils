package syncx

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestPromise(t *testing.T) {
	p := NewPromise[int]()

	go func() {
		time.Sleep(50 * time.Millisecond)
		p.Resolve(42)
	}()

	val, err := p.Await(context.Background())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 42 {
		t.Errorf("Expected 42, got %v", val)
	}
}

func TestPromiseReject(t *testing.T) {
	p := NewPromise[int]()
	errFail := errors.New("fail")

	go func() {
		p.Reject(errFail)
	}()

	_, err := p.Await(context.Background())
	if err != errFail {
		t.Errorf("Expected error %v, got %v", errFail, err)
	}
}

func TestPromiseTimeout(t *testing.T) {
	p := NewPromise[int]()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := p.Await(ctx)
	if err != context.DeadlineExceeded {
		t.Errorf("Expected DeadlineExceeded, got %v", err)
	}
}
