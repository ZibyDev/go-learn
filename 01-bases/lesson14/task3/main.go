package main

import "fmt"

func unique(nums []int) []int {
	check := map[int]bool{}
	result := []int{}
	for _, i := range nums {
		if !check[i] {
			check[i] = true
			result = append(result, i)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 2, 3, 1, 4, 3}
	fmt.Println(unique(slice))
}
