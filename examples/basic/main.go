package main

import (
    "fmt"

    "github.com/brettfirecore/calculator/calculator"
)

func main() {
    result := calculator.AddMany(1, 2, 3, 4)
    fmt.Println(result) // 10

    res, err := calculator.DivideMany(10, 2)
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println(res) // 5
    }
}
