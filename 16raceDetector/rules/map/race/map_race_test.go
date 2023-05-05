package map_race

import (
	"sync"
	"testing"
)

// Test_Race_map shows that map are not synchronized
func Test_Race_map(t *testing.T) {
	// use wait group to wait for all map to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// shared map
	var count = make(map[string]int) // ----- race ----->

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			count["key"]++ // <----- race ----- ( X many )
		}()
	}

	// Wait for all map to finish
	wg.Wait()
}
