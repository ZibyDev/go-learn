package main

import (
	"fmt"
	"time"
)

func slowOperation() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Since(start))
}

func main() {
	slowOperation()
}
