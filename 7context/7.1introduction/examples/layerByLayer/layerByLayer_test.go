package layerByLayer

import (
	"context"
	"fmt"
	"testing"
)

// Test_Check_layerByLayer is to observe the cancellation of the context layer by layer
func Test_Check_layerByLayer(t *testing.T) {
	var ctx = context.Background()
	var cancel context.CancelFunc

	// Create an inner layer context
	ctx = context.WithValue(ctx, "key", 0)
	ctx, cancel = context.WithCancel(ctx)
	defer func() {
		fmt.Println("cancel inner layer")
		cancel()
	}()

	// Create a middle layer context
	ctx = context.WithValue(ctx, "key", 1)
	ctx, cancel = context.WithCancel(ctx)
	defer func() {
		fmt.Println("cancel middle layer")
		cancel()
	}()

	// Create an outer layer context
	ctx = context.WithValue(ctx, "key", 2)
	ctx, cancel = context.WithCancel(ctx)
	defer func() {
		fmt.Println("cancel outer layer")
		cancel()
	}()
}
