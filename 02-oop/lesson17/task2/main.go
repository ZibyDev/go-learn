package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }

func (c Circle) Area() float64    { return 3.14159 * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 7, Height: 4},
		Circle{Radius: 5},
		Circle{Radius: 4},
		Rectangle{Width: 11, Height: 5},
	}
	fmt.Println(totalArea(shapes))
}
