# `rate`

The `rate` package provides efficient rate-limiting utilities.

## TokenBucket
An industry-standard Token Bucket rate limiter, supporting both synchronous (immediate drop) and asynchronous (wait/block) paradigms.

### Usage
```go
// Burst capacity of 10, adds 1 token every 100ms
tb := rate.NewTokenBucket(10, 100*time.Millisecond)

// Synchronous: Returns true if allowed, false if rate limited
if !tb.Allow() {
    return http.StatusTooManyRequests
}

// Asynchronous: Blocks until a token is available or context cancels
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

err := tb.Wait(ctx)
if err != nil {
    // Timed out waiting for rate limit
}
```
