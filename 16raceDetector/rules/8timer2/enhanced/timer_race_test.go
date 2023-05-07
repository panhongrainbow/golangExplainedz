package timer_race

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Test_Race_timer fixed that channel are not synchronized.
func Test_Race_timer2(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Create timer
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop() // timer.Stop() makes timer channel closed

	// Shared variable by goroutines
	var count int32 // ----- race ----->

	// Start 1 goroutine to write to count
	go func() {
		defer wg.Done()
		atomic.AddInt32(&count, 1) // <----- race ----- // correct (1/2) !
	}()

	// Start 1 goroutine to write to count with timer
	go func() {
		defer wg.Done()
		select {
		case <-timer.C: // <- random race -
			// if timer is setted to 0 * time.Second
			// do nothing
		default:
			// if timer is setted to 1 * time.Second
			atomic.AddInt32(&count, 1) // <----- race ----- // correct (2/2) !
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
