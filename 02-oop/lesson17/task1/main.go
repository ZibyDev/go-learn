package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }

func (c Circle) Area() float64    { return 3.14159 * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }

func printArea(s Shape) { fmt.Println(s.Area()) }

func main() {
	printArea(Circle{Radius: 4})
	printArea(Rectangle{Width: 8, Height: 6})
}
