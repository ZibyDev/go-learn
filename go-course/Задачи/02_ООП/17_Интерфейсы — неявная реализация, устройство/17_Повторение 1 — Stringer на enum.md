### Повторение 1 — Stringer на enum (уроки 3, 16, 17)

Объяви тип `Planet int` с константами `Mercury, Venus, Earth, Mars` через `iota` (урок 3). Реализуй метод `String() string`, возвращающий название планеты. Создай срез `[]fmt.Stringer` и положи туда несколько планет. Пройди по срезу и выведи каждую через `fmt.Println` — должны печататься названия.

> Связка: iota-enum (урок 3) + метод String() (урок 16) + интерфейс Stringer и полиморфный срез (урок 17). Это показывает, что `fmt.Stringer` — обычный интерфейс, который твой тип реализует неявно.

```go
package main

import "fmt"

type Planet int

const (
	Mercury Planet = iota
	Venus
	Earth
	Mars
)

func (p Planet) String() string {
	switch p {
	case Mercury:
		return "Mercury"
	case Venus:
		return "Venus"
	case Earth:
		return "Earth"
	case Mars:
		return "Mars"
	}
	return "Not a planet"
}

func main() {
	planets := []fmt.Stringer{Mercury, Venus, Mars, Earth, Venus}
	for _, x := range planets {
		fmt.Println(x)
	}
}
```

```bash
Mercury
Venus
Mars
Earth
Venus
```
