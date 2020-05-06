package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 4)
	for idx := range slice1 {
		slice1[idx] = idx
	}

	for idx, value := range slice1 {
		fmt.Printf("%d: %d\n", idx, value)
	}
}