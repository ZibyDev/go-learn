package main

import "fmt"

func tree(height int) {
	for i := range height {
		for j := 0; j < height-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	tree(4)
	tree(7)
	tree(2)
}
