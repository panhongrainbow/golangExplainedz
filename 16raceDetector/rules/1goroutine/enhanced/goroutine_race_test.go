package goroutine_race

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Test_Race_goroutines fixes that goroutines are not synchronized
func Test_Race_goroutines(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by goroutines
	var count int32 = 0 // ----- race ----->

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // <----- race ----- ( X many ) // correct (1/1) !
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Benchmark_Race_goroutines tests the performance of goroutines.
func Benchmark_Race_goroutines(b *testing.B) {
	// Shared variable by goroutines
	var count int32 = 0

	// Write to the shared variable
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&count, 1)
	}
}
