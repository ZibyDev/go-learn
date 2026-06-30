### Задача 1 — Zero values

Объяви по одной переменной каждого из типов: `int`, `float64`, `bool`, `string`, указатель на `int` — **без инициализации**. Выведи их значения и типы через `fmt.Printf` с форматом `%T` (тип) и `%v` (значение).

Ожидаемый вывод:

```
int: 0
float64: 0
bool: false
string: 
*int: <nil>
```

main.go
```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var u *int

	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("%T %v\n", u, u)
}
```
