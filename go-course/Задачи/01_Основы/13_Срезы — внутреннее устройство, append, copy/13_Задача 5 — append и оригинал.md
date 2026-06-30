### Задача 5 — append и оригинал

Напиши две функции: `appendBad(s []int)` (просто `s = append(s, 100)` внутри) и `appendGood(s []int) []int` (возвращает результат append). В `main` покажи, что после `appendBad` оригинал не изменился, а `appendGood` (с присваиванием результата) — изменился. Объясни в комментарии почему.

```go
package main

import "fmt"

func appendBad(s []int) {
	s = append(s, 100)
}

func appendGood(s []int) []int {
	return append(s, 100)
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	appendBad(slice)
	fmt.Println(slice)
	slice = appendGood(slice)
	fmt.Println(slice)
}
```

Вывод:
```
[1 2 3]
[1 2 3]
[1 2 3 100]
```

Такой результат из-за того, что плохая функция создает внутри новых срез и не возвращает его, из-за чего оригинал не меняется, хорошая функция возвращает результат, который мы присваиваем нашему срезу