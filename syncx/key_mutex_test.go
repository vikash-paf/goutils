package syncx

import (
	"sync"
	"testing"
	"time"
)

func TestKeyMutex(t *testing.T) {
	km := NewKeyMutex[string]()
	var counter1, counter2 int
	var wg sync.WaitGroup

	// Key 1
	wg.Add(2)
	go func() {
		defer wg.Done()
		km.Lock("user1")
		counter1++
		time.Sleep(10 * time.Millisecond)
		km.Unlock("user1")
	}()
	go func() {
		defer wg.Done()
		km.Lock("user1")
		counter1++
		km.Unlock("user1")
	}()

	// Key 2 (should not be blocked by Key 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		km.Lock("user2")
		counter2++
		km.Unlock("user2")
	}()

	wg.Wait()
	if counter1 != 2 || counter2 != 1 {
		t.Errorf("Counters failed: c1=%d, c2=%d", counter1, counter2)
	}
}
