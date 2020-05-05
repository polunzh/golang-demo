package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}
	for idx := range items {
		items[idx] *= 2
	}

	fmt.Println(items)
}
