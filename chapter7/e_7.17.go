package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mapFunc(multi, s))
}

func multi(n int) int {
	return n * 10
}

func mapFunc(f func(int) int, s []int) []int {
	for i := 0; i < len(s); i++ {
		s[i] = f(s[i])
	}

	return s
}
