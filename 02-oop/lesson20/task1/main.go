package main

import "fmt"

type DataSource interface {
	Fetch() string
}

type APISource struct{}
type MockSource struct{}

func (api APISource) Fetch() string   { return "данные из API" }
func (mock MockSource) Fetch() string { return "тестовые данные" }

type App struct {
	source DataSource
}

func (a App) Run() string { return a.source.Fetch() }

func main() {
	a1 := App{source: APISource{}}
	a2 := App{source: MockSource{}}
	fmt.Println(a1.Run())
	fmt.Println(a2.Run())
}
