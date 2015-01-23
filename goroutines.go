// same thing as in Tour of Go
package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println("Hello", name)
    }
}

func main() {
    go sayHello("Ambroży")
    sayHello("Adam")
}