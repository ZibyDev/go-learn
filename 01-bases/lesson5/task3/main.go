package main

import (
	"fmt"
)

func reverse(s string) string {
	var strRune = []rune(s)
	var n = len(strRune)
	for i := 0; i < n/2; i++ {
		strRune[i], strRune[n-i-1] = strRune[n-i-1], strRune[i]
	}
	return string(strRune)
}

func main() {
	var str1 = "Hello"
	var str2 = "Привет"
	fmt.Println(str1 + " " + reverse(str1))
	fmt.Println(str2 + " " + reverse(str2))
}
