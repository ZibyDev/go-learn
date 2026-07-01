package main

import "fmt"

type Animal interface {
	Sound() string
	Describe() string
}

type BaseAnimal struct {
	Name string
}

func (ba BaseAnimal) Describe() string { return ba.Name }

type Dog struct {
	BaseAnimal
}

func (d Dog) Sound() string { return "Woof-woof" }

type Cat struct {
	BaseAnimal
}

func (c Cat) Sound() string { return "Meooooow!" }

func main() {
	arr := []Animal{
		Dog{
			BaseAnimal{
				Name: "Bobik",
			},
		},
		Cat{
			BaseAnimal{
				Name: "Murka",
			},
		},
		Dog{
			BaseAnimal{
				Name: "Anton",
			},
		},
	}
	for _, x := range arr {
		fmt.Println(x.Describe())
		fmt.Println(x.Sound())
	}
}
