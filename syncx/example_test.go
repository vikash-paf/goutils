package syncx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/syncx"
	"time"
)

func ExampleBatcher() {
	processed := 0
	b := syncx.NewBatcher(2, 50*time.Millisecond, func(batch []int) {
		processed += len(batch)
	})

	b.Add(1)
	b.Add(2) // Triggers flush
	
	time.Sleep(10 * time.Millisecond) // Wait for async processing
	fmt.Println(processed)
	// Output: 2
}

func ExampleDebounce() {
	count := 0
	debounced := syncx.Debounce(50*time.Millisecond, func() {
		count++
	})

	debounced()
	debounced()
	debounced()

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count)
	// Output: 1
}
