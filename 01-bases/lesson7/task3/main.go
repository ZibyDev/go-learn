package main

import "fmt"

func isPalindrome(s string) bool {
	var strRune = []rune(s)
	var n = len(strRune)
	for i := 0; i < n/2; i++ {
		if strRune[i] != strRune[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("level"))
	fmt.Println(isPalindrome("hello"))
	fmt.Println(isPalindrome("Привет"))
}
