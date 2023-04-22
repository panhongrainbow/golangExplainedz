package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Sending value on channel")
		ch <- 42
	}()

	go func() {
		fmt.Println("Receiving value from channel")
		value := <-ch
		fmt.Println("Received value:", value)
	}()

	time.Sleep(time.Second)
}
