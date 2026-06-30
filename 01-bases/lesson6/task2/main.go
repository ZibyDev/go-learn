package main

import "fmt"

func calc(a, b float64, op string) (float64, bool) {
	switch {
	case op == "+":
		return a + b, true
	case op == "-":
		return a - b, true
	case op == "*":
		return a * b, true
	case op == "/":
		if b == 0 {
			return 0, false
		}
		return a / b, true
	default:
		return 0, false
	}
}

func main() {
	fmt.Println(calc(10, 5, "+"))
	fmt.Println(calc(10, 0, "/"))
	fmt.Println(calc(10, 5, "%"))
}
