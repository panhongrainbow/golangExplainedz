package channel_race

import (
	"sync"
	"testing"
)

// Test_Race_channel has been fixed as it was not in a synchronized condition due to a malfunctioning channel mechanism.
func Test_Race_channel(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Shared variable by goroutines
	count := 0 // ----- race ----->

	// Make a channel
	c := make(chan int)

	/*
		Fix the race condition by maintaining the normal operation of the channel
	*/
	// Close(c) // correct (1/1)

	// First mover
	go func() {
		defer wg.Done()
		count++ // <----- race ----- ( X 1 )
		c <- 1  // <- channel mechanism -
	}()

	// Second mover
	go func() {
		defer wg.Done()
		<-c     // <- channel mechanism -
		count++ // <----- race ----- ( X 1 )
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
