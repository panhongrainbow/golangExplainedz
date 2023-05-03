package informationBarrier

import (
	"context"
	"fmt"
	"testing"
)

// Create a channel to synchronize the goroutines
var c = make(chan int)

// Test_Check_informationIsolation is a test function about information barrier
func Test_Check_informationIsolation(t *testing.T) {
	// Create a context with a key-value pair
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	ctx, cancel := context.WithCancel(ctx)

	// Start a goroutine
	go messageReceiver(ctx)

	// Cancel the parent context
	cancel()

	// Tell the messageReceiver to continue
	c <- 1

	// Check the value of the key
	value := ctx.Value("key")
	fmt.Println("value in main function:", value)
}

// messageReceiver is a goroutine that is started from the main function
func messageReceiver(ctx context.Context) {
	// Wait for parent context to be canceled
	<-c
	// Check the value of the key
	value := ctx.Value("key")
	fmt.Println("value in messageReceiver :", value)
}
