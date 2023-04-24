package closure_race

import (
	"sync"
	"testing"
)

// Test_Race_closure shows that closure is not in synchronized condition due to a malfunctioning closure mechanism..
func Test_Race_closure(t *testing.T) {
	// Use waitgroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by goroutines
	count := 0 // ----- race ----->

	// Make a closure
	closure := func() func() {
		// Use closure's lock
		var mu sync.Mutex
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
