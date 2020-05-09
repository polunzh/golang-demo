package main

import "fmt"

func main() {
	str := "zhangzhenqiang"
	fmt.Println(len(str)/2, len(str)/2)
	fmt.Println(str[len(str)/2:] + str[:len(str)/2])
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
