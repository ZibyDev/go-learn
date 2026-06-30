package main

import "fmt"

func counter() (count int) {
	defer func() { count += 10 }()
	defer func() { count += 10 }()
	return 5
}

func main() {
	fmt.Println(counter())
}
