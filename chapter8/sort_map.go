package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	keys := make([]string, len(barVal))
	var i int = 0
	for k, v := range barVal {
		fmt.Printf("%s: %d\n", k, v)
		keys[i] = k
		i++
	}

	fmt.Println("\nsorted:")
	sort.Strings(keys)
	for _, val := range keys {
		fmt.Printf("%s: %d\n", val, barVal[val])
	}
}
