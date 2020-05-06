package main

import (
	"fmt"
)

func Append(slice, data []byte) []byte {
	s := make([]byte, len(slice)+len(data))
	copy(s, slice)

	for i := 0; i < len(data); i++ {
		s[len(slice) + i] = data[i]
	}

	return s
}

func main() {
	s := []byte{'a', 'b'}
	fmt.Printf("%s\n", Append(s, []byte{'c', 'd'}))
}
