package main

import "fmt"

func main() {
	x := 100
	if true {
		x := 999
		fmt.Println(x)
	}
	fmt.Println(x)
}
