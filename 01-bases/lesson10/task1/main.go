package main

import "fmt"

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter1())
}
