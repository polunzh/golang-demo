package main

import "fmt"

func main() {
	s := []int{1, 100, 4, 0, 1, 2, 19, -1}
	fmt.Println(bubbleSort(s))
}

func bubbleSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}

	return s
}
