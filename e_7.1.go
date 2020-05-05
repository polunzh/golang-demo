package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3}
	var arr2 = arr1
	var arr3 = &arr1
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr3[2] = 12
	fmt.Println(arr1)
	fmt.Println(arr2)
}
