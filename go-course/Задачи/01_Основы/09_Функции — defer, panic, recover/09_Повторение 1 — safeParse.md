### Повторение 1 — Парсер + ошибки (уроки 6, 8, 9)

Напиши функцию `safeParse(s string) (n int, err error)` которая через `strconv.Atoi` парсит строку, но **оборачивает возможную панику** в ошибку через `defer`+`recover`. Хотя `Atoi` сам не паникует, сделай искусственную проверку: если строка пустая — вызови `panic("пустая строка")`, и перехвати её, вернув как `err`. Проверь на `"42"`, `""`, `"abc"`.

> Связка: именованные возвраты (урок 8) + recover (урок 9) + обработка ошибки парсинга (урок 6).

main.go
```go
package main

import (
	"fmt"
	"strconv"
)

func safeParse(s string) (n int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("поймана ошибка: %v", r)
		}
	}()
	if s == "" {
		panic("пустая строка")
	}
	return strconv.Atoi(s)
}

func main() {
	var s1 = "42"
	var s2 = ""
	var s3 = "abc"
	var n1, err1 = safeParse("42")
	var n2, err2 = safeParse("")
	var n3, err3 = safeParse("abc")
	fmt.Printf("Ваш ввод: \"%v\", Получено число: %v, Ошибка: %v\n", s1, n1, err1)
	fmt.Printf("Ваш ввод: \"%v\", Получено число: %v, Ошибка: %v\n", s2, n2, err2)
	fmt.Printf("Ваш ввод: \"%v\", Получено число: %v, Ошибка: %v\n", s3, n3, err3)
}
```
