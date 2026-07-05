package main

import "fmt"

type Player struct {
	Name   string
	Health int
	Level  int
}

type Option func(*Player)

func WithHealth(h int) Option { return func(p *Player) { p.Health = h } }
func WithLevel(l int) Option  { return func(p *Player) { p.Level = l } }

func NewPlayer(name string, opts ...Option) *Player {
	p := Player{Name: name, Health: 100, Level: 1}
	for _, opt := range opts {
		opt(&p)
	}
	return &p
}

func main() {
	p1 := NewPlayer("Anton")
	p2 := NewPlayer("Igor", WithHealth(50))
	p3 := NewPlayer("Elena", WithHealth(70), WithLevel(3))
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
