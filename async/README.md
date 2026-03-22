# `async`

The `async` package provides basic concurrency handlers.

## Usage

### MapAsync
Performs a parallel map operation over a slice with a maximum concurrency limit.
```go
urls := []string{"http://a.com", "http://b.com"}
results := async.MapAsync(urls, func(url string) string {
    return fetch(url)
}, 5) // max 5 concurrent routines
```

### Retry / RetryWithContext
Attempt an operation multiple times with a backoff delay.
```go
err := async.Retry(3, 100*time.Millisecond, func() error {
    return connectDB()
})
```
