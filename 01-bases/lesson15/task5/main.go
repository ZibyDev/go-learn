package main

import "fmt"

type Timer struct {
	seconds int
}

func NewTimer(minutes int) *Timer {
	return &Timer{seconds: minutes * 60}
}

func main() {
	t1 := NewTimer(2)
	t2 := NewTimer(5)
	fmt.Printf("%+v\n", t1)
	fmt.Printf("%+v\n", t2)
}
