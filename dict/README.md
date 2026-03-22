# `dict`

The `dict` package provides generic helpers for map/dictionary manipulation, saving you from writing repetitive iteration loops.

## Usage

### Keys & Values
Extract all keys or values into a slice.
```go
m := map[string]int{"a": 1, "b": 2}
keys := dict.Keys(m)     // ["a", "b"] (unordered)
values := dict.Values(m) // [1, 2] (unordered)
```

### Merge
Combine multiple maps. Later maps overwrite earlier identical keys.
```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 99, "c": 3}
merged := dict.Merge(m1, m2)
// {"a": 1, "b": 99, "c": 3}
```

### Invert
Swap keys and values.
```go
m := map[string]int{"a": 1, "b": 2}
inverted := dict.Invert(m)
// {1: "a", 2: "b"}
```

### Omit
Clone a map, omitting specific keys.
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
omitted := dict.Omit(m, "a", "c")
// {"b": 2}
```
