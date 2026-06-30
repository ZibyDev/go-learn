package main

import "fmt"

func toInt(x any) (int, bool) {
	v, ok := x.(int)
	return v, ok
}

func main() {
	fmt.Println(toInt(42))
	fmt.Println(toInt("hello"))
	fmt.Println(toInt(3.14))
}
