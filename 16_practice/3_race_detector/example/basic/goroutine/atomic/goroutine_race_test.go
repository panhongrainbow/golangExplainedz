package goroutine_race

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Test_Race_goroutines fixed that goroutines are not synchronized
func Test_Race_goroutines(t *testing.T) {
	// use waitgroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// shared variable by goroutines
	var count int32 = 0 // ----- race ----->

	// start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // <----- race ----- ( X many ) // correct (1/1) !
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
		atomic.AddInt32(&count, 1) // <----- race ----- ( X many )
	}
}
