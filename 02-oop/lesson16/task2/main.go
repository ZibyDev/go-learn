package main

import "fmt"

type Counter struct{ value int }

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() {
	c.value--
}

func (c Counter) Value() int {
	return c.value
}

func main() {
	c1 := Counter{value: 0}
	fmt.Println(c1.Value())
	c1.Increment()
	fmt.Println(c1.Value())
	c1.Increment()
	fmt.Println(c1.Value())
	c1.Decrement()
	fmt.Println(c1.Value())
}
