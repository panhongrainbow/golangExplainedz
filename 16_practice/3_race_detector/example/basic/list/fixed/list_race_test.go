package list_race

import (
	"sync"
	"testing"
)

var mu sync.Mutex // add (1/4) !

// List is a linked list
type List struct {
	value int
	next  *List
}

// Test_Race_list fixed that goroutines are not synchronized
func Test_Race_list(t *testing.T) {
	// use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// shared variable by goroutines
	root := &List{value: -1} // ----- race ----->

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ { // <----- race ----- ( X many )
		i := i // add (2/4) !
		go func() {
			defer wg.Done()
			list := &List{value: i} // <----- race ----- ( X many )
			list.next = new(List)
			mu.Lock()        // add (3/4) !
			root.next = list // <----- race ----- ( X many )
			mu.Unlock()      // add (4/4) !
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Benchmark_Race_list test
func Benchmark_Race_list(b *testing.B) {
	// shared variable by goroutines
	root := &List{value: -1}

	// Reset timer
	b.ResetTimer()

	// Benchmark
	for i := 0; i < b.N; i++ {
		list := &List{value: i} // <----- race ----- ( X many )
		list.next = new(List)
		mu.Lock()
		root.next = list // <----- race ----- ( X many )
		mu.Unlock()
	}
}
