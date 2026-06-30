### Задача 2 — Type switch описатель

Напиши функцию `describe(x any) string`, которая через type switch возвращает описание: для `int` — `"целое N"`, для `float64` — `"дробное N"`, для `string` — `"строка длиной N"`, для `bool` — `"логическое N"`, для остального — `"неизвестно (тип T)"` (используй `%T`). Проверь на разных значениях.

```go
package main

import "fmt"

func describe(x any) string {
	switch v := x.(type) {
	case int:
		return fmt.Sprintf("Целое %d", v)
	case float64:
		return fmt.Sprintf("Дробное %f", v)
	case string:
		return fmt.Sprintf("Строка длиной %d", len(v))
	case bool:
		return fmt.Sprintf("Логическое %v", v)
	default:
		return fmt.Sprintf("Неизвестно (%T)", v)
	}
}

func main() {
	fmt.Println(describe(42))
	fmt.Println(describe(3.14))
	fmt.Println(describe("xdxdxd"))
	fmt.Println(describe(true))
	fmt.Println(describe('a'))
}
```

```bash
Целое 42
Дробное 3.140000
Строка длиной 6
Логическое true
Неизвестно (int32)
```
