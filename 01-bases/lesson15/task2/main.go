package main

import "fmt"

type Counter struct {
	Value int
}

func incByValue(c Counter) {
	c.Value++
}

func incByPointer(c *Counter) {
	c.Value++
}

func main() {
	c := Counter{Value: 1}
	fmt.Println(c)
	incByValue(c)
	fmt.Println(c)
	incByPointer(&c)
	fmt.Println(c)
}
