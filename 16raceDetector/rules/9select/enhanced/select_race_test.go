package select_atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Shared variable by goroutines
var count int32 // ----- race ----->

// inc increments count by one.
func inc(ch1, ch2 chan bool) {
	atomic.AddInt32(&count, 1) // <----- race ----- ( X many ) // correct (1/2) !
	select {
	case <-ch1:
	case <-ch2:
	}
	atomic.AddInt32(&count, 1) // <----- race ----- ( X many )  // correct (2/2) !
}

// Test_Race_select has been fixed as it was not in a synchronized condition due to a malfunctioning select mechanism.
func Test_Race_select(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(10)

	// Start 10 goroutines
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() { // <- race -
			defer wg.Done()
			inc(ch1, ch2) // <- race -
		}()
	}

	// What if close channels accidentally // <- race -
	close(ch1)
	close(ch2)

	// Wait for all goroutines to finish
	wg.Wait()
}
