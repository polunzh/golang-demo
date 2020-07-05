package main

import "fmt"

type Rectangle struct {
	x, y int
}

func area(r *Rectangle) int {
	return r.x * r.y	
}

func perimeter(r *Rectangle) int {
	return 2 * (r.x + r.y)
}

func main() {
	var r = new(Rectangle)
	r.x = 10
	r.y = 50
	fmt.Printf("area: %d\n", area(r))
	fmt.Printf("perimeter: %d\n", perimeter(r))
}