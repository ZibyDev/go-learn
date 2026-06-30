package main

import (
	"fmt"
	"strconv"
)

func safeParse(s string) (n int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("поймана ошибка: %v", r)
		}
	}()
	if s == "" {
		panic("пустая строка")
	}
	return strconv.Atoi(s)
}

func main() {
	var s1 = "42"
	var s2 = ""
	var s3 = "abc"
	var n1, err1 = safeParse("42")
	var n2, err2 = safeParse("")
	var n3, err3 = safeParse("abc")
	fmt.Printf("Ваш ввод: \"%v\", Получено число: %v, Ошибка: %v\n", s1, n1, err1)
	fmt.Printf("Ваш ввод: \"%v\", Получено число: %v, Ошибка: %v\n", s2, n2, err2)
	fmt.Printf("Ваш ввод: \"%v\", Получено число: %v, Ошибка: %v\n", s3, n3, err3)
}
