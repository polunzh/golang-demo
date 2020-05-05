package main

import "fmt"

func main() {
	const length = 16
	var arr [length]int

	for i := 0; i < length; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
}
