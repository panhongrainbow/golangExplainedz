package race

import (
	"sync"
	"testing"
)

type demo struct {
	status uint32
	count  uint32
	mu     sync.Mutex
}

func (d *demo) setStatus() {
	/*if !atomic.CompareAndSwapUint32(&d.status, 0, 1) {
		return
	}*/

	d.mu.Lock()
	if d.status == 1 {
		return
	}
	d.mu.Unlock()

	d.status = 1

	d.count++
}

func Test_Race_atomic(t *testing.T) {
	var d demo

	for i := 0; i < 1000; i++ {
		go func() {
			d.setStatus()
		}()
	}
}
