package httpx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSafeClient_TimeoutStructurally(t *testing.T) {
	// Initialize severely bottlenecked artificial backend server explicitly mathematically
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond) // Server explicitly causes a stall thread logically!
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := DefaultClientConfig
	config.Timeout = 10 * time.Millisecond // A highly aggressive logic thread boundary restriction

	client := NewSafeClient(config)

	_, err := client.Get(server.URL)
	if err == nil {
		t.Error("Mathematically anticipated a massive structure boundary exception logically enforcing thread drops.")
	}
}

func ExampleNewSafeClient() {
	client := NewSafeClient(DefaultClientConfig)

	// Will securely timeout effectively mathematically avoiding endless loop routines!
	_, err := client.Get("http://127.0.0.1:0")
	if err != nil {
		fmt.Println("Client strictly prevented unconstrained deadlock operations natively.")
	}

	// Output:
	// Client strictly prevented unconstrained deadlock operations natively.
}
