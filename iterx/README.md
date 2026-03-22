# `iterx`

The `iterx` package provides functional utilities for working with Go 1.23+ iterators.

## Functional Utilities

### Map, Filter, Reduce, Take
Standard functional primitives for iterators:
```go
nums := []int{1, 2, 3, 4, 5}
seq := iterx.FromSlice(nums)

evens := iterx.Filter(seq, func(n int) bool { return n%2 == 0 })
doubled := iterx.Map(evens, func(n int) int { return n * 2 })
```

### GroupBy
Collects elements from an iterator into a map grouped by a selector function.
```go
nums := []int{1, 2, 3, 4, 5, 6}
seq := iterx.FromSlice(nums)

// Group into evens and odds
grouped := iterx.GroupBy(seq, func(n int) string {
    if n%2 == 0 { return "even" }
    return "odd"
})
// Result: {"even": [2, 4, 6], "odd": [1, 3, 5]}
```
