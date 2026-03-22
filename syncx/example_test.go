package syncx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/syncx"
	"sync/atomic"
	"time"
)

func ExampleBatcher() {
	var processed atomic.Int32
	b := syncx.NewBatcher(2, 50*time.Millisecond, func(batch []int) {
		processed.Add(int32(len(batch)))
	})

	b.Add(1)
	b.Add(2) // Triggers flush
	
	time.Sleep(10 * time.Millisecond) // Wait for async processing
	fmt.Println(processed.Load())
	// Output: 2
}

func ExampleDebounce() {
	var count atomic.Int32
	debounced := syncx.Debounce(50*time.Millisecond, func() {
		count.Add(1)
	})

	debounced()
	debounced()
	debounced()

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count.Load())
	// Output: 1
}
