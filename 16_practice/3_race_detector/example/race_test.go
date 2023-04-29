package example

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Check_data_race(t *testing.T) {
	// As long as the program `accesses pointer variables`
	t.Run("pointer_race", func(t *testing.T) {
		// Setting up the wait group
		var wg sync.WaitGroup
		wg.Add(2)

		// Goroutines need to modify the pointer variable together
		var ptr *int // >>>>> >>>>> >>>>>

		// Defining the write function
		write := func() {
			defer wg.Done()
			x := 10
			ptr = &x // Modifying the pointer variable // <<<<< <<<<< <<<<<
		}

		// Defining the increment function
		incr := func() {
			defer wg.Done()
			*ptr++ // Accessing the variable pointed to by the pointer // <<<<< <<<<< <<<<<
		}

		// Simulating the main program
		go write()
		go incr()

		// Waiting
		wg.Wait()
	})
	t.Run("pointer_no_race", func(t *testing.T) {
		// Setting up the wait group
		var wg sync.WaitGroup
		wg.Add(2)
		c := make(chan int)

		// Using a mutex to protect the pointer variable
		var mu sync.Mutex

		// Defining the write function
		write := func(ptr *int) {
			defer wg.Done()
			mu.Lock()   // add lock
			*ptr = 10   // 修改指针变量
			mu.Unlock() // add unlock
			c <- 1
		}

		// Defining the increment function
		incr := func(ptr *int) {
			defer wg.Done()
			mu.Lock()   // add lock
			*ptr++      // access the variable pointed to by the pointer
			mu.Unlock() // add unlock
		}

		// var ptr1 = new(int)
		var ptr1 int

		// Simulating the main program
		go func() {
			write(&ptr1)
		}()
		<-c
		go incr(&ptr1)

		// Waiting
		wg.Wait()

		// print
		mu.Lock()
		fmt.Println(ptr1)
		mu.Unlock()
	})
	// As long as the program `accesses pointer variables`
	t.Run("interface_race", func(t *testing.T) {
		// Setting up the wait group
		var wg sync.WaitGroup
		wg.Add(2)

		// Create a channel
		c := make(chan int, 1)

		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			defer wg.Done()
			write()
			c <- 1
		}()

		go func() {
			// Waiting
			defer wg.Done()
			read()
			<-c
		}()

		// Waiting
		wg.Wait()
	})
	t.Run("interface_no_race", func(t *testing.T) {
		// Setting up the wait group
		var wg sync.WaitGroup
		wg.Add(2)

		// Create a channel
		c := make(chan int, 1)

		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			defer wg.Done()
			writeMu()
			c <- 1
		}()

		go func() {
			// Waiting
			defer wg.Done()
			readMu()
			<-c
		}()

		// Waiting
		wg.Wait()
	})
	t.Run("map_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)

		var m = make(map[string]int)
		m["key"] = 0

		write := func() {
			m["key"] = 10
		}

		read := func() {
			value := m["key"]
			fmt.Println(value)
		}

		go func() {
			defer wg.Done()
			write()
		}()
		go func() {
			defer wg.Done()
			read()
		}()

		wg.Wait()
	})
	t.Run("map_no_race", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		c := make(chan int, 1)

		var m = make(map[string]int)

		mu := sync.Mutex{}

		write := func() {
			mu.Lock()
			m["key"] = 10
			mu.Unlock()
			c <- 1
		}

		read := func() {
			<-c
			mu.Lock()
			value := m["key"]
			mu.Unlock()
			fmt.Println(value)
		}

		go func() {
			defer wg.Done()
			write()
		}()
		go func() {
			defer wg.Done()
			read()
		}()

		wg.Wait()
	})
	t.Run("map_sync", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)

		var m = sync.Map{}

		write := func() {
			m.Store("key", 10)
		}

		read := func() {
			value, _ := m.Load("key")
			fmt.Println(value)
		}

		go func() {
			defer wg.Done()
			write()
		}()
		go func() {
			defer wg.Done()
			read()
		}()

		wg.Wait()
	})
	t.Run("select_race", func(t *testing.T) {
		var count int

		inc := func(ch1, ch2 chan bool) {
			count++
			select {
			case <-ch1:
			case <-ch2:
			}
			count++ // 这里可能与其他 goroutine 并发访问 count
		}

		ch1 := make(chan bool)
		ch2 := make(chan bool)
		go inc(ch1, ch2)
		go inc(ch1, ch2)
		count++ // 主 goroutine 也访问 count

		// 1 秒后关闭一个 channel
		time.Sleep(1 * time.Second)
		close(ch1)

		// 读取另一个 channel 的值
		<-ch2
	})
	t.Run("select_no_race", func(t *testing.T) {
		var count int

		mu := sync.Mutex{}

		inc := func(ch1, ch2 chan bool) {
			mu.Lock()
			count++
			mu.Unlock()
			select {
			case <-ch1:
			case <-ch2:
			}
			mu.Lock()
			count++ // 这里可能与其他 goroutine 并发访问 count
			mu.Unlock()
		}

		ch1 := make(chan bool)
		ch2 := make(chan bool)
		go inc(ch1, ch2)
		go inc(ch1, ch2)
		count++ // 主 goroutine 也访问 count

		// 1 秒后关闭一个 channel
		time.Sleep(1 * time.Second)
		close(ch1)

		// 读取另一个 channel 的值
		<-ch2
	})
}
