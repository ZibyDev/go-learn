package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func area(r Rectangle) float64 {
	return r.Height * r.Width
}

func perimeter(r Rectangle) float64 {
	return r.Height*2 + r.Width*2
}

func main() {
	rec := Rectangle{Height: 3, Width: 4}
	fmt.Printf("%+v\n", rec)
	fmt.Println(area(rec))
	fmt.Println(perimeter(rec))
}
