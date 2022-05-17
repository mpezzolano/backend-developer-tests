// Package concurrency implements worker pool interfaces, one simple and one a
// bit more complex.
package concurrency

import "fmt"

// SimplePool is a simple worker pool that does not support cancellation or
// closing. All functions are safe to call from multiple goroutines.
type SimplePool interface {
	// Submit a task to be executed asynchronously. This function will return as
	// soon as the task is submitted. If the pool does not have an available slot
	// for the task, this blocks until it can submit.
	Submit(func())
}

// NewSimplePool creates a new SimplePool that only allows the given maximum
// concurrent tasks to run at any one time. maxConcurrent must be greater than
// zero.
func NewSimplePool(maxConcurrent int) SimplePool {
	var sp SimplePool
	jobs := make(chan func())
	sp = &JobPool{maxConcurrent, jobs}

	// Create workers and jobs channel
	for i := 0; i < maxConcurrent; i++ {
		go Worker(i, jobs)
	}

	return sp
}

// Generic worker with some added text to verify things are working
func Worker(id int, jobs <-chan func()) {
	for job := range jobs {
		fmt.Printf("worker %v starting sleep\n", id)
		job()
		fmt.Printf("worker %v awake\n", id)
	}
}

type JobPool struct {
	maxWorkers int
	jobs       chan func()
}

func (jp *JobPool) Submit(task func()) {
	jp.jobs <- task
}
