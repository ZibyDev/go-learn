### Задача 2 — comma ok

Создай карту-инвентарь `map[string]int` с парой товаров, где один товар имеет количество `0`. Напиши функцию `checkItem(inv map[string]int, item string) string` которая через comma ok возвращает: `"нет в наличии"` (ключа нет), `"закончился"` (ключ есть, значение 0), или `"в наличии: N"` (значение > 0). Проверь на трёх случаях.

```go
package main

import "fmt"

func checkItem(inv map[string]int, item string) string {
	if count, ok := inv[item]; ok {
		if count <= 0 {
			return "закончился"
		}
		return fmt.Sprintf("в наличии: %d", count)
	}
	return "нет в наличии"
}

func main() {
	inv := map[string]int{
		"Меч":            5,
		"Зелье здоровья": 0,
	}
	fmt.Println(checkItem(inv, "Меч"))
	fmt.Println(checkItem(inv, "Зелье здоровья"))
	fmt.Println(checkItem(inv, "Щит"))
}
```

```
в наличии: 5
закончился
нет в наличии
```
