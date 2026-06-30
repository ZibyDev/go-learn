### Задача 4 — Pointer receiver и интерфейс

Объяви интерфейс `Accumulator` с методами `Add(x int)` и `Sum() int`. Реализуй тип `intAccumulator` с **pointer receiver** для `Add` (он меняет состояние). Покажи, что в интерфейс нужно класть **указатель** (`&intAccumulator{}`), а не значение. Попробуй положить значение и пойми ошибку компилятора.

> Связка с уроком 16: method set. Раз `Add` на `*T`, то интерфейс реализует только `*T`.

(1)
```go
package main

import "fmt"

type Accumulator interface {
	Add(x int)
	Sum() int
}

type intAccumulator []int

func (i *intAccumulator) Add(x int) { *i = append(*i, x) }
func (i *intAccumulator) Sum() int {
	total := 0
	for _, x := range *i {
		total += x
	}
	return total
}

func main() {
	var intAccumulator1 Accumulator = &intAccumulator{1, 2, 3, 4}
	intAccumulator1.Add(5)
	fmt.Println(intAccumulator1.Sum())
}
```

```bash
15
```

(2)
```go
package main

import "fmt"

type Accumulator interface {
	Add(x int)
	Sum() int
}

type intAccumulator []int

func (i *intAccumulator) Add(x int) { *i = append(*i, x) }
func (i *intAccumulator) Sum() int {
	total := 0
	for _, x := range *i {
		total += x
	}
	return total
}

func main() {
	var intAccumulator1 Accumulator = intAccumulator{1, 2, 3, 4}
	intAccumulator1.Add(5)
	fmt.Println(intAccumulator1.Sum())
}
```

```bash
# command-line-arguments
./main.go:22:36: cannot use intAccumulator{…} (value of slice type intAccumulator) as Accumulator value in variable declaration: intAccumulator does not implement Accumulator (method Add has pointer receiver)
```
