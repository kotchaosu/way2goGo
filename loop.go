package main

import (
    "fmt"
)

func main() {
    // C smell I...
    for i:= 0; i < 5; i++ {
        fmt.Print(i)
    }

    // this raises undefined err
    // fmt.Println(i)

    fmt.Println()
    // while
    i := 0
    for i < 5 {
        fmt.Print(i)
        i++
    }

    // this is not wise
    // for {
    //     fmt.Println("Infinity, babe")
    // }
}