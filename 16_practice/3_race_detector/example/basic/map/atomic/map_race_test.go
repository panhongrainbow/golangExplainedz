package map_race

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

// Test_Race_map fixed that map are not synchronized
func Test_Race_map(t *testing.T) {
	// use waitgroup to wait for all map to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// shared map
	var count = sync.Map{} // ----- race ----->
	count.Store("key", new(int64))

	// start 1000 goroutines
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			value, _ := count.Load("key")
			atomic.AddInt64(value.(*int64), 1) // <----- race ----- ( X many )
		}()
	}

	// Wait for all map to finish
	wg.Wait()

	// check the result sum
	/*var sum int64
	count.Range(func(key, value interface{}) bool {
		sum = sum + *value.(*int64)
		return true
	})
	fmt.Println(sum)*/
}

// Benchmark_Race_map test
func Benchmark_Race_map(b *testing.B) {
	// shared variable by map
	var count = sync.Map{} // ----- race ----->

	for i := 0; i < b.N; i++ {
		if value, ok := count.Load(rand.Intn(101)); ok { // <- sync map -
			atomic.AddInt64(value.(*int64), 1) // <----- race ----- ( X many )
		} else {
			count.Store(rand.Intn(101), new(int64)) // <----- race ----- ( X many )
		}
	}
}
