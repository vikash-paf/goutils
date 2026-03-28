# goutils

[![Go Reference](https://pkg.go.dev/badge/github.com/vikash-paf/goutils.svg)](https://pkg.go.dev/github.com/vikash-paf/goutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/vikash-paf/goutils)](https://goreportcard.com/report/github.com/vikash-paf/goutils)
[![CI](https://github.com/vikash-paf/goutils/actions/workflows/ci.yml/badge.svg)](https://github.com/vikash-paf/goutils/actions/workflows/ci.yml)

`goutils` is a modern, zero-dependency, generic utility library for Go 1.23+. It eliminates boilerplate by providing highly-tested, idiomatic utilities for collections, synchronization, error handling, and data structures.

## Installation

```bash
go get github.com/vikash-paf/goutils
```

## Packages

### Core Data Structures & Collections

#### `slice`
A robust collection of slice manipulation functions.
- **Search & Logic**: `Find`, `FindIndex`, `Some`, `Every`.
- **Transformers**: `Map`, `Filter`, `Reduce`, `Chunk`, `Partition`.
- **Aggregation & Sorting**: `GroupBy`, `CountBy`, `Reverse`, `Shuffle`, `SortBy`, `SortByDesc`.
- **Deduplication**: `Unique`, `UniqueBy`, `DiffState` (Get added/removed items between two states).

#### `set`
A native, generic Set data structure (`map[T]struct{}`) with standard mathematical set operations.
- `New`, `Add`, `Remove`, `Contains`, `Values`.
- `Union`, `Intersection`, `Difference`, `SymmetricDifference`, `IsSubset`, `IsSuperset`.

#### `dict`
Helpers for map/dictionary manipulation.
- `Keys`, `Values`, `Merge`, `Invert`, `Omit`.

#### `tuple`
Simple, generic data pairing.
- `Pair[L, R]`, `NewPair`, `Zip`, `Unzip`.

#### `ds` & `cache` (Advanced Structures)
- **`ds.PriorityQueue[T]`**: A generic Min/Max Heap.
- **`ds.RingBuffer[T]`**: A generic, fixed-size circular buffer.
- **`ds.Stack[T]`** / **`ds.Queue[T]`**: Standard queues and stacks optimized for memory management.
- **`cache.LRU[K, V]`**: A concurrent-safe Least Recently Used Cache.

### Resilience & Concurrency

#### `syncx`
Advanced concurrency primitives and worker pools.
- **`Pool[Job, Result]`**: A robust worker pool. Send jobs, read results dynamically.
- **`Batcher[T]`**: Accumulates items and flushes them to a callback when a size limit is reached or a timeout occurs.
- **`Debounce`** / **`Throttle`**: Execution rate limiting for functions.

#### `resilience`
Fault-tolerance patterns.
- **`CircuitBreaker`**: Prevents repeated execution of operations likely to fail. Transitions through `StateClosed`, `StateOpen`, and `StateHalfOpen` with configurable thresholds.

#### `rate`
Rate limiting utilities.
- **`TokenBucket`**: Industry-standard synchronous and asynchronous (`Wait`) rate limiting.

### Control Flow & Utilities

#### `control` & `opt`
Flow control, Optional types, and error handling.
- **`control.If`**: A generic ternary operator (`control.If(condition, true, false)`).
- **`control.Must`**: Standardizes the "panic on error" initialization pattern.
- **`control.Coalesce`**: Returns the first non-zero value.
- **`control.Try`**: Executes a function, returning a fallback value if it errors.
- **`opt.Opt[T]`**: Null-safe Optional types (similar to Rust's `Option`) featuring `Some`, `None`, and `Unwrap`.

#### `mathx`
Generic mathematical operations using `cmp.Ordered`.
- `Sum`, `Average`, `Clamp`.
- `MinBy`, `MaxBy`: Find the min/max element in a slice of structs based on a field selector.

#### `str`
String manipulation and formatting utilities.
- `IsBlank`, `Reverse`, `Truncate`, `ToCamelCase`, `ToSnakeCase`.

#### `timex`
Time manipulation extensions.
- `StartOfDay`, `EndOfDay`, `StartOfWeek`.
- `IsWeekend`, `IsWeekday`, `AddBusinessDays`.

## Development

### Git Hooks

This project uses git hooks to ensure code quality. To set up the hooks, run:

```bash
git config core.hooksPath .githooks
```

The `pre-commit` hook automatically runs `gofmt -s -w` on staged Go files to ensure consistent formatting.

## License

MIT License. See [LICENSE](LICENSE) for more details.
