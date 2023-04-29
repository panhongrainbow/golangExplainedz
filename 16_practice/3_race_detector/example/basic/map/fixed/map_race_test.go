package map_race

import (
	"math/rand"
	"sync"
	"testing"
)

var mu sync.Mutex // correct (1/3) !

// Test_Race_map fixed that map are not synchronized
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
			mu.Lock()      // correct (2/3)
			count["key"]++ // <----- race ----- ( X many )
			mu.Unlock()    // correct (3/3)
		}()
	}

	// Wait for all map to finish
	wg.Wait()
}

// Benchmark_Race_map test
func Benchmark_Race_map(b *testing.B) {
	// shared variable by map
	var count = make(map[int]int) // ----- race ----->

	for i := 0; i < b.N; i++ {
		mu.Lock()
		count[rand.Intn(101)]++ // <----- race ----- ( X many )
		mu.Unlock()
	}
}
