package list_race

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// List is a linked list
type List struct {
	value int
	next  atomic.Pointer[List] // correct (1/3) !
}

// Test_Race_list fixes that the root list is not synchronized.
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
			list := List{value: i}
			next := root
			for {
				if next.next.Load() == nil { // correct (2/3) !
					if next.next.CompareAndSwap(nil, &list) { // <----- race ----- ( X many ) // correct (3/3) !
						break
					}
				} else {
					next = next.next.Load()
				}
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Count list length
	var count int
	next := root
	for {
		if next.next.Load() == nil {
			break
		} else {
			count++
			next = next.next.Load()
		}
	}
	fmt.Println("list length: ", count)

	// Check the detail carefully
	fmt.Println(root.value)
	fmt.Println(root.next.Load().value)
	fmt.Println(root.next.Load().next.Load().value)
	fmt.Println(root.next.Load().next.Load().next.Load().value)
	fmt.Println(root.next.Load().next.Load().next.Load().next.Load().value)
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
				list := List{value: i}
				next := root
				for {
					if next.next.Load() == nil {
						if next.next.CompareAndSwap(nil, &list) {
							break
						}
					} else {
						next = next.next.Load()
					}
				}
			}()
		}

		wg.Wait()
	}
}
