package interface_race

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Test_Race_interface fixed that interface are not synchronized
func Test_Race_interface(t *testing.T) {
	// use waitgroup to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Create a channel
	c := make(chan int, 1000)

	// start 1000 goroutines
	for i := 0; i < 500; i++ {
		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			wg.Done()
			Write()
			c <- 1
		}()

		go func() {
			// Waiting
			<-c
			wg.Done()
			Set()
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Test_Race_interface fixed that interface are in synchronized
func Test_fixed_interface(t *testing.T) {
	// use waitgroup to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Create a channel
	c := make(chan int, 1)

	for i := 0; i < 500; i++ {
		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			wg.Done()
			MutexWrite()
			c <- 1
		}()

		go func() {
			// Waiting
			<-c
			wg.Done()
			MutexSet()
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Test_atomic_interface fixed that interface are in synchronized
func Test_atomic_interface(t *testing.T) {
	// use waitgroup to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Create a channel
	c := make(chan int, 1)

	for i := 0; i < 500; i++ {
		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			wg.Done()
			AtomicWrite()
			c <- 1
		}()

		go func() {
			// Waiting
			<-c
			wg.Done()
			AtomicSet()
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Benchmark_Race_interface test
func Benchmark_Race_interface(b *testing.B) {
	// shared variable by interface
	var count int32 = 0 // ----- race ----->

	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&count, 1) // <----- race ----- ( X many )
	}
}
