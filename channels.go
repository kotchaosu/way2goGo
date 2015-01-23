package main

import (
    "fmt"
)

func sayHello(name string, done chan<- bool) {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello", name)
    }
    done<-true
}

func main() {
    done := make(chan bool)
    go sayHello("Ambroży", done)

    <-done
}