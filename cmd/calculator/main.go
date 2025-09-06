package main

import (
	"fmt"
	"log"

	"github.com/brettfirecore/calculator/calculator"
)

func main() {
	// Add
	sum := calculator.Add(2, 3)
	fmt.Printf("2 + 3 = %v\n", sum)

	// Subtract
	diff := calculator.Subtract(10, 4)
	fmt.Printf("10 - 4 = %v\n", diff)

	// Multiply
	product := calculator.Multiply(6, 7)
	fmt.Printf("6 * 7 = %v\n", product)

	// Divide
	quotient, err := calculator.Divide(20, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("20 / 5 = %v\n", quotient)

	// Divide by zero (error case)
	_, err = calculator.Divide(10, 0)
	if err != nil {
		fmt.Printf("Divide by zero error: %v\n", err)
	}

	// Square root
	root, err := calculator.Sqrt(81)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("√81 = %v\n", root)

	// Square root of negative number (error case)
	_, err = calculator.Sqrt(-9)
	if err != nil {
		fmt.Printf("Sqrt of negative error: %v\n", err)
	}
}
