package httpx

import (
	"net/http"
	"time"
)

// ClientConfig establishes foundational settings for constructing a robust API client connection.
type ClientConfig struct {
	Timeout             time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

// DefaultClientConfig represents safe and aggressive values built natively for typical 
// REST JSON implementations, rigorously eliminating long-term thread wait leaks structurally.
var DefaultClientConfig = ClientConfig{
	Timeout:             10 * time.Second,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
	IdleConnTimeout:     90 * time.Second,
}

// NewSafeClient orchestrates a natively fully-secured core HTTP execution thread structurally optimized 
// mathematically strictly eliminating infinite thread connection bugs historically typical across generic platforms.
func NewSafeClient(cfg ClientConfig) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        cfg.MaxIdleConns,
		MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
		IdleConnTimeout:     cfg.IdleConnTimeout,
	}

	return &http.Client{
		Timeout:   cfg.Timeout,
		Transport: transport,
	}
}
