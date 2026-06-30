package main

import "fmt"

func removeAt(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}
	sliceCopy := make([]int, 0, len(s)-1)
	sliceCopy = append(sliceCopy, s[:i]...)
	sliceCopy = append(sliceCopy, s[i+1:]...)
	return sliceCopy
}

func main() {
	slice := []int{10, 20, 30, 40}
	fmt.Println(slice)
	result := removeAt(slice, 1)
	fmt.Println(slice)
	fmt.Println(result)
}
