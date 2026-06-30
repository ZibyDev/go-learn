package main

import "fmt"

func safeAccess(arr []int, index int) (value int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("индекс %d вне границ", index)
		}
	}()
	return arr[index], nil
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(safeAccess(arr, 0))
	fmt.Println(safeAccess(arr, 4))
	fmt.Println(safeAccess(arr, 5))
	fmt.Println(safeAccess(arr, -1))
}
