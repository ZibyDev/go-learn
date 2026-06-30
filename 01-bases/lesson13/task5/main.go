package main

import "fmt"

func appendBad(s []int) {
	s = append(s, 100)
}

func appendGood(s []int) []int {
	return append(s, 100)
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	appendBad(slice)
	fmt.Println(slice)
	slice = appendGood(slice)
	fmt.Println(slice)
}
