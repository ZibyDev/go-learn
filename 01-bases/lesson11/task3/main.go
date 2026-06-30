package main

import "fmt"

func divide(a, b int) *int {
	if b == 0 {
		return nil
	} else {
		res := a / b
		return &res
	}
}

func main() {
	x1 := divide(10, 5)
	if x1 != nil {
		fmt.Println(*x1)
	}
	x2 := divide(10, 0)
	if x2 != nil {
		fmt.Println(*x2)
	}
}
