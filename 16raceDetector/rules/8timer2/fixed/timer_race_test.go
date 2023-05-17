package timer_race

import (
	"sync"
	"testing"
	"time"
)

// sync.Mutex is used to synchronize access to count.
var mu sync.Mutex // add (1/5) !

// Test_Race_timer fixes that channel are not synchronized.
func Test_Race_timer2(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Create timer
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop() // timer.Stop() makes timer channel closed

	// Shared variable by goroutines
	var count int // ----- race ----->

	// Start 1 goroutine to write to count
	go func() {
		defer wg.Done()
		mu.Lock()   // add (2/5) !
		count++     // <----- race -----
		mu.Unlock() // add (3/5) !
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
			mu.Lock()   // add (4/5) !
			count++     // <----- race -----
			mu.Unlock() // add (5/5) !
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
