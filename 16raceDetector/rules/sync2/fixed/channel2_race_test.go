package sync2_race

import (
	"sync"
	"testing"
)

// Test_Race_sync2 fixed that channel are not synchronized
func Test_Race_sync2(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(10)

	// Shared variable by goroutines
	var count int // ----- race ----->

	// Create mutex
	mu := sync.Mutex{}

	// Start 10 goroutines to write to channel
	for i := 0; i < 10; i++ {

		go func() {
			mu.Lock()   // Correct (1/2) !
			count++     // <----- race -----
			mu.Unlock() // Correct (2/2) !
			defer wg.Done()
		}()

	}

	// Wait for all goroutines to finish
	wg.Wait()
}
