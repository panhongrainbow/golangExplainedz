package goroutine_race

import (
	"sync"
	"testing"
)

// Test_Race_goroutines shows that goroutines are not synchronized
func Test_Race_goroutines(t *testing.T) {
	// use waitgroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// shared variable by goroutines
	count := 0 // ----- race ----->

	// start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			count++ // <----- race ----- ( X many )
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
