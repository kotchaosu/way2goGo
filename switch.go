package main

import (
    "fmt"
)

func main() {
    i := 3
    switch {
    case i < 5:
        fmt.Println("You got it, again...")
    case i == 4:
        fmt.Println("Go back to you cave!")
    default:
        fmt.Println("Too much coffee, huh?")
    }
}