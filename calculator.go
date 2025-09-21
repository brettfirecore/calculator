// Package calculator does simple calculations.
// This is the code under test — each function here will be tested in calculator_test.go.
package calculator

// Import the errors package from Go’s standard library.
// It provides a simple way to create error values (errors.New).
import "errors"

// Add takes two numbers and returns their sum.
func Add(a, b float64) float64 {
	// The + operator adds the two float64 values together.
	// Example: Add(2, 3) returns 5.
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	// The - operator subtracts the second value from the first.
	// Example: Subtract(10, 4) returns 6.
	return a - b
}

// Multiply takes two numbers and returns their product.
func Multiply(a, b float64) float64 {
	// The * operator multiplies the two float64 values together.
	// Example: Multiply(3, 3) returns 9.
	return a * b
}

// Divide takes two numbers and returns the result of dividing a by b,
// along with an error value if b == 0.
func Divide(a, b float64) (float64, error) {
	// Guard clause: division by zero is invalid.
	// Return 0 as a dummy value, plus a non-nil error.
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}

	// Normal case: return the computed result and nil error.
	return a / b, nil
}
