package main

import "fmt"

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	result = append(result, a...)
	result = append(result, b...)
	return result
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := merge(s1, s2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s4 := make([]int, len(s1))
	copy(s4, s1)
	s4 = append(s4, s2...)
	fmt.Println(s1)
	fmt.Println(s4)
	s4[1] = 10
	fmt.Println(s1)
	fmt.Println(s4)
}
