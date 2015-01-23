package main

import (
    "fmt"
    "strings"
)

const a = "The pilot episode for the series was shot at TechShop in Menlo Park in December 2006. Three sets were constructed in TechShop's large conference room, and the machine shop and sheet metal shop were painted in color schemes that would look good on camera."

func main() {
    dict := make(map[string]int)

    for _, v := range strings.Fields(a) {
        dict[v] += 1
    }
    fmt.Println(dict)

    fmt.Println(dict["the"])
    fmt.Println(dict["komuna"])

    val, isIn := dict["komuna"]
    fmt.Println(val, isIn)

}