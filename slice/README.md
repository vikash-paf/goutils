# `slice`

The `slice` package provides a robust collection of generic slice manipulation functions, taking inspiration from functional programming (map, filter, reduce) and data wrangling utilities.

## Usage

### Map
Transforms elements of a slice.
```go
numbers := []int{1, 2, 3}
strings := slice.Map(numbers, strconv.Itoa)
// ["1", "2", "3"]
```

### Filter
Extracts elements that match a condition.
```go
words := []string{"hello", "go", "world"}
long := slice.Filter(words, func(w string) bool { return len(w) > 3 })
// ["hello", "world"]
```

### Reduce
Accumulates a slice down to a single value.
```go
sum := slice.Reduce([]int{1, 2, 3}, func(acc, val int) int { return acc + val }, 0)
// 6
```

### Unique / UniqueBy
Deduplicates primitive slices, or slices of structs based on a key.
```go
uniq := slice.Unique([]int{1, 2, 2, 3})
// [1, 2, 3]

users := []User{{ID: 1}, {ID: 1}, {ID: 2}}
uniqUsers := slice.UniqueBy(users, func(u User) int { return u.ID })
```

### Chunk / Partition / GroupBy
Split or group slices based on criteria.
```go
chunks := slice.Chunk([]int{1, 2, 3, 4, 5}, 2)
// [[1, 2], [3, 4], [5]]

passed, failed := slice.Partition([]int{1, 2, 3}, func(n int) bool { return n%2 == 0 })
// passed: [2], failed: [1, 3]
```

### DiffState
Compare two slices and extract added and removed items. Perfect for syncing database states.
```go
added, removed := slice.DiffState([]string{"admin", "editor"}, []string{"editor", "viewer"})
// added: ["viewer"], removed: ["admin"]
```
