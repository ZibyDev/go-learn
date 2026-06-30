package main

import (
	"fmt"
	"strconv"
)

func parsePositive(s string) (int, error) {
	if num, err := strconv.Atoi(s); err != nil {
		return 0, err
	} else if num < 0 {
		return 0, fmt.Errorf("число должно быть положительным, получено %d", num)
	} else {
		return num, nil
	}
}

func main() {
	fmt.Println(parsePositive("42"))
	fmt.Println(parsePositive("-5"))
	fmt.Println(parsePositive("abc"))
}
