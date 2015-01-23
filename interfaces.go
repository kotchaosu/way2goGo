package main

import (
    "fmt"
)

type Animaler interface {
    Eat(int)
    MakeSound(int)
}

type Mammal struct {
    Weight int
    HasFur bool
    Sound string
}

func  (m *Mammal) Eat(w int) {
    m.Weight += w * 2
}

func (m *Mammal) MakeSound(n int) {
    for i := 0; i < n; i++ {
        fmt.Println(m.Sound)
    }
}

type Reptile struct {
    Weight int
}

func (r *Reptile) Eat(w int) {
    r.Weight += w
}

func (r *Reptile) MakeSound(n int) {
    fmt.Println("I don' wanna talk")
}

func main() {
    m := &Mammal{70, true, "Bollocks"}
    fmt.Println(m)

    a := Animaler(m)
    fmt.Println(a)

    r := &Reptile{}

    animals := []Animaler{m, r}
    for _, animal := range animals {
        fmt.Println("I'm ", animal)
        animal.MakeSound(4)
    }
}