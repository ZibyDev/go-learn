package main

import "fmt"

func sumNumbers(items []any) int {
	sum := 0
	for _, item := range items {
		if v, ok := item.(int); ok {
			sum += v
		}
	}
	return sum
}

func main() {
	fmt.Println(sumNumbers([]any{1, "two", 3, true, 5}))
}
