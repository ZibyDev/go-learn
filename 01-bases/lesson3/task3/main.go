package main

import "fmt"

type Ability uint8

const (
	CanFly Ability = 1 << iota
	CanSwim
	CanClimb
	CanDig
)

func main() {
	character := CanFly | CanSwim

	fmt.Printf("Can fly: %t\n", character&CanFly != 0)
	fmt.Printf("Can swim: %t\n", character&CanSwim != 0)
	fmt.Printf("Can climb: %t\n", character&CanClimb != 0)
	fmt.Printf("Can dig: %t\n", character&CanDig != 0)

	character |= CanClimb
	fmt.Printf("After learning to climb:\n")
	fmt.Printf("Can fly: %t\n", character&CanFly != 0)
	fmt.Printf("Can swim: %t\n", character&CanSwim != 0)
	fmt.Printf("Can climb: %t\n", character&CanClimb != 0)
	fmt.Printf("Can dig: %t\n", character&CanDig != 0)

	character &^= CanFly
	fmt.Printf("After forgetting how to fly:\n")
	fmt.Printf("Can fly: %t\n", character&CanFly != 0)
	fmt.Printf("Can swim: %t\n", character&CanSwim != 0)
	fmt.Printf("Can climb: %t\n", character&CanClimb != 0)
	fmt.Printf("Can dig: %t\n", character&CanDig != 0)
}
