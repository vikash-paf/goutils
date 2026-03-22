package rate

import (
	"context"
	"testing"
	"time"
)

func TestTokenBucketAllow(t *testing.T) {
	// Capacity 2, adds 1 token every 50ms
	tb := NewTokenBucket(2, 50*time.Millisecond)

	if !tb.Allow() {
		t.Error("Expected 1st Allow() to succeed")
	}
	if !tb.Allow() {
		t.Error("Expected 2nd Allow() to succeed")
	}
	if tb.Allow() {
		t.Error("Expected 3rd Allow() to fail immediately")
	}

	time.Sleep(60 * time.Millisecond)
	if !tb.Allow() {
		t.Error("Expected Allow() to succeed after refill")
	}
	if tb.Allow() {
		t.Error("Expected next Allow() to fail, bucket shouldn't be fully refilled yet")
	}
}

func TestTokenBucketWait(t *testing.T) {
	tb := NewTokenBucket(1, 50*time.Millisecond)

	// Consume the initial token
	_ = tb.Allow()

	ctx, cancel := context.WithTimeout(context.Background(), 75*time.Millisecond)
	defer cancel()

	err := tb.Wait(ctx)
	if err != nil {
		t.Errorf("Expected Wait to succeed within timeout, got error: %v", err)
	}
}

func TestTokenBucketWaitTimeout(t *testing.T) {
	tb := NewTokenBucket(1, 100*time.Millisecond)

	// Consume the initial token
	_ = tb.Allow()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := tb.Wait(ctx)
	if err == nil {
		t.Error("Expected Wait to fail due to context timeout")
	}
}
