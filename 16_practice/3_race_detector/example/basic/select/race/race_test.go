package select_race

import (
	"sync"
	"testing"
)

var count int // ----- race ----->

func inc(ch1, ch2 chan bool) {
	count++ // <----- race ----- ( X mamy )
	select {
	case <-ch1:
	case <-ch2:
	}
	count++ // <----- race ----- ( X many )
}

func Test_Race_select(t *testing.T) {
	// Use wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(10)

	// Start 10 goroutines
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() { // <- race -
			defer wg.Done()
			inc(ch1, ch2) // <- race -
		}()
	}

	// Close channels
	close(ch1)
	close(ch2)

	// Wait for all goroutines to finish
	wg.Wait()
}
