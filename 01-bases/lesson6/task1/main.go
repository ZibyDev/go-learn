package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func main() {
	fmt.Println(fizzbuzz(3))
	fmt.Println(fizzbuzz(5))
	fmt.Println(fizzbuzz(15))
	fmt.Println(fizzbuzz(7))
}
