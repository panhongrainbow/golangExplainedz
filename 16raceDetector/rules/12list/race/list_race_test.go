package list_race

import (
	"fmt"
	"sync"
	"testing"
)

// List is a linked list
type List struct {
	value int
	next  *List
}

// Test_Race_list shows the root list is not synchronized.
func Test_Race_list(t *testing.T) {
	// use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// shared variable by goroutines
	root := &List{value: -1} // ----- race ----->

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ { // <- race -
		i := i
		go func() {
			// append to the list tail
			defer wg.Done()
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
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// count list length
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
