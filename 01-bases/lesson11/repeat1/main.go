package main

import "fmt"

func makeAdder(start *int) (func(int), func() int) {
	add := func(x int) {
		*start = *start + x
	}
	result := func() int {
		return *start
	}
	return add, result
}

func main() {
	s := 0
	add, res := makeAdder(&s)
	add(5)
	fmt.Println(res())
	add(10)
	fmt.Println(res())
	add(15)
	fmt.Println(res())
}
