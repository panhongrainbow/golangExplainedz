package interface_race

import (
	"sync"
	"testing"
)

// Test_Race_interface fixed that interface are not synchronized
func Test_Race_interface(t *testing.T) {
	// use wait group to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Create a channel
	//c := make(chan int, 1000)
	/*
		Disable the channel mechanism here will cause a race condition,
		which is my purpose of doing so
		(我故意的)
	*/
	//close(c)

	// Start 1000 goroutines
	for i := 0; i < 500; i++ {
		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			defer wg.Done()
			Write()
			// c <- 1
		}()

		go func() {
			// Waiting
			// <-c
			defer wg.Done()
			Set()
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Test_Race_interface fixed that interface are in synchronized
func Test_fixed_interface(t *testing.T) {
	// use wait group to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Start 1000 goroutines
	for i := 0; i < 500; i++ {
		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			defer wg.Done()
			MutexWrite()
		}()

		go func() {
			// Waiting
			defer wg.Done()
			MutexSet()
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Test_atomic_interface fixed that interface are in synchronized
func Test_atomic_interface(t *testing.T) {
	// use wait group to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Start 1000 goroutines
	for i := 0; i < 500; i++ {
		// Goroutines need to modify the interface variable together
		go func() {
			// Waiting
			defer wg.Done()
			AtomicWrite()
		}()

		go func() {
			// Waiting
			defer wg.Done()
			AtomicSet()
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Benchmark_Race_fixed_interface test
func Benchmark_Race_fixed_interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go MutexWrite() // <- race -
		go MutexSet()   // <- race -
	}
}

// Benchmark_Race_atomic_interface test
func Benchmark_Race_atomic_interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go AtomicWrite() // <- race -
		go AtomicSet()   // <- race -
	}
}
