### Задача 1 — Swap

Напиши функцию `swap(a, b *int)` которая меняет местами значения двух переменных через указатели. Проверь: до вызова `x=1, y=2`, после — `x=2, y=1`.

```go
package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 1, 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
```

Всё работает
