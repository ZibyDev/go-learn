package main

import "fmt"

func modify(arr [4]int) {
	arr[2] = 1234
}

func main() {
	a := [4]int{1, 2, 3, 4}
	b := a
	b[2] = 1234
	fmt.Println(a)
	fmt.Println(b)
	modify(a)
	fmt.Println(a)
}
