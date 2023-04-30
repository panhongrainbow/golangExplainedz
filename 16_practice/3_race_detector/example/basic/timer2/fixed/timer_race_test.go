package timer_race

import (
	"sync"
	"testing"
	"time"
)

// sync.Mutex is used to synchronize access to count
var mu sync.Mutex // add (1/5) !

// Test_Race_timer fixed that channel are not synchronized
func Test_Race_timer(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Create timer
	timer := time.NewTimer(1 * time.Second)

	// Shared variable by goroutines
	var count int // ----- race ----->

	// Start 1 goroutine to write to count
	go func() {
		defer wg.Done()
		timer.Stop() // timer.Stop() make timer channel closed // <- race -
		mu.Lock()    // add (2/5) !
		count++      // <----- race -----
		mu.Unlock()  // add (3/5) !
	}()

	// Start 1 goroutine to write to count with timer
	go func() {
		defer wg.Done()
		select {
		case <-timer.C: // <- random race -
		// do nothing
		default:
			mu.Lock() // add (4/5) !
			count++
			mu.Unlock() // add (5/5) !
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
