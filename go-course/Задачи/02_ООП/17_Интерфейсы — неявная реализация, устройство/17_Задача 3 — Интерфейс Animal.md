### Задача 3 — Интерфейс Animal

Объяви интерфейс `Animal` с двумя методами: `Sound() string` и `Legs() int`. Реализуй для `Dog`, `Cat`, `Bird` (у птицы 2 ноги, у остальных 4). Напиши функцию `describe(a Animal)`, печатающую звук и количество ног. Создай срез животных и опиши каждое в цикле.

```go
package main

import "fmt"

type Animal interface {
	Sound() string
	Legs() int
}

type Dog struct{}
type Cat struct{}
type Bird struct{}

func (a Dog) Sound() string  { return "Woof!" }
func (a Dog) Legs() int      { return 4 }
func (a Cat) Sound() string  { return "Meow!" }
func (a Cat) Legs() int      { return 4 }
func (a Bird) Sound() string { return "Bird sound!" }
func (a Bird) Legs() int     { return 2 }

func describe(a Animal) {
	fmt.Println(a.Sound(), a.Legs())
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Bird{}, Cat{}, Bird{}}
	for _, animal := range animals {
		describe(animal)
	}
}
```

```bash
Woof! 4
Meow! 4
Bird sound! 2
Meow! 4
Bird sound! 2
```
