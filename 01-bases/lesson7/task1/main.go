package main

import "fmt"

func factorial(n int) int {
	var res = 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	fmt.Println(factorial(5))
}
