package main

import "fmt"

func main() {
	var s = make([]byte, 5)
	t := s[2:4]
	fmt.Println(len(s), cap(s), cap(t), len(t))
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s2)
	fmt.Println(s1)
	s2[1] = 't'
	fmt.Println(s1)
}
