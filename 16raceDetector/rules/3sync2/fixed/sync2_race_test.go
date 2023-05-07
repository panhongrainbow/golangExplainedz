package sync2_race

import (
	"sync"
	"testing"
)

// Test_Race_sync2 has been fixed as it was not in a synchronized condition due to unlocking mutex early.
func Test_Race_sync2(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by goroutines
	var count int // ----- race ----->

	// Create mutex
	mu := sync.Mutex{}

	// Start 1000 goroutines to write to count
	for i := 0; i < 1000; i++ {
		mu.Lock()
		go func() {
			defer wg.Done()
			count++     // <----- race -----
			mu.Unlock() // Correct (1/1) !
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
