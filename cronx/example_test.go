package cronx_test

import (
	"context"
	"fmt"
	"time"
	"github.com/vikash-paf/goutils/cronx"
)

func ExampleEvery() {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	// Execute a task every 100ms.
	job := cronx.Every(ctx, 100*time.Millisecond, func() {
		fmt.Println("Cron heartbeat executed.")
	})

	<-ctx.Done()
	job.Stop()
	// Output: Cron heartbeat executed.
}
