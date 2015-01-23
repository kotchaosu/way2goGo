package main

import (
    "fmt"
)

type Animal struct {
    Weight int
    HasFur bool
    Sound string
}

func (a *Animal) makeSound(n int) {
    for i := 0; i < n; i++ {
        fmt.Println(a.Sound)
    }
}

func (a *Animal) eat(w int) {
    a.Weight += w
}

func main() {
    mammal := &Animal{70, true, "Bollocks!"}
    fmt.Println(mammal)

    mammal.makeSound(4)

    // pointers matters
    mammal.eat(5)
    fmt.Println(mammal)
}