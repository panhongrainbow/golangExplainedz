# Context

> It is called `context` because it can create `parent-child relationships`.

## Usage

There are some suggestions and limitations for using context in Go language:
1. Context should be passed from `parent Goroutine to child Goroutine`.
   The parent Goroutine creates context and passes it down when creating child Goroutine.

2. Do not store context in a structure and use it for a long time.
   Context is short-lived and should be canceled or timed out in time.
   If stored in a structure for a long time, it is easy to cause Goroutine leakage.
   Mainly to prevent other goroutines with `the same parent context` from `receiving unexpected signals`. 

3. When canceling a context, cancel all its subcontexts in time.
   If any subcontext leaks, it will cause Goroutine leakage. `Cancel them layer by layer`.

   ```go
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
   
   /*
   Output
   cancel outer layer
   cancel middle layer
   cancel inner layer
   
   You can see from defer that the cancel of context is canceled from the outer layer to the inner layer !
   */
   ```

4. Between the upper and lower ctx, there is `no` so-called `information barrier`,
   Here is an example

   ```go
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
   
   /*
   value in messageReceiver : value
   value in main function: value
   
   As you can see here, there is no way to prevent the lower ctx from getting information from the upper level
   */
   ```

5. In an `uninterruptible situation`, do `not` directly receive the context parameter, because there is `no message barrier`, and `immediately interrupt` when the upper-level function wants to interrupt.(如果没有讯息屏障，别人要中断就中断)
   Here's an example of a message barrier

   ```go
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
   
   /*
   Output
   clean resource
   clean resource
   clean resource
   
   As you can see here, the message barrier is used to isolate the message from the parent goroutine
   */
   ```

6. The message barrier may generate `goroutine leak`.
   Another condition is that `the interruption of memory cleaning` will also generate a `momory leak`.

7. Select the appropriate context type.

   |                      |                                                              |
   | -------------------- | ------------------------------------------------------------ |
   | context.Background() | return a non-nil, empty context that is never canceled, has no values, and has no deadline. It is typically used as `the top-level context for incoming requests`, in the main function, initialization, and tests |
   | TODO()               | returns a non-nil, empty context that is used when it is unclear which context to use or when the context is not yet available. It is a temporary context input that should be replaced as soon as it gets clearer which context should be used |
   | context.WithCancel   | used when a context needs to be cancelled                    |
   | WithDeadline         | the deadline is known in advance and can be specified as an absolute time |
   | WithTimeout          | implement timeouts and cancellations in network requests, database queries, and other long-running operations |

8. Context's Value-related methods should be used with caution. context's values are `interface types` and can `be modified at any time`. 

9. After the context has expired, `all references` to it will be released.
   ```go
   conn, err := net.Dial("tcp", "example.com:80")
   if err != nil {
       // ...
   }
   
   ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
   defer cancel()  
   
   go func() {
     <-ctx.Done() // After the context has expired
     conn.Close() // Release the connection
   }()
   
   // ... 
   ```

10. Some parts of the context are `not race safe`, here's an example
    ```go
    package dataRace
    
    import (
    	"context"
    	"testing"
    )
    
    func Test_Race_informationIsolation(t *testing.T) {
    	ctx := context.Background() // ----- race ----->
        i := i
    	for i := 0; i < 1000; i++ {
    		go func() {
    			ctx = context.WithValue(ctx, "key", i) // <----- race -----
    		}()
    	}
    }
    ```

    





