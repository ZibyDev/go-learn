package main

import "fmt"

func countVowels(s string) int {
	var countVow = 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			countVow++
		}
	}
	return countVow
}

func main() {
	fmt.Println(countVowels("Hello World"))
}
