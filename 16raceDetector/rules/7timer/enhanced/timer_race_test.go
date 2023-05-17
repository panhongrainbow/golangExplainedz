package timer_race

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Test_Race_timer fixes that channel are not synchronized.
func Test_Race_timer(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Shared variable by goroutines
	var count int32 // ----- race ----->

	// Start 1 goroutine to write to count
	go func() {
		defer wg.Done()
		atomic.AddInt32(&count, 1) // <----- race ----- // correct (1/2) !
	}()

	// Start 1 goroutine to write to count with timer
	select {
	case <-time.After(1 * time.Second): // <- random race -
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // <----- race ----- // correct (2/2) !
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
