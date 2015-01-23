package main

import (
    "fmt"
)

func main() {
    // creating slice
    s1 := []int{1, 3, 4, 5}
    fmt.Println(s1[1:3])
    fmt.Println(s1[:])
    fmt.Println(s1[1:])
    fmt.Println(s1[:3])

    fmt.Println(len(s1), cap(s1), s1)

    s2 := make([]int, 2, 5)
    for i := range s2 {
        s2[i] = 0
    }
    fmt.Println(len(s2), cap(s2), s2)

    // append
    s2 = append(s2, 4)
    s2 = append(s2, 3)
    s2 = append(s2, 2)
    s2 = append(s2, 2)
    // watch how capacity changed
    fmt.Println(len(s2), cap(s2), s2)
}