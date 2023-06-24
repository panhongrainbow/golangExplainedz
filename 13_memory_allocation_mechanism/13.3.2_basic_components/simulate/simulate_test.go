package main

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
[The arena] is used to allocate objects from a contiguous region of memory.
[The bitmap] is used to keep track of which bytes in the arena are in use and which are free.
[The span] is used to manage memory allocation and deallocation,
*/

// Actually, the heapArenaSize and bitmapSize are 32 to 1, but this time we simulate 8 to 1.
// 这是模拟 8 比 1，而不是 32 比 1
const (
	heapArenaSize = 512
	bitmapSize    = heapArenaSize / 8
)

// simulate heapArena and bitmap
var heap [heapArenaSize]byte
var bitmap [bitmapSize]byte

// |----- bitmap[0] -----| (8 bits) ...... bitmap[heapArenaSize / 8] (共 64 个)
//                  1                        0                        0                        0                        0                        0                        0                        0
// |----- heapArena[0] -----|----- heapArena[1] -----|----- heapArena[2] -----|----- heapArena[3] -----|----- heapArena[4] -----|----- heapArena[5] -----|----- heapArena[6] -----|----- heapArena[7] -----| (共 512 个)
//        (8 bits)                 (8 bits)                 (8 bits)                 (8 bits)                 (8 bits)                 (8 bits)                 (8 bits)                 (8 bits)

// One bit in the bitmap manages 8 bits of heapArena space.
// (在 bitmap 的 1 bit 管 8bits 的 heapArena 空间)

// allocate simulates mspan executing allocate here, but actually we use a list to connect memory spaces.
// Here we directly return the pointer.
func allocate(howManyBytes int) *byte {
	// countBytes is primarily used to find contiguous writable spaces.
	var bitmapStart, bitmapEnd, countBytes int
	for i := 0; i < bitmapSize; i++ {
		if bitmap[i] == 0xff { // The value is 1111 1111.
			// Each bit in the bitmap represents 8 bits of heapArena space, and each 8 bits in the bitmap represents 64 bits of heapArena space.
			// All 64 bits of heapArena space are fully occupied.

			// I need to find contiguous space, but the bitmap is fully occupied, causing an interruption.
			// Therefore, the count is reset to 0, and I will search for contiguous space again.
			countBytes = 0 // (就是这一步不能省 !)
			continue       // Do not search for contiguous space in this bitmap anymore
		}
		for j := 0; j < 8; j++ {
			if bitmap[i]&(1<<j) == 0 {
				countBytes++
				if countBytes == howManyBytes { // If the required contiguous space is reached.

					// If j ranges from 0 to 7, adding 0 here, is it logical? So j should be incremented by 1.
					bitmapEnd = i*8 + (j + 1) // (j 由 0 到 7 ，这里加 0 ，合理吗？所以要 j 加 1)
					bitmapStart = bitmapEnd - howManyBytes

					fmt.Println("Allocate the starting position of the bitmap: ", bitmapStart)

					for k := bitmapStart; k < bitmapEnd; k++ {
						bitmap[k/8] |= 1 << (k % 8)
						// If bitmapStart to bitmapEnd is from 0 to 7, all values in bitmap[0] will be set to 1.
						// If bitmapStart to bitmapEnd is from 8 to 15, all values in bitmap[1] will be set to 1.
						// Etc.
					}

					// Return the starting position of the writable space, which is contiguous
					// Therefore, only the starting space needs to be returned.

					return &heap[bitmapStart] // 回传连续的起始空间就好
				}
			} else {
				// In case the space is not contiguous, reset countBytes
				countBytes = 0 // (就是这一步不能省 !)
			}
		}
	}
	return nil
}

// release is used to free up space.
func release(ptr *byte, size int) {
	// Calculate bitmapStart based on the position of the content
	bitmapStart := int(uintptr(unsafe.Pointer(ptr)) - uintptr(unsafe.Pointer(&heap[0])))

	for i := 0; i < size; i++ {
		// Perform an XOR operation; when both are zero or both are one, it will be reset to zero.
		bitmap[(bitmapStart+i)/8] ^= 1 << ((bitmapStart + i) % 8)
	}
}

// Test_simulate performs unit testing.
func Test_simulate(t *testing.T) {
	// Initialize the bitmap.
	for i := 0; i < bitmapSize; i++ {
		bitmap[i] = 0
	}

	// Allocate a space of 12 bytes
	ptr1 := allocate(12)

	fmt.Println("The starting position of the 12-byte contiguous space: ", ptr1)

	fmt.Printf("Current Bitmap: %08b\n", bitmap)

	// Allocate a space of 1 byte
	ptr2 := allocate(1)

	fmt.Println("The starting position of the 1-byte contiguous space: ", ptr2)

	fmt.Printf("Current Bitmap: %08b\n", bitmap)

	// Release a space of 1 byte
	release(ptr2, 1)

	fmt.Printf("Current Bitmap: %08b\n", bitmap)

	// Release a space of 12 bytes
	release(ptr1, 12)

	fmt.Printf("Current Bitmap: %08b\n", bitmap)
}
