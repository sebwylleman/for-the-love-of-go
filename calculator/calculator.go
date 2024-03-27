// Package calculator does simple calculations.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the
// result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b, and
// returns their product.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the
// result of dividing the first by the
// second, or an error if the second value
// is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

// Sqrt returns the square root of its input,
// or an error for negative inputs.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("invalid input, only positive numbers allowed")
	}
	return math.Sqrt(a), nil
}
