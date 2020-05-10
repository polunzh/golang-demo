package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println("length of arr1:", len(arr1))
	fmt.Println("length of slice1:", len(slice1))
	fmt.Println("capacity of slice1:", cap(slice1))

	slice1 = slice1[0:4]
	fmt.Println("slice1:", slice1)
	fmt.Println("length of slice1:", len(slice1))
	fmt.Println("capacity of slice1:", cap(slice1))

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	var t = b[1:4]
	for i := 0; i < len(t); i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()

	t = b[:2]
	for i := 0; i < len(t); i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()

	t = b[2:]
	for i := 0; i < len(t); i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()

	t = b[:]
	for i := 0; i < len(t); i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}
