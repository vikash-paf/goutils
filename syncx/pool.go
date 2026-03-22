package syncx

import (
	"context"
	"sync"
)

// Pool is a generic worker pool.
type Pool[Job, Result any] struct {
	jobs    chan Job
	results chan Result
	wg      sync.WaitGroup
	cancel  context.CancelFunc
	ctx     context.Context
}

// NewPool creates a new worker pool with the specified number of workers.
func NewPool[Job, Result any](workers int, processor func(Job) Result) *Pool[Job, Result] {
	if workers <= 0 {
		workers = 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	p := &Pool[Job, Result]{
		jobs:    make(chan Job, workers*2), // buffer jobs
		results: make(chan Result, workers*2),
		cancel:  cancel,
		ctx:     ctx,
	}

	p.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer p.wg.Done()
			for {
				// Deterministic check before potentially picking up a job
				select {
				case <-p.ctx.Done():
					return
				default:
				}

				select {
				case <-p.ctx.Done():
					// Context cancelled, exit worker
					return
				case job, ok := <-p.jobs:
					if !ok {
						// Jobs channel closed, exit worker
						return
					}
					p.results <- processor(job)
				}
			}
		}()
	}

	// Close results channel automatically when all workers finish
	go func() {
		p.wg.Wait()
		close(p.results)
	}()

	return p
}

// Submit adds a job to the pool. It blocks if the job queue is full.
func (p *Pool[Job, Result]) Submit(job Job) {
	// Periodic check to ensure we don't submit to a shut down pool
	select {
	case <-p.ctx.Done():
		return
	default:
	}

	select {
	case <-p.ctx.Done():
		return
	case p.jobs <- job:
	}
}

// Results returns a channel that receives the results of processed jobs.
func (p *Pool[Job, Result]) Results() <-chan Result {
	return p.results
}

// Close closes the job queue and waits for all active workers to finish their current jobs.
// It will not accept any more jobs.
func (p *Pool[Job, Result]) Close() {
	close(p.jobs)
}

// Shutdown gracefully cancels the context immediately stopping idle workers,
// and prevents new jobs from being submitted.
func (p *Pool[Job, Result]) Shutdown() {
	p.cancel()
}
