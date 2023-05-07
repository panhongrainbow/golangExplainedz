package timer_race

import (
	"sync"
	"testing"
	"time"
)

/*
Test_Race_timer2 shows that channel are not synchronized.
However, race detector does not detect it, probably.
*/
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
		count++ // <----- race -----
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
			count++
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
