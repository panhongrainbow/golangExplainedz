package list_race

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

/*
correct dereferences pointer to int, assigns value, returns; data race prone !
functions created by unsafe point are really afraid of data race.
*/
func correct(p unsafe.Pointer, i int) { // <- race -
	*(*int)(p) = i // ----- race ----->
	return
}

// Test_Race_unsafe fixes that unsafe mechanism is not synchronized.
func Test_Race_unsafe(t *testing.T) {
	// Create a slice of integers
	numbers := []int{1, 0}

	// Get a pointer to the first element of the slice
	p1 := unsafe.Pointer(&(numbers[0]))

	// Get a pointer to the second element of the slice
	p2 := unsafe.Pointer(uintptr(p1) + unsafe.Sizeof(numbers[0]))

	// use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Create a CAS Lock
	var locked int32 = 0 // add (1/3) !

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ { // <- race -
		i := i
		go func() {
			defer wg.Done()
			// Call correct() function
			if atomic.CompareAndSwapInt32(&locked, 0, 1) { // add (2/3) !
				correct(p2, i)                // <----- race -----
				atomic.StoreInt32(&locked, 0) // add (3/3) !
			}
		}()
	}

	// Get the value of the element pointed to by the pointer.
	// fmt.Println(numbers)

	// Wait for all goroutines to finish
	wg.Wait()
}

// Benchmark_Race_unsafe tests the performance of unsafe mechanism.
func Benchmark_Race_unsafe(b *testing.B) {
	// Create a CAS Lock
	var locked int32 = 0 // add (1/3) !

	// Shared variable by goroutines
	numbers := []int{1, 0}

	// Get a pointer to the first element of the slice
	p1 := unsafe.Pointer(&(numbers[0]))

	// Get a pointer to the second element of the slice
	p2 := unsafe.Pointer(uintptr(p1) + unsafe.Sizeof(numbers[0]))

	// Write to the shared variable
	for i := 0; i < b.N; i++ {
		// Call correct() function
		if atomic.CompareAndSwapInt32(&locked, 0, 1) { // add (2/3) !
			correct(p2, 2)
			atomic.StoreInt32(&locked, 0) // add (3/3) !
		}
	}
}
