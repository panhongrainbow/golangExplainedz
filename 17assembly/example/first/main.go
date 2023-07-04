package main

import "fmt"

func main() {
	a, b := 1, 2
	c := Add(a, b)
	fmt.Println(c)
}

//go:noescape
func Add(a, b int) int
