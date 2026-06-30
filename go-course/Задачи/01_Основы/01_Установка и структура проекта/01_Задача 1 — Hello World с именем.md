### Задача 1 — Hello World с именем

Создай новый Go модуль. Напиши программу которая выводит:

```
Привет! Меня зовут <твоё имя> и я изучаю Go.
Версия Go: <версия>
```

Версию Go получи программно через пакет `runtime` (`runtime.Version()`).

main.go
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Привет! Меня зовут Никита и я изучаю Go.")
	fmt.Println("Версия Go: ", runtime.Version())
}
```
