# `resilience`

The `resilience` package provides patterns for maintaining system stability and preventing cascading failures in distributed environments.

## Retry (with Exponential Backoff & Jitter)
A resilient retry mechanism that protects downstream services from thundering herds by employing exponential backoff and randomized jitter.
```go
ctx := context.Background()

config := resilience.RetryConfig{
    MaxRetries: 3,
    BaseDelay:  100 * time.Millisecond,
    MaxDelay:   2 * time.Second,
    Jitter:     true, // Adds up to 50% random jitter to delays
}

err := resilience.Retry(ctx, config, func(ctx context.Context) error {
    return makeNetworkCall(ctx)
})

if err != nil {
    fmt.Println("Operation finally failed:", err)
}
```
