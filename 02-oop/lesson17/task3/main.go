package main

import "fmt"

type Animal interface {
	Sound() string
	Legs() int
}

type Dog struct{}
type Cat struct{}
type Bird struct{}

func (a Dog) Sound() string  { return "Woof!" }
func (a Dog) Legs() int      { return 4 }
func (a Cat) Sound() string  { return "Meow!" }
func (a Cat) Legs() int      { return 4 }
func (a Bird) Sound() string { return "Bird sound!" }
func (a Bird) Legs() int     { return 2 }

func describe(a Animal) {
	fmt.Println(a.Sound(), a.Legs())
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Bird{}, Cat{}, Bird{}}
	for _, animal := range animals {
		describe(animal)
	}
}
