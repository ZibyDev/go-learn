package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Привет! Меня зовут Никита и я изучаю Go.")
	fmt.Printf("Версия Go: %s\n", runtime.Version())
}
