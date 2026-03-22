package ds_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/ds"
)

func ExampleStack() {
	var s ds.Stack[int]
	s.Push(1)
	s.Push(2)
	val, _ := s.Pop()
	fmt.Println(val)
	// Output: 2
}

func ExampleQueue() {
	var q ds.Queue[int]
	q.Enqueue(1)
	q.Enqueue(2)
	val, _ := q.Dequeue()
	fmt.Println(val)
	// Output: 1
}

func ExamplePriorityQueue() {
	pq := ds.NewPriorityQueue[int](func(a, b int) bool { return a < b })
	pq.Push(3)
	pq.Push(1)
	pq.Push(2)
	val, _ := pq.Pop()
	fmt.Println(val)
	// Output: 1
}

func ExampleRingBuffer() {
	rb := ds.NewRingBuffer[int](3)
	rb.Push(1)
	rb.Push(2)
	rb.Push(3)
	rb.Push(4) // Overwrites 1
	fmt.Println(rb.Values())
	// Output: [2 3 4]
}

func ExampleBloomFilter() {
	bf := ds.NewBloomFilter(100, 0.01)
	bf.AddString("hello")
	fmt.Println(bf.ContainsString("hello"))
	fmt.Println(bf.ContainsString("world"))
	// Output:
	// true
	// false
}

func ExampleUnionFind() {
	uf := ds.NewUnionFind[string]()
	uf.Add("A")
	uf.Add("B")
	uf.Union("A", "B")
	fmt.Println(uf.Connected("A", "B"))
	// Output: true
}

func ExampleFenwickTree() {
	ft := ds.NewFenwickTree[int](5)
	ft.Add(0, 10)
	ft.Add(2, 5)
	fmt.Println(ft.PrefixSum(2))
	// Output: 15
}

func ExampleTrie() {
	t := ds.NewTrie[int]()
	t.Insert("hello", 1)
	t.Insert("hell", 2)
	val, _ := t.Search("hello")
	fmt.Println(val)
	fmt.Println(t.StartsWith("hel"))
	// Output:
	// 1
	// true
}
