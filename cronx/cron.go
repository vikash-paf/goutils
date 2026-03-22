package cronx

import (
	"context"
	"time"
)

// Job represents a background task that runs at a specific interval.
type Job struct {
	ticker *time.Ticker
	quit   chan struct{}
}

// Every starts a new background job that executes the given task at every interval.
func Every(ctx context.Context, interval time.Duration, task func()) *Job {
	job := &Job{
		ticker: time.NewTicker(interval),
		quit:   make(chan struct{}),
	}

	go func() {
		for {
			select {
			case <-job.ticker.C:
				task()
			case <-ctx.Done():
				job.ticker.Stop()
				return
			case <-job.quit:
				job.ticker.Stop()
				return
			}
		}
	}()

	return job
}

// Stop terminates the background job.
func (j *Job) Stop() {
	close(j.quit)
}
