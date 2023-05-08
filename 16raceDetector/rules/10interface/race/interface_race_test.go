package interface_race

import (
	"sync"
	"testing"
)

// Test_Race_interface fixed that interface are not synchronized.
func Test_Race_interface(t *testing.T) {
	// use wait group to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Start 1000 goroutines
	for i := 0; i < 500; i++ {
		// Goroutines modifies the interface variable together
		go func() {
			defer wg.Done()
			Write() // <- race -
		}()

		go func() {
			defer wg.Done()
			Set() // <- race -
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Test_Race_interface fixed that interface are in synchronized.
func Test_fixed_interface(t *testing.T) {
	// use wait group to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Start 1000 goroutines
	for i := 0; i < 500; i++ {
		// Goroutines modifies the interface variable together
		go func() {
			defer wg.Done()
			MutexWrite() // fixed (1/2) !
		}()

		go func() {
			defer wg.Done()
			MutexSet() // fixed (2/2) !
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Test_atomic_interface fixed that interface are in synchronized.
func Test_atomic_interface(t *testing.T) {
	// use wait group to wait for all interface to finish
	var wg sync.WaitGroup
	wg.Add(1000)

	// Start 1000 goroutines
	for i := 0; i < 500; i++ {
		// Goroutines modifies the interface variable together
		go func() {
			// Waiting
			defer wg.Done()
			AtomicWrite() // fixed (1/2) !
		}()

		go func() {
			// Waiting
			defer wg.Done()
			AtomicSet() // fixed (2/2) !
		}()
	}

	// Wait for all interface to finish
	wg.Wait()
}

// Benchmark_Race_fixed_interface test
func Benchmark_Race_fixed_interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go MutexWrite()
		go MutexSet()
	}
}

// Benchmark_Race_atomic_interface test
func Benchmark_Race_atomic_interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go AtomicWrite()
		go AtomicSet()
	}
}
