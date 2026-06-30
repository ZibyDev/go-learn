package main

import "fmt"

type Cell struct {
	Row, Col int
}

func main() {
	chess := map[Cell]string{
		{Row: 1, Col: 1}: "пешка",
		{Row: 4, Col: 3}: "конь",
		{Row: 2, Col: 1}: "ладья",
		{Row: 6, Col: 8}: "король",
	}
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if v, ok := chess[Cell{Row: i, Col: j}]; ok {
				fmt.Println(i, j, v)
			}
		}
	}
}
