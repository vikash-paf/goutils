# `syncx`

The `syncx` package provides advanced concurrency primitives and workflow controllers.

## Worker Pool
A robust worker pool to handle massive streams of jobs asynchronously.
```go
pool := syncx.NewPool[string, int](5, func(job string) int {
    // Process job
    return len(job)
})

pool.Submit("hello")
pool.Close() // Wait for workers to finish

for result := range pool.Results() {
    fmt.Println(result)
}
```

## Batcher
Accumulates items and flushes them to a callback when a size limit is reached OR a timeout occurs. Excellent for bulk database inserts.
```go
b := syncx.NewBatcher(100, 5*time.Second, func(batch []string) {
    db.InsertMany(batch)
})
b.Add("item1")
b.Add("item2")
```

## Debounce & Throttle
Rate limit function execution.
```go
// Debounce: Delay execution until 100ms of quiet time
debounced := syncx.Debounce(100*time.Millisecond, myFunc)

// Throttle: Execute at most once per 100ms
throttled := syncx.Throttle(100*time.Millisecond, myFunc)
```
