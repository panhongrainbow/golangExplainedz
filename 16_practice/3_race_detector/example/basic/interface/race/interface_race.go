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

// Write function
func Write() {
	obj = new(T) // alter the interface variable
}

// Set function
func Set() {
	obj = new(T)
	obj.Set(10) // set the value 10 pointed to by the interface variable
}

// mu mutex
var mu sync.Mutex

// MutexWrite function
func MutexWrite() {
	mu.Lock()
	obj = new(T) // alter the interface variable
	mu.Unlock()
}

// MutexSet function
func MutexSet() {
	mu.Lock()
	obj = new(T)
	obj.Set(10) // set the value 10 pointed to by the interface variable
	mu.Unlock()
}

// obj interface
var obj2 atomic.Value

func init() {
	obj2.Store(new(T)) // not atomic !!!
}

// mu2 mutex
var mu2 sync.Mutex

// AtomicWrite function
func AtomicWrite() {
	// obj2.Store(new(T)) // big no, no, not atomic !!!
	t := new(T)
	obj2.Store(t) // atomic !!!
}

// AtomicSet function
func AtomicSet() {
	mu2.Lock()
	obj2.Load().(*T).Set(10) // big no, no, not atomic !!! need to use mutex
	mu2.Unlock()
}
