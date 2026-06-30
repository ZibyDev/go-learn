### Задача 1 — FizzBuzz логика (одно число)

Напиши функцию `fizzbuzz(n int) string` которая возвращает:

- `"FizzBuzz"` если n делится на 3 и на 5
- `"Fizz"` если делится только на 3
- `"Buzz"` если делится только на 5
- строковое представление числа в остальных случаях (вспомни `strconv.Itoa` из урока 5)

Проверь на числах 3, 5, 15, 7.

> Используй обычные `if-else` или `switch` без выражения — на твой выбор.

main.go
```go
package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func main() {
	fmt.Println(fizzbuzz(3))
	fmt.Println(fizzbuzz(5))
	fmt.Println(fizzbuzz(15))
	fmt.Println(fizzbuzz(7))
}
```
