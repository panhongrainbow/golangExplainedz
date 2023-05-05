package timer_race

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Test_Race_timer fixed that channel are not synchronized
func Test_Race_timer(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Create timer
	timer := time.NewTimer(1 * time.Second)

	// Shared variable by goroutines
	var count int32 // ----- race ----->

	// Start 1 goroutine to write to count
	go func() {
		defer wg.Done()
		timer.Stop()               // timer.Stop() make timer channel closed // <- race -
		atomic.AddInt32(&count, 1) // <----- race ----- // correct (1/2) !
	}()

	// Start 1 goroutine to write to count with timer
	go func() {
		defer wg.Done()
		select {
		case <-timer.C: // <- random race -
		// do nothing
		default:
			atomic.AddInt32(&count, 1) // <----- race ----- // correct (1/2) !
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
