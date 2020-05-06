package main

import "fmt"

func main() {
	print(10)
}

func print(n int) {
	if n < 1 {
		return
	}

	fmt.Println(n)
	print(n - 1)
}