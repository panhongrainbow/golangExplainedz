package enhanced

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Test_Race_closure has been fixed as it was not in a synchronized condition due to a malfunctioning closure mechanism.
func Test_Race_closure(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by goroutines
	var count int32 // ----- race -----> // correct (1/2) !

	// Make a closure
	closure := func() func() {
		return func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // <----- race ----- ( X many ) // correct (2/2) !
		}
	}

	/*for i := 0; i < 1000; i++ {
		go closure() //  Not correct, it will result in being unable to unlock
	}*/

	// Create 1000 goroutines
	for i := 0; i < 1000; i++ {
		/*
			The return value should be stored in the fn variable to be garbage collected;
			Otherwise, it will result in being unable to unlock
		*/
		fn := closure()
		go func() {
			fn() // Call the fn variable instead of closure()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
