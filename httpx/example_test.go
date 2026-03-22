package httpx_test

import (
	"fmt"
	"github.com/vikash-paf/goutils/httpx"
	"time"
)

func ExampleNewSafeClient() {
	cfg := httpx.ClientConfig{
		Timeout: 5 * time.Second,
	}
	client := httpx.NewSafeClient(cfg)
	fmt.Println(client.Timeout)
	// Output: 5s
}
