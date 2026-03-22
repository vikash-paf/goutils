# `httpx`

The `httpx` package provides a pre-configured HTTP client with sensible timeouts to prevent goroutine leaks.

## SafeHTTPClient

An HTTP client struct with mandatory timeouts and optimized connection pooling.
```go
package main

import (
	"fmt"
	"github.com/your-org/goutils/httpx"
)

func main() {
	client := httpx.NewSafeClient(httpx.DefaultClientConfig)
	
	// Fails with a timeout error if the request takes too long.
	_, err := client.Get("http://example.com/slow-endpoint")
	if err != nil {
		fmt.Println("Bounded error correctly:", err)
	}
}
```
