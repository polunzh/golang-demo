package main

import "fmt"

func main() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func fibonacci() func() int {
	v1, v2 := 0, 1
	return func () int {
		v1, v2 = v2, v1 + v2

		return v2
	}
}