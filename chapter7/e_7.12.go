package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(RemoveStringSlice(s, 1, 3))
}

func splitStr(s string, i int) []string {
	slice := byte[](s)
}
