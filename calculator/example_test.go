package calculator_test

import (
	"fmt"
	"github.com/brettfirecore/calculator/calculator"
)

func ExampleAddMany() {
	fmt.Println(calculator.AddMany(1, 2, 3, 4))
	// Output: 10
}


func ExampleDivide() {
    q, err := calculator.Divide(20, 5)
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    fmt.Println(q)
    // Output: 4
}

func ExampleSqrt() {
	r, err := calculator.Sqrt(81)
	if err != nil {
		return
	}
		fmt.Println(r)
	// Output: 9
}
