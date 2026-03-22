package syncx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/syncx"
	"sync"
	"sync/atomic"
	"time"
)

func ExampleBatcher() {
	var wg sync.WaitGroup
	wg.Add(1)

	b := syncx.NewBatcher(2, 50*time.Millisecond, func(batch []int) {
		fmt.Printf("Processed batch of size %d\n", len(batch))
		wg.Done()
	})

	b.Add(1)
	b.Add(2) // Triggers flush

	wg.Wait()
	// Output: Processed batch of size 2
}

func ExampleDebounce() {
	var wg sync.WaitGroup
	wg.Add(1)
	var count atomic.Int32

	debounced := syncx.Debounce(50*time.Millisecond, func() {
		count.Add(1)
		wg.Done()
	})

	debounced()
	debounced()
	debounced()

	wg.Wait()
	fmt.Println(count.Load())
	// Output: 1
}
