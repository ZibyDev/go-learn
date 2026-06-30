package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	c1 := Circle{3}
	c2 := Circle{1}
	c3 := Circle{7}
	fmt.Println(c1.Area(), c1.Circumference())
	fmt.Println(c2.Area(), c2.Circumference())
	fmt.Println(c3.Area(), c3.Circumference())
}
