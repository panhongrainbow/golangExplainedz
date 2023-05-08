package list_race

import (
	"fmt"
	"sync"
	"testing"
)

// Create a mutex
var mu sync.Mutex // add (1/3) !

// List is a linked list
type List struct {
	value int
	next  *List
}

// Test_Race_list fixed that the root list is not synchronized.
func Test_Race_list(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Shared variable by goroutines
	root := &List{value: -1} // ----- race ----->

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ { // <- race -
		i := i
		go func() {
			// Append to the list tail
			defer wg.Done()
			mu.Lock() // add (2/3) !
			list := &List{value: i}
			next := root
			for {
				if next.next == nil {
					next.next = list // <----- race ----- ( X many )
					break
				} else {
					next = next.next
				}
			}
			mu.Unlock() // add (3/3) !
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Count list length
	var count int
	next := root
	for {
		if next.next == nil {
			break
		} else {
			count++
			next = next.next
		}
	}
	fmt.Println("list length: ", count)
}

// Benchmark_Race_list test
func Benchmark_Race_list(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(1000)
		root := &List{value: -1}

		for i := 0; i < 1000; i++ {
			i := i
			go func() {
				defer wg.Done()
				mu.Lock()
				list := &List{value: i}
				next := root
				for {
					if next.next == nil {
						next.next = list
						break
					} else {
						next = next.next
					}
				}
				mu.Unlock()
			}()
		}

		wg.Wait()
	}
}
