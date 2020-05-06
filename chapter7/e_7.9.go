package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = enlarge(slice, 6)
	fmt.Println(len(slice))
}

func enlarge(slice []int, factor int) []int {
	newSlice := make([]int, len(slice)+factor)
	copy(newSlice, slice)
	slice = newSlice
	return slice
}
