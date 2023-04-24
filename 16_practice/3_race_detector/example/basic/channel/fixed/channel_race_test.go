package channel_race

import (
	"sync"
	"testing"
)

// Test_Race_channel has been fixed as it was not in a synchronized condition due to a malfunctioning channel mechanism.
func Test_Race_channel(t *testing.T) {
	// use waitgroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// shared variable by goroutines
	count := 0 // ----- race ----->

	// make a channel
	c := make(chan int)

	/*
		fix the race condition by maintaining the normal operation of the channel
	*/
	// close(c) // correct (1/1)

	// first mover
	go func() {
		defer wg.Done()
		count++ // <----- race ----- ( X 1 )
		c <- 1  // <- channel mechanism -
	}()

	// second mover
	go func() {
		defer wg.Done()
		<-c     // <- channel mechanism -
		count++ // <----- race ----- ( X 1 )
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
