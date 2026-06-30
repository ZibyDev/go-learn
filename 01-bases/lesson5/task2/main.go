package main

import (
	"fmt"
)

func main() {
	str := "Hi Мир"
	for index, r := range str {
		fmt.Printf("Index: %v, %c (код руны: %d)\n", index, r, r)
	}
}
