package calculator_test

import (
	"fmt"

	"github.com/brettfirecore/calculator/calculator"
)

func ExampleAddMany() {
	fmt.Println(calculator.AddMany(1, 2, 3, 4))
	// Output: 10
}

func ExampleSubtractMany() {
	fmt.Println(calculator.SubtractMany(10, 1, 2, 3))
	// Output: 4
}

func ExampleMultiplyMany() {
	fmt.Println(calculator.MultiplyMany(2, 3, 4))
	// Output: 24
}

func ExampleDivideMany() {
	res, err := calculator.DivideMany(100, 5, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(res)
	// Output: 10
}
