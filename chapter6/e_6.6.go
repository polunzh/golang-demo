package main

import "fmt"

func main() {
	fmt.Println(factorial1(4))
}

func factorial1(n int) (res int) {
	if n <= 1 {
		res = 1
		return
	}

	res = n * factorial1(n - 1)
	return
}
