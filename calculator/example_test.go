package calculator_test

import (
	"fmt"
	"log"

	"github.com/brettfirecore/calculator/calculator"
)

func ExampleAdd() {
	fmt.Println(calculator.Add(2, 3))
	// Output: 5
}

func ExampleDivide() {
	q, err := calculator.Divide(20, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q)
	// Output: 4
}

func ExampleSqrt() {
	r, err := calculator.Sqrt(81)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	// Output: 9
}
