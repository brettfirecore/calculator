// Package calculator does simple calculations.
package calculator

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two number a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b // swap b and a around so its 2 from 4 not 4 from 2 
}
