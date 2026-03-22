# `algo`

The `algo` package provides generic implementations of common algorithms optimized for performance and type safety.

## Algorithms

### TopK
A generic function to efficiently find the $K$ largest or smallest elements in a slice using a min-heap. Ideal for finding "top N" items without the overhead of sorting the entire dataset. Time complexity: $O(N \log K)$.

```go
items := []int{10, 4, 25, 8, 3, 1, 15}

// Find the 3 largest numbers. Less function returns a < b to keep larger elements.
largest3 := algo.TopK(items, 3, func(a, b int) bool { return a < b })
fmt.Println(largest3) // Output: [10 15 25] (order within result not guaranteed)

// Find the 2 smallest numbers. Less function returns a > b to keep smaller elements.
smallest2 := algo.TopK(items, 2, func(a, b int) bool { return a > b })
fmt.Println(smallest2) // Output: [1 3]
```

### BinarySearch
Searches for a target value in a sorted slice of items. Note: The slice must be sorted prior to calling.

```go
items := []int{1, 3, 5, 7, 9}
idx := algo.BinarySearch(items, 5, func(v int) int { return v })
fmt.Println(idx) // Output: 2
```
