package channel_race

import (
	"sync"
	"testing"
)

// Test_Race_channel shows that channel is not in synchronized condition due to a malfunctioning channel mechanism..
func Test_Race_channel(t *testing.T) {
	// use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// shared variable by goroutines
	count := 0 // ----- race ----->

	// make a channel
	c := make(chan int)

	/*
		Disable the channel mechanism here will cause a race condition,
		which is my purpose of doing so
		(我故意的)
	*/
	close(c)

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
