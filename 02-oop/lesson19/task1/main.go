package main

import "fmt"

type Entity struct {
	ID   int
	Name string
}

func (e Entity) Info() string { return fmt.Sprintf("ID:%d Name:%s", e.ID, e.Name) }

type Player struct {
	Entity
	Score int
}

func main() {
	p := Player{Entity: Entity{ID: 1, Name: "Anton"}, Score: 100}
	fmt.Println(p.ID, p.Name)
	fmt.Println(p.Info())
}
