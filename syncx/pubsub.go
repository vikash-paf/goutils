package syncx

import "sync"

// PubSub is a generic, thread-safe, in-memory publish-subscribe broker.
// It allows decoupled components to communicate via channels without blocking
// the publisher if a consumer is too slow.
type PubSub[T any] struct {
	mu          sync.RWMutex
	subscribers map[chan T]struct{}
}

// NewPubSub creates a new publish-subscribe broker.
func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		subscribers: make(map[chan T]struct{}),
	}
}

// Subscribe creates and returns a channel that will receive published messages.
// The bufferSize prevents the subscriber from missing messages under load.
func (ps *PubSub[T]) Subscribe(bufferSize int) chan T {
	ch := make(chan T, bufferSize)
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	ps.subscribers[ch] = struct{}{}
	return ch
}

// Unsubscribe removes a specific channel from the broker and closes it.
func (ps *PubSub[T]) Unsubscribe(ch chan T) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.subscribers[ch]; ok {
		delete(ps.subscribers, ch)
		close(ch)
	}
}

// Publish sends a message to all currently active subscribers.
// If a subscriber's channel buffer is full, the message is dropped for that
// specific subscriber entirely. This ensures that one slow consumer cannot
// cause the entire application to deadlock or slow down the publisher.
func (ps *PubSub[T]) Publish(msg T) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for ch := range ps.subscribers {
		select {
		case ch <- msg:
		default:
			// Non-blocking write: channel is full, drop the message.
		}
	}
}

// Close removes all subscribers and cleanly terminates all active channels.
func (ps *PubSub[T]) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for ch := range ps.subscribers {
		close(ch)
	}
	// Clear the map
	ps.subscribers = make(map[chan T]struct{})
}
