### Задача 3 — Среднее variadic

Напиши variadic функцию `average(nums ...float64) float64` которая возвращает среднее арифметическое любого количества чисел. Для пустого вызова `average()` верни `0` (не дай делению на ноль случиться). Проверь на `average(1, 2, 3, 4)` и на распаковке среза `average(data...)`.

> Подумай про граничный случай пустого вызова — `len(nums) == 0`.

main.go
```go
package main

import "fmt"

func average(nums ...float64) float64 {
	var length = len(nums)
	if length == 0 {
		return 0
	}
	var sum float64 = 0
	for _, r := range nums {
		sum += r
	}
	return sum / float64(length)
}

func main() {
	var list = []float64{1, 2, 3, 4, 5}
	fmt.Println(average(1, 2, 3, 4))
	fmt.Println(average(list...))
	fmt.Println(average())
}
```
