package main

import "fmt"

func applyDamage(health *int, damage int) {
	if damage > *health {
		*health = 0
	} else {
		*health -= damage
	}
}

func main() {
	health := 100
	fmt.Println(health)
	applyDamage(&health, 30)
	fmt.Println(health)
	applyDamage(&health, 50)
	fmt.Println(health)
	applyDamage(&health, 50)
	fmt.Println(health)
}
