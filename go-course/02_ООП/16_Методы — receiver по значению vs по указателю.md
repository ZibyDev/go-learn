# Урок 16 — Методы: receiver по значению vs по указателю

**Блок:** [[Go — План курса#📁 Блок 02 — ООП по-голански (~40 ч)|02 ООП по-голански]] **Время:** 4–5 часов **Теги:** #go/ооп #go/методы #go/структуры

---

## 📌 Теория

### 16.1 Что такое метод

Метод — это функция, привязанная к определённому типу. От обычной функции отличается **получателем** (receiver) — особым параметром перед именем метода, который указывает, к какому типу метод относится.

```go
type Rectangle struct {
    Width, Height float64
}

// Метод area с receiver типа Rectangle
func (r Rectangle) area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 3, Height: 4}
    fmt.Println(rect.area())  // 12 — вызов через точку
}
```

Разбор синтаксиса: `func (r Rectangle) area() float64`

- `(r Rectangle)` — **receiver**: `r` это имя, `Rectangle` это тип
- `area` — имя метода
- вызывается как `rect.area()`, а не `area(rect)`

> 💡 В Go нет классов, но методы на типах дают похожий результат: данные (структура) + поведение (методы на ней). Это и есть «ООП по-голангски» — без классов и наследования.

---

### 16.2 Зачем методы, если есть функции

Сравни два подхода:

```go
// Функция
func area(r Rectangle) float64 {
    return r.Width * r.Height
}
area(rect)        // вызов функции

// Метод
func (r Rectangle) area() float64 {
    return r.Width * r.Height
}
rect.area()       // вызов метода
```

Разница не только синтаксическая:

- Методы **группируют поведение** вокруг типа (читается как «у прямоугольника есть площадь»)
- Методы нужны для реализации **интерфейсов** (урок 17) — это главная причина
- Методы дают тип-ориентированный API: `rect.area()`, `circle.area()` — единообразно

---

### 16.3 Receiver по значению

Когда receiver передаётся **по значению**, метод работает с **копией**. Изменения внутри метода не влияют на оригинал:

```go
type Counter struct {
    count int
}

func (c Counter) Increment() {
    c.count++   // меняется КОПИЯ, оригинал не тронут
}

func main() {
    c := Counter{count: 0}
    c.Increment()
    fmt.Println(c.count)  // 0 — не изменилось!
}
```

> ⚠️ Это та же логика, что с функциями (урок 8) и структурами (урок 15): значение копируется. Если метод должен **изменить** структуру — нужен receiver по указателю.

---

### 16.4 Receiver по указателю

Receiver-указатель (`*T`) позволяет методу изменять оригинал:

```go
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++   // меняем ОРИГИНАЛ через указатель
}

func main() {
    c := Counter{count: 0}
    c.Increment()
    c.Increment()
    fmt.Println(c.count)  // 2 — изменилось!
}
```

> 💡 Обрати внимание: вызов `c.Increment()` выглядит одинаково для обоих receiver'ов. Go **автоматически** берёт адрес (`&c`) при вызове метода с pointer receiver на адресуемом значении. Не нужно писать `(&c).Increment()`.

---

### 16.5 Когда какой receiver — правила выбора

Это один из самых частых вопросов новичков. Вот практические правила:

**Используй pointer receiver (`*T`), если:**

1. Метод **изменяет** receiver (мутирует поля)
2. Структура **большая** (копировать дорого)
3. Для **консистентности** — если хоть один метод типа использует pointer receiver, делай все остальные тоже pointer (смешивать не рекомендуется)

**Используй value receiver (`T`), если:**

1. Метод только **читает** данные, не меняет
2. Структура **маленькая** (несколько полей примитивов)
3. Тип «неизменяемый» по смыслу (например, `time.Time`)

```go
type Point struct{ X, Y int }

// Читающий метод — value receiver ок (маленькая структура)
func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Изменяющий метод — pointer receiver обязателен
func (p *Point) Move(dx, dy int) {
    p.X += dx
    p.Y += dy
}
```

> 💡 **Практическое правило по умолчанию:** если сомневаешься — бери pointer receiver. Это безопаснее (изменения работают) и эффективнее (нет копирования). Многие команды используют pointer receiver почти везде для консистентности.

---

### 16.6 Автоматическое преобразование при вызове

Go удобно сглаживает разницу между значением и указателем при **вызове** метода:

```go
type Counter struct{ count int }
func (c *Counter) Inc()  { c.count++ }
func (c Counter) Get() int { return c.count }

func main() {
    c := Counter{}        // значение
    c.Inc()               // Go сам берёт &c → (&c).Inc()
    fmt.Println(c.Get())  // 1

    p := &Counter{}       // указатель
    p.Inc()               // напрямую
    fmt.Println(p.Get())  // Go сам разыменует → (*p).Get()
}
```

> 💡 Правило: для **адресуемого** значения Go автоматически конвертирует между `T` и `*T` при вызове метода. И value-метод можно звать на указателе, и pointer-метод на значении (если оно адресуемо).

> ⚠️ Исключение: значение в карте **не адресуемо**, поэтому pointer-метод на нём не вызвать напрямую:

```go
m := map[string]Counter{"a": {}}
// m["a"].Inc()  // ❌ нельзя — m["a"] не адресуемо
```

---

### 16.7 Методы можно определять не только на структурах

Метод можно привязать к **любому** именованному типу, объявленному в твоём пакете — не только к структуре:

```go
type Celsius float64

func (c Celsius) ToFahrenheit() Celsius {
    return c*9/5 + 32
}

func main() {
    temp := Celsius(25)
    fmt.Println(temp.ToFahrenheit())  // 77
}
```

Ты уже это видел в уроке 3 — метод `String()` на типе-enum (`Weekday`, `Direction`):

```go
type Weekday int
func (w Weekday) String() string { ... }
```

> ⚠️ Нельзя определять методы на типах из **чужих** пакетов (например, нельзя добавить метод к `int` или к `time.Time`). Только на типах, объявленных в твоём пакете. Обходится через создание своего типа: `type MyInt int`.

---

### 16.8 Метод String() — интерфейс Stringer (мостик к уроку 17)

Если у типа есть метод `String() string`, то `fmt.Println` и `fmt.Printf("%v")` автоматически его используют:

```go
type Color struct {
    R, G, B uint8
}

func (c Color) String() string {
    return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func main() {
    c := Color{255, 128, 0}
    fmt.Println(c)  // #FF8000 — вызвался String() автоматически!
}
```

> 💡 Это работает потому, что `Color` неявно реализует интерфейс `fmt.Stringer` (тип с методом `String() string`). Как именно это происходит — главная тема урока 17 (интерфейсы). Здесь просто увидь связь: метод + автоматическое использование.

---

### 16.9 Набор методов (method set)

У типа есть «набор методов» — это важно для интерфейсов (урок 17):

- Тип `T` (значение): в наборе методы **только с value receiver**
- Тип `*T` (указатель): в наборе **и value, и pointer** receiver методы

```go
type T struct{}
func (t T) ValueMethod()   {}
func (t *T) PointerMethod() {}

// У значения T: ValueMethod
// У указателя *T: ValueMethod И PointerMethod
```

> 💡 Пока запомни общий принцип: указатель `*T` имеет доступ ко всем методам, а значение `T` — только к value-методам. Это станет важным при реализации интерфейсов в уроке 17.

---

## 💡 Примеры

### Пример 1 — Value vs pointer receiver

```go
package main

import "fmt"

type BankAccount struct {
    balance int
}

// value receiver — только читает
func (a BankAccount) Balance() int {
    return a.balance
}

// pointer receiver — изменяет
func (a *BankAccount) Deposit(amount int) {
    a.balance += amount
}

func main() {
    acc := BankAccount{balance: 100}
    acc.Deposit(50)
    fmt.Println(acc.Balance())  // 150
}
```

### Пример 2 — Методы на не-структурном типе

```go
package main

import "fmt"

type Temperature float64

func (t Temperature) IsFreezing() bool {
    return t <= 0
}

func (t Temperature) String() string {
    return fmt.Sprintf("%.1f°C", float64(t))
}

func main() {
    t := Temperature(-5)
    fmt.Println(t)            // -5.0°C  (через String())
    fmt.Println(t.IsFreezing()) // true
}
```

### Пример 3 — Цепочка методов (через возврат receiver)

```go
package main

import "fmt"

type StringBuilder struct {
    parts []string
}

func (b *StringBuilder) Add(s string) *StringBuilder {
    b.parts = append(b.parts, s)
    return b   // возвращаем себя для цепочки
}

func (b *StringBuilder) Build() string {
    result := ""
    for _, p := range b.parts {
        result += p
    }
    return result
}

func main() {
    b := &StringBuilder{}
    result := b.Add("Hello").Add(", ").Add("Go").Build()  // цепочка
    fmt.Println(result)  // Hello, Go
}
```

### Пример 4 — Метод изменяет через указатель

```go
package main

import "fmt"

type Player struct {
    Name   string
    Health int
}

func (p *Player) TakeDamage(dmg int) {
    p.Health -= dmg
    if p.Health < 0 {
        p.Health = 0
    }
}

func (p Player) IsAlive() bool {
    return p.Health > 0
}

func main() {
    hero := Player{Name: "Hero", Health: 100}
    hero.TakeDamage(30)
    fmt.Println(hero.Health)   // 70
    hero.TakeDamage(80)
    fmt.Println(hero.Health)   // 0
    fmt.Println(hero.IsAlive()) // false
}
```

---

## ⚠️ Подводные камни

### 1. Value receiver не меняет оригинал

```go
func (c Counter) Inc() { c.count++ }  // меняет копию
// c.count останется прежним
```

### 2. Смешивание value и pointer receiver у одного типа

```go
func (t T) Read()  {}   // value
func (t *T) Write() {}  // pointer
// Технически можно, но не рекомендуется — выбери одно для всего типа
```

### 3. Pointer-метод на неадресуемом значении

```go
m := map[string]Counter{"a": {}}
// m["a"].Inc()  // ❌ значение в карте неадресуемо
// Решение: достать, изменить, положить обратно
```

### 4. Нельзя добавить метод к чужому типу

```go
// func (i int) Double() int {}  // ❌ int из другого пакета
type MyInt int
func (i MyInt) Double() MyInt { return i * 2 }  // ✅ свой тип
```

### 5. nil pointer receiver

```go
type T struct{ x int }
func (t *T) Get() int { return t.x }
var p *T
// p.Get()  // panic: nil pointer dereference
```

---

## ✅ Задачи

- [x] [[Задачи/02_ООП/16_Методы — receiver по значению vs по указателю/16_Повторение 1 — Метод String на структуре]]
- [x] [[Задачи/02_ООП/16_Методы — receiver по значению vs по указателю/16_Задача 1 — Площадь и периметр]]
- [x] [[Задачи/02_ООП/16_Методы — receiver по значению vs по указателю/16_Задача 2 — Счётчик с указателем]]
- [x] [[Задачи/02_ООП/16_Методы — receiver по значению vs по указателю/16_Задача 3 — Игровой персонаж]]
- [x] [[Задачи/02_ООП/16_Методы — receiver по значению vs по указателю/16_Задача 4 — Метод на своём типе]]
- [x] [[Задачи/02_ООП/16_Методы — receiver по значению vs по указателю/16_Задача 5 — Стек]]

---

## 🔄 Задачи на повторение (условия)

### Повторение 1 — Метод String на структуре (уроки 5, 15, 16)

Объяви структуру `Duration{Hours, Minutes int}`. Напиши метод `String() string` (value receiver), который форматирует длительность как `"2ч 30м"`. Создай несколько Duration и выведи их через `fmt.Println` — метод должен вызваться автоматически. Проверь, что для `Duration{0, 45}` выводится `"0ч 45м"`.

> Связка: структуры (урок 15) + форматирование строк `fmt.Sprintf` (урок 5) + методы и автоматический вызов String() (урок 16). Вспомни метод `String()` из урока 3.

---

## 📝 Задачи (условия)

### Задача 1 — Площадь и периметр

Объяви структуру `Circle{Radius float64}`. Напиши два метода с **value receiver**: `Area() float64` (π·r²) и `Circumference() float64` (2·π·r). Используй `math.Pi`. Создай круг и выведи обе характеристики.

> Методы только читают — подумай, какой receiver подходит.

---

### Задача 2 — Счётчик с указателем

Объяви структуру `Counter{value int}`. Напиши методы `Increment()`, `Decrement()` и `Value() int`. Подумай, какие методы должны иметь pointer receiver, а какой — value. Создай счётчик, несколько раз увеличь/уменьши, выведи итог.

---

### Задача 3 — Игровой персонаж

Объяви структуру `Hero{Name string; HP int; MaxHP int}`. Напиши методы:

- `TakeDamage(dmg int)` — уменьшает HP, не ниже 0 (pointer receiver)
- `Heal(amount int)` — увеличивает HP, не выше MaxHP (pointer receiver)
- `IsAlive() bool` — HP > 0 (value receiver)
- `Status() string` — возвращает `"Name: HP/MaxHP"` (value receiver)

Создай героя, нанеси урон, полечи, выведи статус на каждом шаге.

---

### Задача 4 — Метод на своём типе

Объяви тип `Meters float64` (не структуру!). Напиши методы `ToKilometers() float64` и `ToCentimeters() float64`. Также добавь `String() string`, форматирующий как `"1500 м"`. Проверь на значении `Meters(1500)`.

> Это показывает, что методы работают не только на структурах. Вспомни `Celsius`/`Temperature` из теории.

---

### Задача 5 — Стек

Объяви структуру `Stack{items []int}`. Напиши методы (подумай про receiver для каждого):

- `Push(x int)` — добавить элемент
- `Pop() (int, bool)` — снять верхний (вернуть значение и `ok`; если пусто — `0, false`)
- `Peek() (int, bool)` — посмотреть верхний без снятия
- `Len() int` — количество элементов

Проверь: push нескольких чисел, peek, pop до опустошения, pop из пустого стека.

> Связка с уроком 13: стек на срезе. `Push` это `append`, `Pop` это срез `s[:len-1]` + возврат последнего. Какие методы меняют стек (pointer), а какие читают (value)?

---

## 📋 Краткое резюме

- Метод — функция с получателем (receiver): `func (r Type) Name() {...}`, вызов через точку `obj.Name()`.
- В Go нет классов: данные (структура) + методы на ней заменяют ООП-классы.
- **Value receiver** (`T`) — работает с копией, изменения не видны снаружи; для читающих методов и маленьких типов.
- **Pointer receiver** (`*T`) — работает с оригиналом, может менять; для мутирующих методов и больших структур.
- При сомнении бери pointer receiver; не смешивай value и pointer у одного типа (консистентность).
- Go автоматически конвертирует между `T` и `*T` при вызове на адресуемом значении.
- Значение в карте неадресуемо — pointer-метод на нём напрямую не вызвать.
- Методы можно определять на любом своём именованном типе, не только на структурах; на чужих типах — нельзя.
- Метод `String() string` автоматически используется `fmt` (интерфейс Stringer — мостик к уроку 17).
- Набор методов: у `T` только value-методы, у `*T` — и value, и pointer (важно для интерфейсов).

---

## 🔗 Связанные темы

- [[01_Основы/15_Структуры — поля, теги, анонимные поля]] — предыдущий урок (структуры)
- [[02_ООП/17_Интерфейсы — неявная реализация, устройство]] — следующий урок (методы → интерфейсы)
- [[01_Основы/11_Указатели — как работают, когда использовать]] — pointer receiver
- [[01_Основы/03_Константы, iota, типизированные константы]] — метод String() на enum
- [[02_ООП/19_Встраивание (embedding) вместо наследования]] — методы и встраивание

---

#go/ооп #go/методы #go/структуры #урок