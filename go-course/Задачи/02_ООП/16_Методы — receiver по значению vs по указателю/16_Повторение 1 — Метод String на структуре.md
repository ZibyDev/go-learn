### Повторение 1 — Метод String на структуре (уроки 5, 15, 16)

Объяви структуру `Duration{Hours, Minutes int}`. Напиши метод `String() string` (value receiver), который форматирует длительность как `"2ч 30м"`. Создай несколько Duration и выведи их через `fmt.Println` — метод должен вызваться автоматически. Проверь, что для `Duration{0, 45}` выводится `"0ч 45м"`.

> Связка: структуры (урок 15) + форматирование строк `fmt.Sprintf` (урок 5) + методы и автоматический вызов String() (урок 16). Вспомни метод `String()` из урока 3.

```go
package main

import "fmt"

type Duration struct {
	Hours, Minutes int
}

func (d Duration) String() string {
	return fmt.Sprintf("%dч %dм", d.Hours, d.Minutes)
}

func main() {
	d1 := Duration{Hours: 2, Minutes: 30}
	d2 := Duration{Hours: 0, Minutes: 45}
	d3 := Duration{Hours: 12, Minutes: 0}
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
}
```

```output
2ч 30м
0ч 45м
12ч 0м
```
