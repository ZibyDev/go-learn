### Задача 1 — Enum дней недели

Создай тип `Weekday int` и константы для дней недели (`Monday`, `Tuesday`, ... `Sunday`) через `iota`. Добавь метод `String()` который возвращает название дня. Выведи любой день — должно печататься название, а не число.

main.go
```go
package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (w Weekday) String() string {
	return [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}[w]
}

func main() {
	fmt.Println(Monday)
	fmt.Println(Friday)
}
```
