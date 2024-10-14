package testing

import (
	"math"
)

func Add(n ...float64) float64 {
	sum := 0.0

	for _, v := range n {
		sum += v
	}

	return sum
}

func Subtract(n ...float64) float64 {
	if len(n) == 0 {
		return 0
	}

	sub := n[0]

	for _, v := range n[1:] {
		sub -= v
	}

	return sub
}

func Multiply(n ...float64) float64 {
	mul := 1.0

	for _, v := range n {
		mul *= v
	}

	return mul
}

func Divide(n ...float64) float64 {
	if len(n) == 0 {
		return math.NaN()
	}

	div := n[0]

	for _, v := range n[1:] {
		div /= v
	}

	return div
}
