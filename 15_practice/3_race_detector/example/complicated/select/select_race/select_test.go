package select_race

import (
	"testing"
	"time"
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

func Test_Check(t *testing.T) {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	for i := 0; i <= 10; i++ {
		go inc(ch1, ch2) // <- race -
		go inc(ch1, ch2) // <- race -
	}
	count++ // <----- race ----- ( X 1 )

	time.Sleep(500 * time.Microsecond)
	close(ch1)

	close(ch2)
	<-ch2
}
