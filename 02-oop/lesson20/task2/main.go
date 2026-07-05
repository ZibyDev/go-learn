package main

import "fmt"

type Calculator struct{}

func (c Calculator) Add(a, b int) int { return a + b }

type LoggingCalculator struct {
	Calculator
}

func (lc LoggingCalculator) Add(a, b int) int {
	fmt.Printf("[calc] ")
	res := lc.Calculator.Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, res)
	return res
}

func main() {
	lc := LoggingCalculator{}
	fmt.Println(lc.Add(3, 5))
	fmt.Println(lc.Calculator.Add(3, 5))
}
