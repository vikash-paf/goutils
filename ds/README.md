# `ds`

The `ds` package provides standard data structures optimized for memory management and thread safety.

## Structures

### PriorityQueue
A generic Min/Max Heap.
```go
// Create a Min-Heap
pq := ds.NewPriorityQueue[int](func(a, b int) bool { return a < b })

pq.Push(3)
pq.Push(1)

val, ok := pq.Pop() // val = 1, ok = true
```

### RingBuffer
A generic, fixed-size circular buffer. Great for keeping a sliding window of recent items (e.g., recent logs) without allocating new memory.
```go
rb := ds.NewRingBuffer[int](3)
rb.Push(1)
rb.Push(2)
rb.Push(3)
rb.Push(4) // Overwrites 1

fmt.Println(rb.Values()) // [2, 3, 4]
```

### Queue
A FIFO queue optimized to prevent memory leaks from continuous slice growth.
```go
var q ds.Queue[string]
q.Enqueue("job1")
val, ok := q.Dequeue()
```

### Stack
A standard LIFO stack.
```go
var s ds.Stack[int]
s.Push(42)
val, ok := s.Pop()
```

### BloomFilter
A space-efficient probabilistic data structure that is used to test whether an element is a member of a set. Useful for reducing expensive database lookups.
```go
// Create a Bloom filter for 10,000 items with a 1% false positive rate
bf := ds.NewBloomFilter(10000, 0.01)

bf.AddString("user_123")

// False positives are possible; false negatives are not.
if bf.ContainsString("user_123") {
    fmt.Println("User might exist")
}
```
