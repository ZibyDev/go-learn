package main

import "fmt"

func average(nums ...float64) float64 {
	var length = len(nums)
	if length == 0 {
		return 0
	}
	var sum float64 = 0
	for _, r := range nums {
		sum += r
	}
	return sum / float64(length)
}

func main() {
	var list = []float64{1, 2, 3, 4, 5}
	fmt.Println(average(1, 2, 3, 4))
	fmt.Println(average(list...))
	fmt.Println(average())
}
