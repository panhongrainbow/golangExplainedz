package channel_race

import (
	"sync"
	"testing"
)

// Test_Race_channel shows that channel is not in synchronized condition due to a malfunctioning channel mechanism.
func Test_Race_channel(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Shared variable by goroutines
	count := 0 // ----- race ----->

	// Make a channel
	c := make(chan int)

	/*
		CLose the channel and make the order of the two goroutines unpredictable
		Disable the channel mechanism here will cause a race condition,
		which is my purpose of doing so
		(我故意的)
	*/
	close(c) // <- race -

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
