# Memory management mspan

The basic components of Golang’s memory allocation are `mcache, mcentral, and mheap`. The basic unit of Go memory management is the `mspan`.

## Memory Request Allocation

### Introduction

When a Go program starts, it asks the operating system for some memory.

It then divides this memory into smaller parts and manages it on its own.

On an`X64` system, the memory is divided into three areas of `512MB, 16GB, and 512GB`.

### heap arena

<img src="../../../assets/image-20230413001950194.png" alt="image-20230413001950194" style="zoom:80%;" /> 

Go divides `heap arena` into small parts called `8KB pages`.

Some `8KB pages` are grouped together and called an `mspan`(多个 8K 页面，会组成 mspan)

(2023/4/12)

### bitmap

<img src="../../../assets/image-20230413004702839.png" alt="image-20230413004702839" style="zoom:80%;" /> 

`The bitmap area` shows which addresses in the arena area have objects.(不是指向 page ，而是 page 内的内存地址)

<img src="../../../assets/image-20230413004702839a.png" alt="image-20230413004702839" style="zoom:80%;" />

It uses a flag with `some bits` to show if `the object` has `a pointer` or `GC tag`.

<img src="../../../assets/image-20230414094920504.png" alt="image-20230414094920504" style="zoom:80%;" />

The space required for the `bitmap` is `512GB divided by 32`, which equals `16GB`.

(2023/4/14)

