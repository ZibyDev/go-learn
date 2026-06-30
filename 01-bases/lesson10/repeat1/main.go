package main

import "fmt"

func filter(nums []int, pred func(int) bool) []int {
	result := []int{}
	for _, num := range nums {
		if pred(num) {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(filter([]int{1, 2, 3, 4, 5, 6}, func(x int) bool { return x%2 == 0 }))
}
