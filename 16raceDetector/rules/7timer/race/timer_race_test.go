package timer_race

import (
	"sync"
	"testing"
	"time"
)

/*
Test_Race_timer shows that channel are not synchronized.
However, race detector does not detect it, probably.
*/
func Test_Race_timer(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Shared variable by goroutines
	var count int // ----- race ----->

	// Start 1 goroutine to write to count
	go func() {
		defer wg.Done()
		count++ // <----- race -----
	}()

	// Start 1 goroutine to write to count with timer
	select {
	case <-time.After(1 * time.Second): // <- random race -
		go func() {
			defer wg.Done()
			count++ // <----- race -----
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
