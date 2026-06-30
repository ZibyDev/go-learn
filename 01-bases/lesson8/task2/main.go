package main

import "fmt"

func minMax(a, b, c int) (int, int) {
	var max, min = a, a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min, max
}

func main() {
	fmt.Println(minMax(5, 2, 8))
	fmt.Println(minMax(3, 3, 3))
}
