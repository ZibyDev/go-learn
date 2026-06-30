package main

import "fmt"

func main() {
	const factor = 100

	var factorInt int = factor
	var factorFloat64 float64 = factor
	var factorInt64 int64 = factor

	fmt.Println(factorInt)
	fmt.Println(factorFloat64)
	fmt.Println(factorInt64)

	const typedFactor int = 100
	var typedFactorInt int = typedFactor
	var typedFactorFloat64 float64 = float64(typedFactor) // Явное приведение
	var typedFactorInt64 int64 = int64(typedFactor)       // Явное приведение

	fmt.Println(typedFactorInt)
	fmt.Println(typedFactorFloat64)
	fmt.Println(typedFactorInt64)
}
