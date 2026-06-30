package main

import "fmt"

func fib(n int) int {
	var a, b = 0, 1
	if n == 0 {
		return a
	}
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(7))
}
