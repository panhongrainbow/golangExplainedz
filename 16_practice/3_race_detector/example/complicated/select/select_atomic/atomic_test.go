package select_atomic

import (
	"sync/atomic"
	"testing"
	"time"
)

var count int32 // ----- race ----->

func inc(ch1, ch2 chan bool) {
	atomic.AddInt32(&count, 1) // <----- race ----- ( X mamy )
	select {
	case <-ch1:
	case <-ch2:
	}
	atomic.AddInt32(&count, 1) // <----- race ----- ( X many )
}

func Test_Check(t *testing.T) {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i <= 10; i++ {
		go inc(ch1, ch2) // <- race -
		go inc(ch1, ch2) // <- race -
	}
	atomic.AddInt32(&count, 1) // <----- race ----- ( X 1 )

	time.Sleep(500 * time.Microsecond)
	close(ch1)

	close(ch2)
	<-ch2
}
