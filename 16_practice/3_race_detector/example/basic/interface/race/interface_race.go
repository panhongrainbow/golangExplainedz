package interface_race

import (
	"sync"
	"sync/atomic"
)

// I interface
type I interface {
	Set(x int)
}

// T struct
type T struct {
	x int
}

// Set method
func (t *T) Set(x int) {
	t.x = x // alter the value pointed to by the interface variable
}

// obj interface
var obj I

// write function
func Write() {
	t := new(T)
	obj = t // alter the interface variable
}

// Set function
func Set() {
	obj.Set(10) // set the value 10 pointed to by the interface variable
}

// mu mutex
var mu sync.Mutex

// MutexWrite function
func MutexWrite() {
	t := new(T)
	mu.Lock()
	obj = t // alter the interface variable
	mu.Unlock()
}

// MutexSet function
func MutexSet() {
	mu.Lock()
	obj.Set(10) // set the value 10 pointed to by the interface variable
	mu.Unlock()
}

// obj interface
var obj2 atomic.Value

// AtomicWrite function
func AtomicWrite() {
	obj2.Store(new(T))
}

// AtomicSet function
func AtomicSet() {
	// t := obj2.Load().(*T)
	// atomic.
}
