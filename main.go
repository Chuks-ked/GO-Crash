package main

import "fmt"

func main() {
    x := 110
    p := &x // p holds the memory address of x

    fmt.Println("Value of x:", x)
    fmt.Println("Address of x:", p)
    fmt.Println("Value at address p points to:", *p)
}
