package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1.4
	var b float64 = 1.5
	var c float64 = 2.5
	var d float64 = -1.5

	fmt.Println(a)
	fmt.Println(int(a))
	fmt.Println(math.Round(a))
	fmt.Println(math.Floor(a))
	fmt.Println(math.Ceil(a))
	fmt.Println()
	fmt.Println(b)
	fmt.Println(int(b))
	fmt.Println(math.Round(b))
	fmt.Println(math.Floor(b))
	fmt.Println(math.Ceil(b))
	fmt.Println()
	fmt.Println(c)
	fmt.Println(int(c))
	fmt.Println(math.Round(c))
	fmt.Println(math.Floor(c))
	fmt.Println(math.Ceil(c))
	fmt.Println()
	fmt.Println(d)
	fmt.Println(int(d))
	fmt.Println(math.Round(d))
	fmt.Println(math.Floor(d))
	fmt.Println(math.Ceil(d))
}
