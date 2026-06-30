package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Qty   int
}

func totalValue(products []Product) int {
	sum := 0
	for _, x := range products {
		sum += x.Price * x.Qty
	}
	return sum
}

func groupByPrice(products []Product) map[int][]string {
	result := map[int][]string{}
	for _, x := range products {
		result[x.Price] = append(result[x.Price], x.Name)
	}
	return result
}

func main() {
	inv := []Product{
		{Name: "Зелье здоровья", Price: 5, Qty: 5},
		{Name: "Зелье маны", Price: 100, Qty: 1},
		{Name: "Щит", Price: 20, Qty: 1},
		{Name: "Меч", Price: 10, Qty: 5},
		{Name: "Топор", Price: 20, Qty: 10},
	}
	fmt.Println(totalValue(inv))
	fmt.Println(groupByPrice(inv))
}
