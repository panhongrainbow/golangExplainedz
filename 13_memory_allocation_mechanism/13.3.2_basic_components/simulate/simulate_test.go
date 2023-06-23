package main

import (
	"fmt"
	"testing"
	"unsafe"
)

const (
	heapSize   = 512
	bitmapSize = heapSize / 8
)

var heap [heapSize]byte
var bitmap [bitmapSize]byte

func allocate(size int) *byte {
	var start, count int
	for i := 0; i < bitmapSize; i++ {
		if bitmap[i] == 0xff {
			count = 0
			continue
		}
		for j := 0; j < 8; j++ {
			if bitmap[i]&(1<<j) == 0 {
				count++
				if count == size {
					start = i*8 + j - size + 1
					for k := start; k < start+size; k++ {
						bitmap[k/8] |= 1 << (k % 8)
					}
					return &heap[start]
				}
			} else {
				count = 0
			}
		}
	}
	return nil
}

func release(ptr *byte, size int) {
	offset := int(uintptr(unsafe.Pointer(ptr)) - uintptr(unsafe.Pointer(&heap[0])))

	for i := 0; i < size; i++ {
		bitmap[(offset+i)/8] ^= 1 << ((offset + i) % 8)
	}
}

func Test_simulate(t *testing.T) {
	for i := 0; i < bitmapSize; i++ {
		bitmap[i] = 0
	}

	ptr1 := allocate(12)

	fmt.Printf("Bitmap: %08b\n", bitmap)

	ptr2 := allocate(1)

	fmt.Printf("Bitmap: %08b\n", bitmap)

	release(ptr2, 1)

	fmt.Printf("Bitmap: %08b\n", bitmap)

	release(ptr1, 12)

	fmt.Printf("Bitmap: %08b\n", bitmap)
}
