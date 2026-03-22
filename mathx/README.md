# `mathx`

The `mathx` package provides generic mathematical operations using `cmp.Ordered` and custom numeric constraints. It works across ints, floats, and more.

## Usage

### Sum & Average
```go
total := mathx.Sum([]float64{10.5, 20.2, 5.3})
mean := mathx.Average([]int{1, 2, 3, 4, 5})
```

### MinBy & MaxBy
Find the min/max element in a slice of structs based on a field selector. Returns a pointer to the element (or nil if empty).
```go
users := []User{{Age: 30}, {Age: 20}, {Age: 40}}

youngest := mathx.MinBy(users, func(u User) int { return u.Age })
// youngest.Age == 20
```

### Clamp
Restrict a value within a `[min, max]` range.
```go
val := mathx.Clamp(150, 0, 100) // Returns 100
```
