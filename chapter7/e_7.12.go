package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(RemoveStringSlice(s, 1, 3))
}

func RemoveStringSlice(s []int, start int, end int) []int {
	res := make([]int, len(s)-(end-start))
	at := copy(res, s[0:start])
	copy(res[at:], s[end:])
	return res
}
