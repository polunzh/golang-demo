package main

import "fmt"

type Interval struct {
	start int
	end int
}

func main() {
	// intr := Interval{0, 3}
	// intr := Interval{start: 0, end: 3}
	intr := Interval{end: 3}
	fmt.Printf("%v", intr)
}
