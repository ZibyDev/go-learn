package main

import (
	"fmt"
	"unicode/utf8"
)

func groupByLength(words []string) map[int][]string {
	result := map[int][]string{}
	for _, i := range words {
		n := utf8.RuneCountInString(i)
		result[n] = append(result[n], i)
	}
	return result
}

func main() {
	slice := []string{"go", "cat", "do", "dog", "a"}
	fmt.Println(groupByLength(slice))
}
