package main

import (
	"fmt"
)

func wordStats(s string) (vowels, consonants, digits int) {
	for _, r := range s {
		switch {
		case r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
			r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U':
			vowels++
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'):
			consonants++
		case r >= '0' && r <= '9':
			digits++
		}
	}
	return vowels, consonants, digits
}

func main() {
	fmt.Println(wordStats("Go123 rocks"))
}
