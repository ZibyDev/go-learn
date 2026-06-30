package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 = "42"
	var s2 = "abc"
	var s3 = "100"

	if num, err := strconv.Atoi(s1); err != nil {
		fmt.Printf("'%v' не число: %v\n", num, err)
	} else {
		fmt.Println(num)
	}
	if num, err := strconv.Atoi(s2); err != nil {
		fmt.Printf("'%v' не число: %v\n", num, err)
	} else {
		fmt.Println(num)
	}
	if num, err := strconv.Atoi(s3); err != nil {
		fmt.Printf("'%v' не число: %v\n", num, err)
	} else {
		fmt.Println(num)
	}
}
