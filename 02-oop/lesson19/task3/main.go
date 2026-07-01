package main

import "fmt"

type Logger struct {
	Prefix string
}

func (l Logger) Log(msg string) { fmt.Printf("[%s] %s\n", l.Prefix, msg) }

type Service struct {
	Logger
	Name string
}

func main() {
	s := Service{Logger: Logger{Prefix: "Anton"}, Name: "Cool service"}
	s.Log(s.Name)
}
