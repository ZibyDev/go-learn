package main

import "fmt"

func main() {
	s := make([]int, 0, 2)
	fmt.Println(s, len(s), cap(s))
	for i := 1; i < 10; i++ {
		s = append(s, i)
		fmt.Println(s, len(s), cap(s))
	}
}
