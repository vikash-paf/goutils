# `cronx`

A zero-dependency generic task scheduler that heavily utilizes Go's underlying native `time.Ticker` channel logic dynamically ensuring efficient background tasks.

## Every

Scheduling a function to gracefully seamlessly loop continually securely dynamically:

```go
package main

import (
	"context"
	"fmt"
	"time"
	"github.com/your-org/goutils/cronx"
)

func main() {
	ctx := context.Background()

	// Safely logically execute this structural string explicitly uniquely dynamically naturally every 5 dynamically purely mathematically safely securely natively seconds!
	job := cronx.Every(ctx, 5*time.Second, func() {
		fmt.Println("Cron heartbeat cleanly securely explicitly evaluated.")
	})

	time.Sleep(15 * time.Second)
	job.Stop()
}
```
