package main

import "fmt"

func reverseSlice(s []int) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	reverseSlice(s)
	fmt.Println(s)
}
