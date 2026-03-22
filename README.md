# goutils

[![Go Reference](https://pkg.go.dev/badge/github.com/vikash/goutils.svg)](https://pkg.go.dev/github.com/vikash/goutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/vikash/goutils)](https://goreportcard.com/report/github.com/vikash/goutils)

`goutils` is a modern, zero-dependency, generic utility library for Go 1.21+. It provides heavily requested functional programming features, pointer manipulation helpers, map utilities, and string transformers in an idiomatic and typesafe manner.

By leveraging Go generics, `goutils` eliminates the need for reflection, keeping performance high and type safety strong.

## Installation

```bash
go get github.com/vikash/goutils
```

## Packages

### `slice`
A robust collection of slice manipulation functions.
- `Map`: Transform elements of a slice into another slice.
- `Filter`: Extract elements that match a condition.
- `Reduce`: Accumulate a slice down to a single value.
- `Unique`: Deduplicate primitive slices.
- `UniqueBy`: Deduplicate slices of structs or interfaces using a key.
- `Chunk`: Split a slice into smaller, fixed-size batches.
- `GroupBy`: Group slice elements into a map by a specific key.

### `dict`
Helpers for map/dictionary manipulation.
- `Keys`: Extract all keys from a map as a slice.
- `Values`: Extract all values from a map as a slice.
- `Merge`: Combine multiple maps, prioritizing later arguments.
- `Invert`: Swap the keys and values in a map.
- `Omit`: Create a map without the specified keys.

### `ptr`
Safely work with pointers to primitive types or structs.
- `Of`: Quickly get a pointer to a value (e.g., `ptr.Of("test")`).
- `Val`: Safely dereference a pointer, returning the zero value if nil.
- `ValOrDefault`: Safely dereference a pointer, returning a fallback value if nil.
- `Equal`: Compare the values of two pointers safely, even if one or both are nil.

### `str`
String manipulation and formatting utilities.
- `IsBlank`: Check if a string is empty or contains only whitespace.
- `Reverse`: Reverse a string, supporting UTF-8/Runes.
- `Truncate`: Shorten a string safely and add an omission token like `...`.
- `ToCamelCase`: Transform strings to camelCase format.
- `ToSnakeCase`: Transform strings to snake_case format.

### `async`
Concurrency and async workflow utilities.
- `MapAsync`: Perform a parallel map operation over a slice with a maximum concurrency limit.
- `Retry`: Attempt an operation multiple times with a backoff delay.
- `RetryWithContext`: Similar to Retry but respects context cancellation.

## License

MIT License. See [LICENSE](LICENSE) for more details.
