package main

import "fmt"

type Logger struct{ Prefix string }

func (l Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Prefix, msg)
}

type Counter struct{ count int }

func (c *Counter) Inc()       { c.count += 1 }
func (c *Counter) Count() int { return c.count }

type Worker struct {
	Logger
	*Counter
	Name string
}

func main() {
	w := Worker{
		Logger:  Logger{Prefix: "worker"},
		Counter: &Counter{},
		Name:    "worker",
	}

	w.Inc()
	w.Inc()
	w.Log(fmt.Sprintf("Count: %d", w.Count()))
}
