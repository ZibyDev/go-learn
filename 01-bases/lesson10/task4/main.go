package main

import "fmt"

func memoizedSquare() func(int) int {
	cache := map[int]int{}
	return func(n int) int {
		v, ok := cache[n]
		if ok {
			return v
		}
		fmt.Println("вычисляю", n)
		res := n * n
		cache[n] = res
		return res
	}
}

func main() {
	f := memoizedSquare()
	fmt.Println(f(2))
	fmt.Println(f(4))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(7))
	fmt.Println(f(4))
}
