package slice_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/slice"
)

func ExampleMap() {
	nums := []int{1, 2, 3}
	doubled := slice.Map(nums, func(n int) int { return n * 2 })
	fmt.Println(doubled)
	// Output: [2 4 6]
}

func ExampleFilter() {
	nums := []int{1, 2, 3, 4}
	evens := slice.Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens)
	// Output: [2 4]
}

func ExampleUnique() {
	nums := []int{1, 2, 2, 3, 1}
	unique := slice.Unique(nums)
	fmt.Println(unique)
	// Output: [1 2 3]
}

func ExampleChunk() {
	nums := []int{1, 2, 3, 4, 5}
	chunks := slice.Chunk(nums, 2)
	fmt.Println(chunks)
	// Output: [[1 2] [3 4] [5]]
}

func ExampleFind() {
	nums := []int{1, 2, 3, 4}
	found := slice.Find(nums, func(n int) bool { return n > 2 })
	fmt.Println(*found)
	// Output: 3
}

func ExamplePartition() {
	nums := []int{1, 2, 3, 4}
	passed, failed := slice.Partition(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(passed, failed)
	// Output: [2 4] [1 3]
}
