package main

import "fmt"

type Describable interface {
	Describe() string
}

type Product struct {
	Name string
}

func (p Product) Describe() string {
	return fmt.Sprintf("Продукт: %q", p.Name)
}

func tryDescribe(x any) {
	if v, ok := x.(Describable); ok {
		fmt.Println(v.Describe())
	} else {
		fmt.Println("не описываемо")
	}
}

func main() {
	tryDescribe(Product{Name: "Anton"})
	tryDescribe(42)
}
