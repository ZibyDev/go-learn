package main

import "fmt"

func divisors(n int) (count int, sum int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(divisors(12))
	fmt.Println(divisors(0))
	fmt.Println(divisors(2))
	fmt.Println(divisors(120))
	fmt.Println(divisors(43))
}
