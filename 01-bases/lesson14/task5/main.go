package main

import "fmt"

func intersection(a, b []int) []int {
	inA := map[int]bool{}
	for _, x := range a {
		inA[x] = true
	}
	result := []int{}
	added := map[int]bool{}
	for _, x := range b {
		if !added[x] && inA[x] {
			result = append(result, x)
			added[x] = true
		}
	}
	return result
}

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 6}
	fmt.Println(intersection(s1, s2))
}
