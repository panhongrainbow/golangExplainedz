package goroutine_race

import (
	"sync"
	"testing"
)

var mu sync.Mutex // correct (1/3) !

// Test_Race_goroutines fixed that goroutines are not synchronized
func Test_Race_goroutines(t *testing.T) {
	// use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// shared variable by goroutines
	count := 0 // ----- race ----->

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()   // correct (2/3)
			count++     // <----- race ----- ( X many )
			mu.Unlock() // correct (3/3)
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Benchmark_Race_goroutines test
func Benchmark_Race_goroutines(b *testing.B) {
	// shared variable by goroutines
	var count int32 = 0 // ----- race ----->

	for i := 0; i < b.N; i++ {
		mu.Lock()
		count++ // <----- race ----- ( X many )
		mu.Unlock()
	}
}
