// Package calculator provides basic float64 arithmetic with simple error handling.
package calculator

import (
	"errors"
	"math"
)

// Add returns the sum of two float64 numbers.
func Add(a, b float64) float64 {
	return a + b
}

// Multiply returns the product of two float64 numbers.
func Multiply(a, b float64) float64 {
	return a * b
}

// Subtract returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Divide returns the result of dividing a by b.
// If b is zero, it returns an error to avoid division by zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// Return 0 and an error if division by zero is attempted
		return 0, errors.New("division by zero not allowed")
	}
	// Return the result of the division and no error
	return a / b, nil
}

// Sqrt returns the square root of a float64 value.
// It returns an error if the input is negative.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("square root of negative number not allowed")
	}
	return math.Sqrt(x), nil
}
