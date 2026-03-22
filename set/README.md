# `set`

The `set` package implements a generic, native-feeling Set data structure backed by a `map[T]struct{}`. It provides constant-time $O(1)$ lookups and mathematical set operations.

## Usage

### Basic Operations
```go
s := set.New("apple", "banana")
s.Add("cherry")
s.Remove("banana")

fmt.Println(s.Contains("apple")) // true
fmt.Println(s.Values())          // ["apple", "cherry"] (unordered)
```

### Mathematical Set Operations
```go
s1 := set.New(1, 2, 3)
s2 := set.New(3, 4, 5)

set.Union(s1, s2)               // [1, 2, 3, 4, 5]
set.Intersection(s1, s2)        // [3]
set.Difference(s1, s2)          // [1, 2] (in s1, not in s2)
set.SymmetricDifference(s1, s2) // [1, 2, 4, 5] (in either, but not both)
```
