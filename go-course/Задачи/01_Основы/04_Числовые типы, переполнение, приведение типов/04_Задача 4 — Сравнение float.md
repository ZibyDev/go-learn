### Задача 4 — Сравнение float

Напиши функцию `nearlyEqual(a, b, epsilon float64) bool`. Проверь её на:

- `0.1 + 0.2` и `0.3` (должны быть равны с epsilon `1e-9`)
- `1.0` и `1.5` (не равны)

Выведи также результат прямого сравнения `==` для контраста.

main.go
```go
package main

import (
	"fmt"
	"math"
)

func nearlyEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func main() {
	epsilon := 1e-9
	a := 0.1 + 0.2
	b := 0.3
	fmt.Println(a == b)
	fmt.Println(nearlyEqual(a, b, epsilon))

	c := 1.0
	d := 1.5
	fmt.Println(c == d)
	fmt.Println(nearlyEqual(c, d, epsilon))
}
```
Примечание: при запуске программы и сравнении a == b вывело значение true