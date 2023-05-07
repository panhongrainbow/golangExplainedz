package map_race

import (
	"math/rand"
	"sync"
	"testing"
)

var mu sync.Mutex // correct (1/3) !

// Test_Race_map fixed that map is not synchronized
func Test_Race_map(t *testing.T) {
	// Use wait group to wait for all goroutine to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by map
	var count = make(map[string]int) // ----- race ----->

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()      // correct (2/3) !
			count["key"]++ // <----- race ----- ( X many )
			mu.Unlock()    // correct (3/3) !
		}()
	}

	// Wait for all map to finish
	wg.Wait()
}

// Benchmark_Race_map tests the performance of map.
func Benchmark_Race_map(b *testing.B) {
	// Shared variable by map
	var count = make(map[int]int)

	// Start goroutines to write to map
	for i := 0; i < b.N; i++ {
		go func() {
			mu.Lock()
			count[rand.Intn(100)]++
			mu.Unlock()
		}()
	}
}
