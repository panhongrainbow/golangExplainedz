package select_fixed

import (
	"sync"
	"testing"
	"time"
)

var mu sync.Mutex

var count int // ----- race ----->

func inc(ch1, ch2 chan bool) {
	mu.Lock()   // add !
	count++     // <----- race ----- ( X mamy )
	mu.Unlock() // add !
	select {
	case <-ch1:
	case <-ch2:
	}
	mu.Lock()   // add !
	count++     // <----- race ----- ( X many )
	mu.Unlock() // add !
}

func Test_Check(t *testing.T) {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i <= 10; i++ {
		go inc(ch1, ch2) // <- race -
		go inc(ch1, ch2) // <- race -
	}
	mu.Lock()   // add !
	count++     // <----- race ----- ( X 1 )
	mu.Unlock() // add !

	time.Sleep(500 * time.Microsecond)
	close(ch1)

	close(ch2)
	<-ch2
}
