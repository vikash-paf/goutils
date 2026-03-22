// Package rate provides rate limiting utilities.
package rate

import (
	"context"
	"sync"
	"time"
)

// TokenBucket is a thread-safe rate limiter based on the token bucket algorithm.
type TokenBucket struct {
	capacity   float64
	tokens     float64
	refillRate float64 // tokens per second
	lastRefill time.Time
	mu         sync.Mutex
}

// NewTokenBucket creates a new TokenBucket.
// capacity is the maximum burst size. refillInterval is the duration required to add 1 token.
func NewTokenBucket(capacity int, refillInterval time.Duration) *TokenBucket {
	if capacity <= 0 {
		panic("TokenBucket capacity must be > 0")
	}
	if refillInterval <= 0 {
		panic("TokenBucket refillInterval must be > 0")
	}

	refillRateSec := 1.0 / refillInterval.Seconds()

	return &TokenBucket{
		capacity:   float64(capacity),
		tokens:     float64(capacity),
		refillRate: refillRateSec,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

// Allow checks if a token is currently available and consumes it. Returns true if allowed.
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()
	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		return true
	}
	return false
}

// Wait blocks until a token is available or the context is canceled.
func (tb *TokenBucket) Wait(ctx context.Context) error {
	for {
		tb.mu.Lock()
		tb.refill()

		if tb.tokens >= 1.0 {
			tb.tokens -= 1.0
			tb.mu.Unlock()
			return nil
		}

		// Calculate time until next token
		missing := 1.0 - tb.tokens
		waitSec := missing / tb.refillRate
		waitTime := time.Duration(waitSec * float64(time.Second))
		tb.mu.Unlock()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(waitTime):
			// Loop again to attempt claiming the token
		}
	}
}
