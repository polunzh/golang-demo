package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(InsertStringSlice(s1, s2, 3))
}

func InsertStringSlice(s1 []int, s2 []int, n int) []int {
	res := make([]int, len(s1)+len(s2))
	at := copy(res, s1[0:n])
	at += copy(res[at:], s2)
	copy(res[at:], s1[n:])
	return res
}
