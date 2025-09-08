// Package calculator with errors and math import
package calculator

import (
	"fmt"
	"math"
)

// --- Basic two-arg API (what your tests expect) ---

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("square root of negative number")
	}
	return math.Sqrt(x), nil
}

// --- Variadic versions (your new API) ---

func AddMany(inputs ...float64) float64 {
	var sum float64
	for _, v := range inputs {
		sum += v
	}
	return sum
}

func MultiplyMany(inputs ...float64) float64 {
	prod := 1.0
	for _, v := range inputs {
		prod *= v
	}
	return prod
}

func SubtractMany(inputs ...float64) float64 {
	switch len(inputs) {
	case 0:
		return 0
	case 1:
		return inputs[0]
	default:
		res := inputs[0]
		for _, v := range inputs[1:] {
			res -= v
		}
		return res
	}
}

func DivideMany(inputs ...float64) (float64, error) {
	switch len(inputs) {
	case 0:
		return 0, fmt.Errorf("DivideMany: need at least one input")
	case 1:
		return inputs[0], nil
	default:
		res := inputs[0]
		for _, v := range inputs[1:] {
			if v == 0 {
				return 0, fmt.Errorf("DivideMany: divide by zero")
			}
			res /= v
		}
		return res, nil
	}
}
