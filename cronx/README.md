# `cronx`
 
A zero-dependency task scheduler for Go, built on top of `time.Ticker`. It provides a simple way to run background tasks at regular intervals with support for context cancellation.
 
## Every
 
Use `Every` to schedule a function to run at a specific interval.
 
```go
package main
 
import (
	"context"
	"fmt"
	"time"
	"github.com/vikash-paf/goutils/cronx"
)
 
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
 
	// Execute a task every 5 seconds.
	job := cronx.Every(ctx, 5*time.Second, func() {
		fmt.Println("Cron heartbeat executed.")
	})
 
	<-ctx.Done()
	job.Stop()
}
```
