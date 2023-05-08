package list_race

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

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
	root.next = new(List)    // Add (1/3) ! (root.next must be initialized to avoid CAS failure)

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ { // <- race -
		i := i // add ! (2/3)
		go func() {
			defer wg.Done()
			list := &List{value: i} // <----- race ----- ( X many )
			list.next = new(List)
			if atomic.CompareAndSwapPointer( // <----- race ----- ( X many ) // Correct (3/3) ! by using CAS
				(*unsafe.Pointer)(unsafe.Pointer(root.next)),
				unsafe.Pointer(root.next),
				unsafe.Pointer(list.next),
			) {
				return
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Benchmark_Race_list test
func Benchmark_Race_list(b *testing.B) {
	// shared variable by goroutines
	root := &List{value: -1} // ----- race ----->
	root.next = new(List)    // Add (1/3) ! (root.next must be initialized to avoid CAS failure)

	// Reset timer
	b.ResetTimer()

	// Benchmark
	for i := 0; i < b.N; i++ {
		list := &List{value: i} // <----- race ----- ( X many )
		list.next = new(List)

		// CAS
		atomic.CompareAndSwapPointer( // <----- race ----- ( X many ) // Correct (3/3) ! by using CAS
			(*unsafe.Pointer)(unsafe.Pointer(root.next)),
			unsafe.Pointer(root.next),
			unsafe.Pointer(list.next),
		)
	}
}
