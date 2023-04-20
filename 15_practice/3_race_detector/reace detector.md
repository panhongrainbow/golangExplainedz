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