### Задача 3 — go tool команды

Намеренно создай файл с неправильным форматированием (неровные отступы, лишние пробелы). Затем:

1. Запусти `go fmt` и посмотри как файл изменился
2. Добавь неиспользуемую переменную и убедись что `go vet` или компилятор об этом сообщает
3. Попробуй `go build` и `go run` — почувствуй разницу

Пункт 1
main.go до `go fmt`
```go
package main
import "fmt"

func main(){
fmt.Println("Hello, Go!")
}
```
main.go после `go fmt`
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

Пункт 2
Сообщение о неиспользуемой переменной:
```bash
zibynik@MacBook-Pro-Asin task3 % go vet ./...
# task3
# [task3]
vet: ./main.go:6:2: declared and not used: x
```

Пункт 3
`go build` - создает файл приложения
`go run` - сразу запускает программу