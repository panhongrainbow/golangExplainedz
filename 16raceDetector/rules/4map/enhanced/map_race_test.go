package map_race

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

// Test_Race_map fixes that map is not synchronized
func Test_Race_map(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by map
	var count = sync.Map{}         // correct (1/4) !
	count.Store("key", new(int64)) // correct (2/4) !

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			value, _ := count.Load("key")      // correct (3/4) !
			atomic.AddInt64(value.(*int64), 1) // correct (4/4) !
		}()
	}

	// Wait for all map to finish
	wg.Wait()

	// Check the result sum
	var sum int64
	count.Range(func(key, value interface{}) bool {
		sum = sum + *value.(*int64)
		return true
	})
	fmt.Println("result of sync map is ", sum)
}

// Benchmark_Race_map tests the performance of the fixed map.
func Benchmark_Race_map(b *testing.B) {
	// Shared variable by map
	var count = sync.Map{} // ----- race ----->

	// Start goroutines to write to the sync map
	for i := 0; i < b.N; i++ {
		go func() {
			if value, ok := count.Load(rand.Intn(101)); ok { // <- sync map -
				atomic.AddInt64(value.(*int64), 1) // <----- race ----- ( X many )
			} else {
				count.Store(rand.Intn(101), new(int64)) // <----- race ----- ( X many )
			}
		}()
	}
}
