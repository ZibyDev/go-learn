package main

import "fmt"

func wordFrequency(words []string) map[string]int {
	result := map[string]int{}
	for _, i := range words {
		result[i]++
	}
	return result
}

func main() {
	slice := []string{"a", "b", "a", "c", "a", "b"}
	fmt.Println(wordFrequency(slice))
}
