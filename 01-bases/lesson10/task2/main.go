package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(2))
	fmt.Println(double(3))
	fmt.Println(double(4))
	fmt.Println(triple(2))
	fmt.Println(triple(3))
	fmt.Println(triple(4))
}
