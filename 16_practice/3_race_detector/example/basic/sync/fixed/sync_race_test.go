package sync_race

import (
	"sync"
	"testing"
)

// Test_Race_sync has been fixed as it was not in a synchronized condition due to a malfunctioning sync mechanism.
func Test_Race_sync(t *testing.T) {
	// use waitgroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// shared variable by goroutines
	count := 0 // ----- race ----->

	/*
		fix the race condition by maintaining the normal operation of the sync mechanism
		(我故意的)
	*/
	mu := sync.Mutex{}
	// mu2 := sync.Mutex{} // correct (1/3)

	// make two goroutines
	go func() {
		defer wg.Done()
		mu.Lock()   // <- sync mechanism -
		count++     // <----- race ----- ( X 1 )
		mu.Unlock() // <- sync mechanism -
	}()
	go func() {
		defer wg.Done()
		mu.Lock()   // <- sync mechanism - // correct (2/3)
		count++     // <----- race ----- ( X 1 )
		mu.Unlock() // <- sync mechanism - // correct (3/3)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
