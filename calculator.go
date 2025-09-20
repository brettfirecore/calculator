// Package calculator does simple calculations.
// Package calculator defines simple arithmetic functions.
// By putting code in a package, it can be imported and tested.
package calculator

// Add takes two numbers and returns their sum.
func Add(a, b float64) float64 {
	// The + operator adds the two float64 values together
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	// The - operator subtracts the second value from the first
	return a - b
}

// Multiply takes two numbers and returns their product.
func Multiply(a, b float64) float64 {
	// The * operator multiplies the two values together
	return a * b
}
