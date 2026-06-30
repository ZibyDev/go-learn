package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s3 := make([]int, 3)
	copy(s3, s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s2[0] = 10
	s3[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
