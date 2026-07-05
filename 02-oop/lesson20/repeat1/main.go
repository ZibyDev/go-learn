package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64    { return 3.14159 * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Height * r.Width }

func report(shapes []Shape) {
	for _, x := range shapes {
		a := x.Area()
		switch v := x.(type) {
		case Circle:
			fmt.Printf("It's Circle with radius = %g. It's area = %g\n", v.Radius, a)
		case Rectangle:
			fmt.Printf("It's Rectangle with width = %g and height = %g. It's area = %g\n", v.Width, v.Height, a)
		}
	}
}

func main() {
	s := []Shape{
		Circle{Radius: 3.4},
		Rectangle{Width: 8, Height: 2.3},
	}
	report(s)
}
