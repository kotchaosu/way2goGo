package main

import (
    "fmt"
    "time"
)

func sendBollocks(delay int, c chan string) {
    time.Sleep(time.Second * time.Duration(delay))
    c <- "Bollocks"
}

func main() {
    c1, c2 := make(chan string), make(chan string)

    go sendBollocks(1, c1)
    go sendBollocks(2, c2)

    for i := 0; i < 2; i++ {
        select {
        case <-c1:
            fmt.Println("Mam to!")
        case name := <-c2:
            fmt.Printf("Czy to %s?", name)
        }
    }

}