package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Go Привет"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}
