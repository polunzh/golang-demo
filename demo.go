package main

import "fmt"

func main() {
	items := []byte{10, 20, 30, 40, 50}
	fmt.Println(AppendByte(items, items...))
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	fmt.Println(slice)
	slice = slice[0:n]
	fmt.Println(slice)
	copy(slice[m:n], data)
	return slice
}
