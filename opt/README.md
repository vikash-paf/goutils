# `opt`

The `opt` package implements a null-safe Optional type (`Opt[T]`), similar to Rust's `Option`. It eliminates the need to use pointers (`*string`) just to represent missing data, preventing catastrophic nil pointer dereferences.

## Usage

### Creation & Checking
```go
var val opt.Opt[int] = opt.Some(42)
var missing opt.Opt[int] = opt.None[int]()

fmt.Println(val.IsSome()) // true
fmt.Println(missing.IsNone()) // true
```

### Safe Access
```go
val := opt.Some("hello")

// UnwrapOr returns the value, or a default if missing
fmt.Println(val.UnwrapOr("default")) // "hello"
fmt.Println(opt.None[string]().UnwrapOr("default")) // "default"

// Unwrap panics if None. Use with caution.
fmt.Println(val.Unwrap()) // "hello"
```

### Mapping
Transform an optional value safely.
```go
val := opt.Some(10)
mapped := opt.Map(val, func(v int) string { return fmt.Sprintf("Num: %d", v) })
// mapped is an opt.Opt[string] containing "Num: 10"
```
