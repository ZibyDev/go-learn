### Задача 1 — Безопасный assertion

Напиши функцию `toInt(x any) (int, bool)`, которая через «comma ok» assertion пытается извлечь `int` из `any`. Верни значение и флаг успеха. Проверь на `42`, `"hello"`, `3.14` — выведи результат для каждого.

```go
package main

import "fmt"

func toInt(x any) (int, bool) {
	v, ok := x.(int)
	return v, ok
}

func main() {
	fmt.Println(toInt(42))
	fmt.Println(toInt("hello"))
	fmt.Println(toInt(3.14))
}
```

```bash
42 true
0 false
0 false
```
