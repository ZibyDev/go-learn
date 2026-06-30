package main

import "fmt"

type Planet int

const (
	Mercury Planet = iota
	Venus
	Earth
	Mars
)

func (p Planet) String() string {
	switch p {
	case Mercury:
		return "Mercury"
	case Venus:
		return "Venus"
	case Earth:
		return "Earth"
	case Mars:
		return "Mars"
	}
	return "Not a planet"
}

func main() {
	planets := []fmt.Stringer{Mercury, Venus, Mars, Earth, Venus}
	for _, x := range planets {
		fmt.Println(x)
	}
}
