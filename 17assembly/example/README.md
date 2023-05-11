 

# Assembly

## NOP

`NOP instructions` are used to `adjust the size of a code fragment`, as `a placeholder for future instructions`, to `delay execution of other instructions`, for `debugging, testing`, and to `give other hardware devices time to respond`.

Here is an example in Go:

```go
go
package main

func main() {
  // Do something
  fmt.Println("Hello, world!")
  // Do nothing - this is effectively a NOP instruction
  _ = 1  
  // Do something else
  fmt.Println("Goodbye, world!")
}
```

