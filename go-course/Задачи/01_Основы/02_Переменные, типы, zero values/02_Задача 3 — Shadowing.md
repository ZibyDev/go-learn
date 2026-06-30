### Задача 3 — Shadowing

Напиши программу которая **намеренно** демонстрирует shadowing:

- Объяви переменную `x = 100` в `main`
- Внутри блока `if` объяви новую `x = 999`
- Выведи `x` внутри блока и после него
- Убедись что внешняя `x` не изменилась

main.go
```go
package main

import "fmt"

func main() {
	x := 100
	if true {
		x := 999
		fmt.Println(x)
	}
	fmt.Println(x)
}
```
увидел shadowing в действии