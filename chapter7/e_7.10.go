package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = filter(slice, even)
	fmt.Println(slice)
}

func even(value int) bool {
	return value%2 == 0
}

func filter(slice []int, fn func(int) bool) []int {
	var res []int
	for _, value := range slice {
		if fn(value) {
			res = append(res, value)
		}
	}

	return res
}
