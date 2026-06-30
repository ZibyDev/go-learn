package main

import "fmt"

type Season int

const (
	Winter Season = iota
	Spring
	Summer
	Autumn
)

func (s Season) String() string {
	switch s {
	case 0:
		return "Winter"
	case 1:
		return "Spring"
	case 2:
		return "Summer"
	default:
		return "Autumn"
	}
}

func SeasonByNumber(x int) Season {
	switch x {
	case 1, 2, 12:
		return Winter
	case 3, 4, 5:
		return Spring
	case 6, 7, 8:
		return Summer
	default:
		return Autumn
	}
}

func main() {
	fmt.Println(SeasonByNumber(1))
	fmt.Println(SeasonByNumber(4))
	fmt.Println(SeasonByNumber(7))
	fmt.Println(SeasonByNumber(10))
}
