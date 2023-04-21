# race detector list

> Here I list all the places where `race detector` is needed, including examples and instructions for actual operation.

## Introduction

If the code contains the following content, then consider using `race detect`.

1. When `global variables` or `static variables` are shared among multiple Goroutines.
2. As long as the program contains `goroutine`
3. As long as the program contains `channel`
4. As long as the program contains `syc package`
5. As long as the program contains `closure`
6. As long as the program `accesses pointer variables`
7. As long as the program `accesses interface variables`
8. Using `maps` for concurrent reading and writing
9. Reading and writing `configuration information` in multiple Goroutines

(2023/4/19)

## Instructions order

The CPU determines how the compiler and processor reorder instructions.

### CPU Memory Model:

[reference](https://go.dev/ref/mem)

its memory model as `Sequential Consistency` (not Relaxed), which means that in a concurrent environment, all CPUs must access memory in sequence and there is no reordering between them.

For example:
```go
a = 1   // Statement 1
b = 2   // Statement 2
x = b   // Statement 3 
y = a   // Statement 4
```

In `Sequential Consistency`, whether it's one CPU at different times or multiple CPUs at the same time, the order must be `1 -> 2 -> 3 -> 4`. It cannot be any other order.

```In the absence of data races, Go programs behave as if all the goroutines were multiplexed onto a single processor. This property is sometimes referred to as DRF-SC: data-race-free programs execute in a sequentially consistent manner. ``` [The Go Memory Model](https://go.dev/ref/mem)

I means as long as there are no data races, the behavior of Go routines will be the same as the main program.

### Memory Barrier

> Memory barriers only ensure `the order of instruction reordering`, but do `not prevent race conditions`.

#### Introduction

When the `compiler or CPU` may reorder instructions to improve performance, the `Go runtime` inserts `memory barriers` to prohibit specific types of reordering to achieve memory consistency.

For example:



## Atomic

> Atomic operations are implemented at the hardware level and can be used to fix data races when multiple threads access and modify the same variable at the same time.
>
> [bug reference](https://pkg.go.dev/sync/atomic)

### Introduction

Atomic operations are implemented at the hardware level using special cache and memory coherency and bus synchronization protocols (channel and lock) .

These protocols ensure that atomic operations are executed in a thread-safe manner, without interference from other threads. 

### bug



On 386, the 64-bit functions use instructions unavailable before the Pentium MMX.



On non-Linux ARM, the 64-bit functions use instructions unavailable before the ARMv6k core.



On ARM, 386, and 32-bit MIPS, it is the caller's responsibility to arrange for 64-bit alignment of 64-bit words accessed atomically via the primitive atomic functions (types [Int64](https://pkg.go.dev/sync/atomic#Int64) and [Uint64](https://pkg.go.dev/sync/atomic#Uint64) are automatically aligned). The first word in an allocated struct, array, or slice; in a global variable; or in a local variable (because the subject of all atomic operations will escape to the heap) can be relied upon to be 64-bit aligned.

(2023/4/19)