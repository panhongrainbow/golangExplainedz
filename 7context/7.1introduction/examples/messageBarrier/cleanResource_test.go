package messageBarrier

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Create a wait group
var wg = sync.WaitGroup{}

// Test_Check_cleanResource is to observe interupting the cleanResource
func Test_Check_cleanResource(t *testing.T) {
	// Wait for the clean function to finish
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	// Pass ctx to clean function
	go clean(ctx)

	// Because of information barrier, clean function won't be interrupted by this cancel function
	cancel()

	// Wait for the clean function to finish
	wg.Wait()
}

// clean is to represent a function that clean resource
func clean(ctx context.Context) {
	// When the clean function is finished, notify the wait group
	defer wg.Done()

	// <----- message barrier ----->
	// Remove the connection with the parent context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

LOOP:
	// The combination of for and select
	for {
		select {
		case <-ctx.Done(): // cancel parent ctx
			return
		default:
			// Clean resource for 3 times
			for i := 0; i < 3; i++ {
				fmt.Println("clean resource")
				time.Sleep(1 * time.Second)
			}
			break LOOP
		}
	}
}
