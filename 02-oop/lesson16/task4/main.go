package main

import "fmt"

type Meters float64

func (m Meters) ToKilometers() float64  { return float64(m) / 1000 }
func (m Meters) ToCentimeters() float64 { return float64(m) * 1000 }
func (m Meters) String() string         { return fmt.Sprintf("%v м", float64(m)) }

func main() {
	fmt.Println(Meters(1500).ToKilometers())
	fmt.Println(Meters(1500).ToCentimeters())
	fmt.Println(Meters(1500))
}
