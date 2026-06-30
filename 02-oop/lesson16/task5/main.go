package main

import "fmt"

type Stack struct{ items []int }

func (s *Stack) Push(x int) { s.items = append(s.items, x) }

func (s *Stack) Pop() (int, bool) {
	n := len(s.items)
	if n == 0 {
		return 0, false
	}
	pop := s.items[n-1]
	s.items = s.items[:n-1]
	return pop, true
}

func (s *Stack) Peek() (int, bool) {
	n := len(s.items)
	if n == 0 {
		return 0, false
	}
	return s.items[n-1], true
}

func (s *Stack) Len() int { return len(s.items) }

func main() {
	stack := Stack{items: []int{}}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
}
