package syncx

import (
	"sync"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	ps := NewPubSub[string]()
	defer ps.Close()

	ch1 := ps.Subscribe(10)
	ch2 := ps.Subscribe(10)

	ps.Publish("hello")

	msg1 := <-ch1
	msg2 := <-ch2

	if msg1 != "hello" || msg2 != "hello" {
		t.Errorf("expected 'hello', got '%s' and '%s'", msg1, msg2)
	}

	ps.Unsubscribe(ch1)
	ps.Publish("world")

	// ch1 is closed, reading from it yields zero value
	if _, ok := <-ch1; ok {
		t.Error("ch1 should be closed")
	}

	msg2 = <-ch2
	if msg2 != "world" {
		t.Errorf("expected 'world', got '%s'", msg2)
	}
}

func TestPubSub_NonBlocking(t *testing.T) {
	ps := NewPubSub[int]()
	defer ps.Close()

	// Zero buffer size forces dropping if consumer isn't perfectly synchronized
	ch1 := ps.Subscribe(0) 
	ch2 := ps.Subscribe(10)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond) // Ensure publish happens before consumer is ready
		// ch1 misses the message because it blocked and publisher dropped it
	}()

	ps.Publish(42) // Triggers select default case for ch1

	msg2 := <-ch2
	if msg2 != 42 {
		t.Errorf("expected ch2 to get message, got %d", msg2)
	}
	wg.Wait()
	
	// Ensure ch1 is indeed empty/didn't receive it because we didn't wait
	select {
	case <-ch1:
		t.Error("ch1 should not have message")
	default:
		// expected
	}
}
