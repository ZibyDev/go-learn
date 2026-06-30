package main

import "fmt"

type Base struct {
	ID   int
	Name string
}

type User struct {
	Base
	Email string
}

func main() {
	u1 := User{Base{ID: 1, Name: "John"}, "johm@mail.com"}
	fmt.Println(u1.ID)
	fmt.Println(u1.Name)
	fmt.Println(u1.Email)
}
