package main

import (
	"fmt"
	_ "github.com/panhongrainbow/golangExplainedz/17assembly/example/first/add2"
	_ "unsafe"
)

func main() {
	a, b := 1, 2
	c := Add(a, b)
	fmt.Println(c)
}

//go:noescape
//go:linkname Add github.com/panhongrainbow/golangExplainedz/17assembly/example/first/add2.Add2
func Add(a, b int) int
