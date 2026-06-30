package main

import "fmt"

func filter(s []int, keep func(int) bool) []int {
	newSlice := []int{}
	for _, v := range s {
		if keep(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)
	slice = filter(slice, func(x int) bool { return x%3 == 0 })
	fmt.Println(slice)
}
