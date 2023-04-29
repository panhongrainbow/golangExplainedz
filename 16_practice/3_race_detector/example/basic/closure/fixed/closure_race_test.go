package closure_race

import (
	"sync"
	"testing"
)

// Test_Race_closure has been fixed as it was not in a synchronized condition due to a malfunctioning closure mechanism.
func Test_Race_closure(t *testing.T) {
	// use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by goroutines
	count := 0 // ----- race ----->

	// Use the shared lock
	var mu sync.Mutex // correct (1/1)

	// Make a closure
	closure := func() func() {
		return func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
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
