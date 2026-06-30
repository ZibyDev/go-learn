package main

import "fmt"

func calculator() (add func(int), result func() int) {
	sum := 0
	add = func(x int) {
		sum += x
	}
	result = func() int {
		return sum
	}
	return
}

func main() {
	add, result := calculator()
	add(10)
	fmt.Println(result())
	add(5)
	fmt.Println(result())
	add(35)
	fmt.Println(result())
}
