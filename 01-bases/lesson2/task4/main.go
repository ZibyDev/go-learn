package main

import "fmt"

func main() {
	name := "Hero"
	lvl := 1
	health := 100.0
	mana := 50.0
	alive := true
	x, y := 0, 0
	fmt.Printf("Имя: %s\nУровень: %d\nЗдоровье: %f\nМана: %f\nЖив: %t\nКоординаты: (%d, %d)\n", name, lvl, health, mana, alive, x, y)
}
