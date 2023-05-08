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
	t.x = x // Alter the value pointed to by the interface variable
}

// obj interface
var obj I // ----- race ----->

// Write function
func Write() {
	obj = new(T) // Alter the interface variable // <----- race -----
}

// Set function
func Set() {
	obj = new(T)
	obj.Set(10) // Set the value 10 pointed to by the interface variable // <----- race -----
}

// Mu mutex
var mu sync.Mutex

// MutexWrite function
func MutexWrite() {
	mu.Lock()    // add (1/4) !
	obj = new(T) // Alter the interface variable
	mu.Unlock()  // add (2/4) !
}

// MutexSet function
func MutexSet() {
	mu.Lock() // add (3/4) !
	obj = new(T)
	obj.Set(10) // Set the value 10 pointed to by the interface variable
	mu.Unlock() // add (4/4) !
}

// obj interface
var obj2 atomic.Value

func init() {
	obj2.Store(new(T)) // Not atomic !!!
}

// mu2 mutex
var mu2 sync.Mutex

// AtomicWrite function
func AtomicWrite() {
	// obj2.Store(new(T)) // Big no, no, not atomic !!!
	t := new(T)
	obj2.Store(t) // Atomic !!! // correct (1/2) !
}

// AtomicSet function
func AtomicSet() {
	mu2.Lock()
	obj2.Load().(*T).Set(10) // Big no, no, not atomic !!! need to use mutex // correct (2/2) !
	mu2.Unlock()
}
