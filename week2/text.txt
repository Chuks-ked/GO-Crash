package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var filename string
    fmt.Print("Enter filename: ")
    fmt.Scanln(&filename)

    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        count++
    }

    fmt.Println("Total lines:", count)
}
