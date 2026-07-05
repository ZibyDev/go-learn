package main

import "fmt"

type Reader interface {
	Read() string
}
type Writer interface {
	Write(s string)
}

type Buffer struct {
	s string
}

func (b *Buffer) Read() string   { return b.s }
func (b *Buffer) Write(s string) { b.s = s }

func readData(r Reader) {
	fmt.Println(r.Read())
}
func writeData(w Writer, s string) {
	w.Write(s)
}

func main() {
	b := Buffer{s: "default"}
	readData(&b)
	writeData(&b, "new msg")
	readData(&b)
}
