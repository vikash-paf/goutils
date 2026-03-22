package resilience

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(2, 50*time.Millisecond)

	if cb.GetState() != StateClosed {
		t.Error("Initial state should be Closed")
	}

	failFn := func() error { return errors.New("fail") }
	successFn := func() error { return nil }

	// First failure
	_ = cb.Execute(failFn)
	if cb.GetState() != StateClosed {
		t.Error("State should be Closed after 1 failure")
	}

	// Second failure -> Opens Circuit
	_ = cb.Execute(failFn)
	if cb.GetState() != StateOpen {
		t.Error("State should be Open after 2 failures")
	}

	// Third attempt -> Should fail fast with ErrCircuitOpen
	err := cb.Execute(successFn)
	if err != ErrCircuitOpen {
		t.Errorf("Expected ErrCircuitOpen, got %v", err)
	}

	// Wait for reset timeout
	time.Sleep(60 * time.Millisecond)

	if cb.GetState() != StateHalfOpen {
		t.Error("State should be HalfOpen after timeout")
	}

	// Test success to close the circuit
	err = cb.Execute(successFn)
	if err != nil {
		t.Errorf("Expected success, got %v", err)
	}

	if cb.GetState() != StateClosed {
		t.Error("State should be Closed after successful attempt in HalfOpen state")
	}
}

func TestExecuteGen(t *testing.T) {
	cb := NewCircuitBreaker(1, 10*time.Millisecond)

	val, err := ExecuteGen(cb, func() (int, error) {
		return 42, nil
	})

	if err != nil || val != 42 {
		t.Errorf("Expected 42, nil; got %v, %v", val, err)
	}

	// Trip it
	_, _ = ExecuteGen(cb, func() (int, error) {
		return 0, errors.New("fail")
	})

	// Try again -> Open
	val, err = ExecuteGen(cb, func() (int, error) {
		return 42, nil
	})

	if err != ErrCircuitOpen || val != 0 {
		t.Errorf("Expected 0, ErrCircuitOpen; got %v, %v", val, err)
	}
}

func ExampleCircuitBreaker() {
	// Opens the circuit after 2 consecutive failures, retries after 100ms
	cb := NewCircuitBreaker(2, 100*time.Millisecond)

	// A function that fails
	flakyFunc := func() error {
		return errors.New("temporary failure")
	}

	// Trigger failures
	_ = cb.Execute(flakyFunc)
	_ = cb.Execute(flakyFunc)

	// Now the circuit is open
	err := cb.Execute(flakyFunc)
	if err == ErrCircuitOpen {
		fmt.Println("Circuit is open")
	}

	// Wait for reset timeout
	time.Sleep(110 * time.Millisecond)

	// Now it's in Half-Open state and will allow one call
	_ = cb.Execute(func() error {
		fmt.Println("Allowed one call in Half-Open state")
		return nil
	})

	// Output:
	// Circuit is open
	// Allowed one call in Half-Open state
}
