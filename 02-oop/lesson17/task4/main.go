package main

import "fmt"

type Accumulator interface {
	Add(x int)
	Sum() int
}

type intAccumulator []int

func (i *intAccumulator) Add(x int) { *i = append(*i, x) }
func (i *intAccumulator) Sum() int {
	total := 0
	for _, x := range *i {
		total += x
	}
	return total
}

func main() {
	var intAccumulator1 Accumulator = &intAccumulator{1, 2, 3, 4}
	intAccumulator1.Add(5)
	fmt.Println(intAccumulator1.Sum())
}
