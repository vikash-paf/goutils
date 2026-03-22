# `resilience`

The `resilience` package provides fault-tolerance patterns to prevent cascading system failures.

## CircuitBreaker
Prevents repeated execution of an operation that is likely to fail. It transitions through `StateClosed` (normal), `StateOpen` (failing, rejecting calls), and `StateHalfOpen` (testing recovery).

### Usage
```go
// Opens the circuit after 3 consecutive failures, retries after 5 seconds.
cb := resilience.NewCircuitBreaker(3, 5*time.Second)

err := cb.Execute(func() error {
    return http.Get("http://flaky-service.com")
})

if err == resilience.ErrCircuitOpen {
    fmt.Println("Service is down, rejecting call immediately without network I/O.")
}
```
