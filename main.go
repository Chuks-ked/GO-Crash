package main

import "fmt"

func double(n *int) {
    *n = *n * 2
}

func main() {
    num := 5
    double(&num)
    fmt.Println(num) // 10
}
