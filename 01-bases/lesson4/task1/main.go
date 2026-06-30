package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 250
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d (max=%d)\n", x, math.MaxUint8)
		x++
	}
}
