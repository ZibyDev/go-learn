package main

import "fmt"

func safeAddUint8(a, b uint8) (uint8, bool) {
	if a > 255-b {
		return 0, false
	}
	return a + b, true
}

func main() {
	fmt.Println(safeAddUint8(100, 50))
	fmt.Println(safeAddUint8(200, 100))
}
