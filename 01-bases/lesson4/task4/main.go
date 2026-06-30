package main

import (
	"fmt"
	"math"
)

func nearlyEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func main() {
	epsilon := 1e-9
	x := 0.1
	y := 0.2
	a := x + y
	b := 0.3
	fmt.Println(a == b)
	fmt.Println(nearlyEqual(a, b, epsilon))

	c := 1.0
	d := 1.5
	fmt.Println(c == d)
	fmt.Println(nearlyEqual(c, d, epsilon))
}
