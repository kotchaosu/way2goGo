package main

import (
    "fmt"
)

func main() {
    if i := 3; i < 5 {
        fmt.Println("You got it")
    } else if i == 4 {
        fmt.Println("Go back to you cave!")
    } else {
        fmt.Println("Too much coffee, huh?")
    }
}