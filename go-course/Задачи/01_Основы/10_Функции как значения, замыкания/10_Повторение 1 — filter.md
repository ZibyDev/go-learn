### Повторение 1 — Замыкание + срез (уроки 8, 10)

Напиши функцию `filter(nums []int, pred func(int) bool) []int` которая возвращает новый срез только из тех элементов, для которых `pred` вернул `true`. Затем с её помощью отфильтруй чётные числа из `[]int{1,2,3,4,5,6}`, передав анонимную функцию-предикат.

> Связка: функция как аргумент (урок 10) + работа со срезом через append (срез базово, подробно в уроке 13). Внутри используй `result := []int{}` и `append(result, n)`.

main.go
```go
package main

import "fmt"

func filter(nums []int, pred func(int) bool) []int {
	result := []int{}
	for _, num := range nums {
		if pred(num) {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(filter([]int{1, 2, 3, 4, 5, 6}, func(x int) bool { return x%2 == 0 }))
}
```
