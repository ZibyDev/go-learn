package main

import "fmt"

func grade(score int) string {
	switch {
	case score < 0 || score > 100:
		return "Invalid"
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func main() {
	fmt.Println(grade(95))
	fmt.Println(grade(83))
	fmt.Println(grade(71))
	fmt.Println(grade(65))
	fmt.Println(grade(40))
	fmt.Println(grade(-5))
	fmt.Println(grade(150))
}
