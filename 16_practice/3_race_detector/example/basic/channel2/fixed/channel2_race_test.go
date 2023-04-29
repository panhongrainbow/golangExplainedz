package channel2_race

import (
	"fmt"
	"sync"
	"testing"
)

var mu sync.Mutex // correct (1/3) !

// Test_Race_channel2 fixed that channel are not synchronized
func Test_Race_channel2(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(20)

	// Shared variable by goroutines
	c := make(chan int) // ----- race ----->
	defer close(c)

	// Start 10 goroutines to write to channel
	for i := 0; i < 10; i++ {
		i := i // Correct (1/1) !
		go func() {
			defer wg.Done()
			fmt.Println(i)
			c <- i // <----- race -----
		}()
	}

	// Start 10 goroutines to read from channel
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			v := <-c
			fmt.Println("v", v)
			v++
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
