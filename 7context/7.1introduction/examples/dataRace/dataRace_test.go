package dataRace

import (
	"context"
	"testing"
)

func Test_Race_informationIsolation(t *testing.T) {
	ctx := context.Background() // ----- race ----->
	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			ctx = context.WithValue(ctx, "key", i) // <----- race -----
		}()
	}
}
