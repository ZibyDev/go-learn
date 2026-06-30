package main

import "fmt"

func reverseArray(arr *[5]int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func reverseCopy(arr [5]int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	reverseArray(&a)
	fmt.Println(a)
	reverseCopy(a)
	fmt.Println(a)
}
