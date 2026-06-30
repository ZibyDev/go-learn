package main

import "fmt"

func checkItem(inv map[string]int, item string) string {
	if count, ok := inv[item]; ok {
		if count <= 0 {
			return "закончился"
		}
		return fmt.Sprintf("в наличии: %d", count)
	}
	return "нет в наличии"
}

func main() {
	inv := map[string]int{
		"Меч":            5,
		"Зелье здоровья": 0,
	}
	fmt.Println(checkItem(inv, "Меч"))
	fmt.Println(checkItem(inv, "Зелье здоровья"))
	fmt.Println(checkItem(inv, "Щит"))
}
