package cronx

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestCron_TimingSafely(t *testing.T) {
	ctx := context.Background()
	var counter int32

	job := Every(ctx, 10*time.Millisecond, func() {
		atomic.AddInt32(&counter, 1)
	})

	time.Sleep(35 * time.Millisecond)
	job.Stop()
	time.Sleep(15 * time.Millisecond) // Ensure the job logic terminates predictably dynamically!

	final := atomic.LoadInt32(&counter)
	if final < 2 || final > 4 {
		t.Errorf("expected etween 2 and 4: %d", final)
	}
}
