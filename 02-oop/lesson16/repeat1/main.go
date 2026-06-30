package main

import "fmt"

type Duration struct {
	Hours, Minutes int
}

func (d Duration) String() string {
	return fmt.Sprintf("%dч %dм", d.Hours, d.Minutes)
}

func main() {
	d1 := Duration{Hours: 2, Minutes: 30}
	d2 := Duration{Hours: 0, Minutes: 45}
	d3 := Duration{Hours: 12, Minutes: 0}
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
}
