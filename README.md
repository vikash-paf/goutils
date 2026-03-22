# goutils

[![Go Reference](https://pkg.go.dev/badge/github.com/vikash/goutils.svg)](https://pkg.go.dev/github.com/vikash/goutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/vikash/goutils)](https://goreportcard.com/report/github.com/vikash/goutils)

`goutils` is a modern, zero-dependency, generic utility library for Go 1.21+. It provides heavily requested functional programming features, pointer manipulation helpers, native data structures (Sets, LRU Cache, Stacks), and concurrency transformers in an idiomatic and typesafe manner.

By leveraging Go generics, `goutils` eliminates the need for reflection, keeping performance high and type safety strong.

## Installation

```bash
go get github.com/vikash/goutils
```

## Packages

### `set`
A native, generic Set data structure (`map[T]struct{}`) with standard mathematical set operations.
- `New` / `FromSlice`: Create sets quickly.
- `Add`, `Remove`, `Contains`: Basic operations.
- `Values`: Retrieve all elements as a slice.
- `Union`, `Intersection`, `Difference`, `SymmetricDifference`: Core mathematical operations.
- `IsSubset`, `IsSuperset`: Set comparison.

### `cache`
A generic, thread-safe Least Recently Used (LRU) Cache.
- `LRU[K, V]`: Safe for concurrent use via `sync.RWMutex`.
- `NewLRU`, `Set`, `Get`, `Contains`, `Remove`, `Clear`, `Len`: Manage your cache effortlessly.

### `ds`
Standard data structures optimized for memory management.
- `Stack[T]`: LIFO queue (`Push`, `Pop`, `Peek`, `Len`).
- `Queue[T]`: FIFO queue optimized to prevent memory leaks from continuous slice growth (`Enqueue`, `Dequeue`, `Peek`, `Len`).

### `opt`
Null-safe Optional types (similar to Rust's `Option`) to avoid pointers and `nil` checks.
- `Opt[T]`: Value-based wrapper.
- `Some`, `None`: Constructors.
- `IsSome`, `IsNone`: State checking.
- `Unwrap`, `UnwrapOr`, `Map`: Access or transform the value safely.

### `slice`
A robust collection of slice manipulation functions.
- **Search & Logic**: `Find`, `FindIndex`, `Some`, `Every`.
- **Transformers**: `Map`, `Filter`, `Reduce`, `Chunk`, `Partition`.
- **Aggregation & Sorting**: `GroupBy`, `CountBy`, `Reverse`, `Shuffle`, `SortBy`, `SortByDesc`.
- **Deduplication & State**: `Unique`, `UniqueBy`, `DiffState` (Get added/removed items between two states).

### `syncx`
Advanced concurrency primitives for workflow control.
- `Debounce`: Delays invoking a function until a specified quiet period has elapsed.
- `Throttle`: Ensures a function is called at most once per specified duration.
- `Batcher[T]`: Collects items and flushes them to a processor either when a size limit is reached or a timeout occurs.

### `control`
Flow control and error handling helpers designed to reduce boilerplate.
- `If`: A generic ternary operator (`control.If(condition, trueVal, falseVal)`).
- `Must`: Standardizes the "panic on error" initialization pattern.
- `Coalesce`: Returns the first non-zero value from a list of arguments.
- `Try`: Executes a function and returns a fallback value if it errors.

### `mathx`
Generic mathematical operations using `cmp.Ordered` and custom numeric constraints.
- `Sum`, `Average`: Easily compute totals and means across any numeric slice type.
- `MinBy`, `MaxBy`: Find the min/max element in a slice of structs based on a field selector.
- `Clamp`: Restrict a value within a `[min, max]` range.

### `tuple`
Simple, generic data pairing.
- `Pair[L, R]`: A standard two-value tuple.
- `NewPair`: Quickly instantiate a pair.
- `Zip`: Combine two slices into a single slice of Pairs.
- `Unzip`: Split a slice of Pairs into two separate slices.

### `dict`
Helpers for map/dictionary manipulation.
- `Keys`, `Values`: Extract keys or values from a map.
- `Merge`: Combine multiple maps, prioritizing later arguments.
- `Invert`: Swap the keys and values in a map.
- `Omit`: Create a map without the specified keys.

### `ptr`
Safely work with pointers to primitive types or structs.
- `Of`, `Val`, `ValOrDefault`: Safely reference and dereference values.
- `Equal`: Compare the values of two pointers safely.

### `str`
String manipulation and formatting utilities.
- `IsBlank`, `Reverse`, `Truncate`.
- `ToCamelCase`, `ToSnakeCase`: Useful casing transformers.

### `async`
Simple async utilities.
- `MapAsync`: Perform a parallel map operation over a slice with a maximum concurrency limit.
- `Retry`, `RetryWithContext`: Attempt operations multiple times with backoff delays.

## License

MIT License. See [LICENSE](LICENSE) for more details.
