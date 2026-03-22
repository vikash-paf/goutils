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

// DefaultClientConfig represents safe values for production services.
var DefaultClientConfig = ClientConfig{
	Timeout:             10 * time.Second,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
	IdleConnTimeout:     90 * time.Second,
}

// NewSafeClient returns a pre-configured http.Client with specific timeouts and transport settings.
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
