package main

import "fmt"

func tryDouble(x int) {
	x *= 2
	fmt.Println("Удвоение без указателя:", x)
}

func doublePtr(x *int) {
	*x *= 2
	fmt.Println("Удвоение с указателем", *x)
}

func main() {
	x := 2
	fmt.Println("Первоначальное значение до функций:", x)
	tryDouble(x)
	fmt.Println("Первоначальное значение после tryDouble:", x)
	doublePtr(&x)
	fmt.Println("Первоначальное значение после doublePtr:", x)
}
