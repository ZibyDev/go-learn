### Задача 2 — MinMax массива

Напиши функцию `minMax(arr [6]int) (min, max int)` которая находит минимум и максимум в массиве. Инициализируй оба значения первым элементом, потом пройди по остальным. Проверь на `[6]int{5, 2, 8, 1, 9, 3}`.

```go
package main

import "fmt"

func minMax(arr [6]int) (min, max int) {
	min = arr[0]
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}

func main() {
	arr := [6]int{5, 2, 8, 1, 9, 3}
	min, max := minMax(arr)
	fmt.Println(min, max)
}
```
Вывод:
```
1 9
```
