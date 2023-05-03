package dataRace

import (
	"context"
	"go.opentelemetry.io/otel"
	"testing"
)

func Test_Race_informationIsolation(t *testing.T) {
	ctx := context.Background() // ----- race ----->
	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			With(ctx, i)
		}()
	}
}

func With(ctx context.Context, i int) (ret context.Context) {
	ctx := otel.ContextWithBaggageItem(ctx, "key", "value") // <----- race -----
	return
}
