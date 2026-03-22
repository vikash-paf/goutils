// Package resilience provides fault-tolerance patterns to prevent cascading system failures,
// such as the Circuit Breaker pattern.
//
// Usage:
//
//	cb := resilience.NewCircuitBreaker(3, 5*time.Second)
//	err := cb.Execute(func() error {
//	    return http.Get("http://flaky-service.com")
//	})
package resilience

import (
	"errors"
	"sync"
	"time"
)

// State represents the current state of the Circuit Breaker.
type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

var (
	// ErrCircuitOpen is returned when the CircuitBreaker is open and rejecting calls.
	ErrCircuitOpen = errors.New("circuit breaker is open")
)

// CircuitBreaker prevents repeated execution of an operation that is likely to fail.
type CircuitBreaker struct {
	maxFailures  int
	resetTimeout time.Duration

	mu          sync.Mutex
	state       State
	failures    int
	lastFailure time.Time
}

// NewCircuitBreaker creates a new CircuitBreaker.
func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        StateClosed,
	}
}

// Execute runs the provided function. If the circuit is open, it returns ErrCircuitOpen immediately.
func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()

	if cb.state == StateOpen {
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			cb.state = StateHalfOpen
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	}
	cb.mu.Unlock()

	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.maxFailures || cb.state == StateHalfOpen {
			cb.state = StateOpen
		}
		return err
	}

	// Success
	cb.failures = 0
	cb.state = StateClosed
	return nil
}

// ExecuteGen allows running a function that returns a value and an error safely through the CircuitBreaker.
func ExecuteGen[T any](cb *CircuitBreaker, fn func() (T, error)) (T, error) {
	var result T
	err := cb.Execute(func() error {
		res, err := fn()
		result = res
		return err
	})
	return result, err
}

// GetState returns the current state of the circuit breaker.
func (cb *CircuitBreaker) GetState() State {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == StateOpen && time.Since(cb.lastFailure) > cb.resetTimeout {
		cb.state = StateHalfOpen
	}
	return cb.state
}
