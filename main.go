package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
    return "Woof"
}

func (c Cat) Speak() string {
    return "Meow"
}

func makeSound(a Animal) {
    fmt.Println(a.Speak())
}

func main() {
    dog := Dog{}
    cat := Cat{}
    makeSound(dog)
    makeSound(cat)
}
