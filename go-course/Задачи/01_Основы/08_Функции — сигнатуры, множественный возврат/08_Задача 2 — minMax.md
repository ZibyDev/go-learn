### Задача 2 — minMax

Напиши функцию `minMax(a, b, c int) (int, int)` которая возвращает минимум и максимум из трёх чисел. Без именованных возвратов — верни явно через `return min, max`. Проверь на `(5, 2, 8)` и `(3, 3, 3)`.

main.go
```go
package main

import "fmt"

func minMax(a, b, c int) (int, int) {
	var max, min = a, a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min, max
}

func main() {
	fmt.Println(minMax(5, 2, 8))
	fmt.Println(minMax(3, 3, 3))
}
```
