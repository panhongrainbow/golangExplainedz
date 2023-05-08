package list_race

import (
	"sync"
	"testing"
)

// List is a linked list
type List struct {
	value int
	next  *List
}

// Test_Race_list shows that goroutines are not synchronized
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
			defer wg.Done()
			list := &List{value: i} // <----- race ----- ( X many )
			list.next = new(List)
			root.next = list // <----- race ----- ( X many )
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
