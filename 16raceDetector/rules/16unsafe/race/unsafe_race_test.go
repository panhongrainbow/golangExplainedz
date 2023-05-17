package list_race

import (
	"sync"
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

// Test_Race_unsafe shows that the program is not in synchronized condition due to a malfunctioning unsafe pointer.
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

	// Start 1000 goroutines
	for i := 0; i < 1000; i++ { // <- race -
		i := i
		go func() {
			defer wg.Done()
			// Call correct() function
			correct(p2, i) // <----- race -----

		}()
	}

	// Get the value of the element pointed to by the pointer.
	// fmt.Println(numbers)

	// Wait for all goroutines to finish
	wg.Wait()
}
