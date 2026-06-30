### Задача 2 — Размеры через iota

Создай константы для единиц измерения скорости передачи данных через `iota` и битовые сдвиги:

```
Bps  = 1
Kbps = 1000
Mbps = 1000000
Gbps = 1000000000
```

> Подсказка: здесь степень 10, а не 2. Подумай как использовать iota с умножением.

main.go
```go
package main

import "fmt"

const factor = 1000

const (
	Bps  = 1
	Kbps = Bps * factor
	Mbps = Kbps * factor
	Gbps = Mbps * factor
)

func main() {
	fmt.Println(Bps)
	fmt.Println(Kbps)
	fmt.Println(Mbps)
	fmt.Println(Gbps)
}
```

Этого сделать невозможно через iota, так как мы работаем с степень, а степень в константу через math записать нельзя