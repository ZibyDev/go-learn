package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	slice[0] = 100
	fmt.Println(arr)
	fmt.Println(slice)
}
