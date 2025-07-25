package main

import (
	"fmt"

	"calculator" // or your module path from go.mod
)

func main() {
	result := calculator.Add(2, 2)
	fmt.Println("2 + 2 =", result)
}
