### Повторение 1 — Функции + циклы + строки (уроки 5, 7, 8)

Напиши функцию `wordStats(s string) (vowels, consonants, digits int)` с именованными возвратами, которая за один проход по строке считает количество гласных, согласных (латинских букв) и цифр. Используй `for range` и `switch`. Проверь на `"Go123 rocks"`.

> Связка: руны из урока 5, циклы из урока 7, switch из урока 6, именованные возвраты из урока 8. Для проверки "буква/цифра" можно сравнивать руну с диапазонами: `r >= '0' && r <= '9'`.

main.go
```go
package main

import (
	"fmt"
)

func wordStats(s string) (vowels, consonants, digits int) {
	for _, r := range s {
		switch {
		case r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
			r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U':
			vowels++
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'):
			consonants++
		case r >= '0' && r <= '9':
			digits++
		}
	}
	return vowels, consonants, digits
}

func main() {
	fmt.Println(wordStats("Go123 rocks"))
}
```
