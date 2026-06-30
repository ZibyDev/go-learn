### Задача 2 — Таймер через defer

Напиши функцию  которая засекает время выполнения через `defer`. Используй `time.Now()` в начале и `defer` с замыканием, которое выведет `time.Since(start)` в конце. Внутри сделай задержку через `time.Sleep(100 * time.Millisecond)`.

> Подсказка: `start := time.Now()` в начале, затем `defer func() { fmt.Println(time.Since(start)) }()`. Подумай, почему здесь нужно именно замыкание, а не `defer fmt.Println(time.Since(start))`.

main.go
```go
package main

import (
	"fmt"
	"time"
)

func slowOperation() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(start))
}

func main() {
	slowOperation()
}
```

Если обернуть вывод в defer func(), то вывод будет такой:
```
101.106084ms
101.20125ms
```
Если не обернуть вывод fmt в defer func(), то вывод будет такой:
```
101.045667ms
17.333µs
```
