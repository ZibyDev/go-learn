### Задача 4 — defer меняет возврат

Напиши функцию `counter() (count int)` с именованным возвратом, которая:

- внутри `return 5`
- но имеет два `defer`, каждый из которых увеличивает `count` на 10
- предскажи результат и объясни порядок (LIFO) и как defer видит именованный возврат

Проверь — результат должен быть `25`.

main.go
```go
package main

import "fmt"

func counter() (count int) {
	defer func() { count += 10 }()
	defer func() { count += 10 }()
	return 5
}

func main() {
	fmt.Println(counter())
}
```

Результат 25, сначала работает 5, потом 2 defer, делая значение count = 15, потом 1 defer делает значение count = 25