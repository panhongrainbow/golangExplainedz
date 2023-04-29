package sync_race

import (
	"sync"
	"testing"
)

// Test_Race_channel shows that channel is not in synchronized condition due to a malfunctioning sync mechanism.
func Test_Race_sync(t *testing.T) {
	// use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// shared variable by goroutines
	count := 0 // ----- race ----->

	/*
		Disable the sync mechanism here will cause a race condition,
		which is my purpose of doing so
		(我故意的)
	*/
	mu := sync.Mutex{}
	mu2 := sync.Mutex{}

	// make two goroutines
	go func() {
		defer wg.Done()
		mu.Lock()   // <- sync mechanism -
		count++     // <----- race ----- ( X 1 )
		mu.Unlock() // <- sync mechanism -
	}()
	go func() {
		defer wg.Done()
		mu2.Lock()   // <- sync mechanism -
		count++      // <----- race ----- ( X 1 )
		mu2.Unlock() // <- sync mechanism -
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
