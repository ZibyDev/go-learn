package main

import "fmt"

type Reader interface{ Read() string }
type Writer interface{ Write(s string) }

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	data string
}

func (f *File) Read() string   { return f.data }
func (f *File) Write(s string) { f.data = s }

func useReadWriter(rw ReadWriter) {
	rw.Write("Hello")
	fmt.Println("Прочитано:", rw.Read())
}

func main() {
	f := &File{}
	useReadWriter(f)

	var r Reader = f
	var w Writer = f
	var rw ReadWriter = f
	fmt.Println(r.Read(), w, rw)
}
