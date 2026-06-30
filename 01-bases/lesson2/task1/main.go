package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var u *int

	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("%T %v\n", u, u)
}
