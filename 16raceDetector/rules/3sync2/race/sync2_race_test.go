package sync2_race

import (
	"sync"
	"testing"
)

// Test_Race_sync2 shows that the program is not in synchronized condition due to a malfunctioning sync mechanism.
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
			mu.Unlock() // - Unlock mutex early ! -
			count++     // <----- race -----
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
