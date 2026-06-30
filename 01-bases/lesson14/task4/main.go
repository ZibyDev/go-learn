package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"a": 4,
		"D": 2,
		"r": 1,
		"M": 5,
		"A": 3,
	}
	s := []string{}
	for k := range m {
		s = append(s, k)
	}
	sort.Strings(s)
	for _, i := range s {
		fmt.Printf("%s: %d\n", i, m[i])
	}
}
