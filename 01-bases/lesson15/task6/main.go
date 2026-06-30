package main

import "fmt"

type Inventory struct {
	Items []string
}

func main() {
	inv1 := Inventory{Items: []string{"Меч", "Щит", "Топор"}}
	inv2 := inv1
	inv2.Items[1] = "Зелье"
	fmt.Println(inv2)
	fmt.Println(inv1)

	inv3 := Inventory{Items: make([]string, len(inv1.Items))}
	copy(inv3.Items, inv1.Items)
	inv3.Items[1] = "Щит"
	fmt.Println(inv3)
	fmt.Println(inv1)
}
