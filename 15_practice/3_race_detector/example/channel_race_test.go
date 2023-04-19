package example

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"sync"
	"testing"
)

func Test_Check_data_race(t *testing.T) {
	t.Run("global_variable_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)

		count := 0

		go func() {
			defer wg.Done()
			count++
		}()

		go func() {
			defer wg.Done()
			count++
		}()

		wg.Wait()
		require.Equal(t, 2, count)
	})
	t.Run("global_variable_no_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		var mu sync.Mutex

		count := 0

		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()

		wg.Wait()
		fmt.Println(count)
	})
	t.Run("channel_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		c := make(chan int)
		var count int
		close(c)

		read := func() {
			defer wg.Done()
			<-c
			count = count + 4
			fmt.Println("read!")
		}

		write := func() {
			defer wg.Done()
			count = count + 2
			fmt.Println("write!")
			c <- 1
		}

		go read()
		go write()
		wg.Wait()
	})
	t.Run("channel_no_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		c := make(chan int)
		var count int
		var mu sync.Mutex

		read := func() {
			defer wg.Done()
			<-c
			if count != 2 {
				log.Fatal("should write first, read later")
			}
			mu.Lock()
			count = count + 4
			mu.Unlock()
			fmt.Println("read!")
		}

		write := func() {
			defer wg.Done()
			mu.Lock()
			count = count + 2
			mu.Unlock()
			fmt.Println("write!")
			c <- 1
		}

		go read()
		go write()
		wg.Wait()
	})
	t.Run("sync_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		var count int
		var mu sync.Mutex  // using different sync Mutex
		var mu1 sync.Mutex // using different sync Mutex

		inc := func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}

		dec := func() {
			defer wg.Done()
			mu1.Lock()
			count--
			mu1.Unlock()
		}

		go inc()
		go dec()
		wg.Wait()
	})
	t.Run("sync_no_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		var count int
		var mu sync.Mutex

		inc := func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}

		dec := func() {
			defer wg.Done()
			mu.Lock()
			count--
			mu.Unlock()
		}

		inc()
		dec()
		wg.Wait()
	})
	t.Run("closure_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(4)
		var count int

		makeCounter := func() func() {
			wg.Done()
			var mu sync.Mutex // use different locks
			return func() {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}

		inc1 := makeCounter() // use different locks
		inc2 := makeCounter() // use different locks

		go inc1()
		go inc1()
		go inc2()
		go inc2()

		wg.Done()
	})
	t.Run("closure_no_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(4)
		var count int

		var mu sync.Mutex // use the same lock
		makeCounter := func() func() {
			wg.Done()
			return func() {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}

		inc1 := makeCounter()
		inc2 := makeCounter()

		go inc1()
		go inc1()
		go inc2()
		go inc2()

		wg.Done()
	})
}
