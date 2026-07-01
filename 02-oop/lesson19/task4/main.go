package main

import "fmt"

type Mover struct{}

func (m Mover) Move() string { return "Move" }

type Attacker struct{}

func (a Attacker) Attack() string { return "Attack" }

type Warrior struct {
	Mover
	Attacker
	Name string
}

func main() {
	w := Warrior{Name: "Anton"}
	fmt.Println(w.Move())
	fmt.Println(w.Attack())
}
