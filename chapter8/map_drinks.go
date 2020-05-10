// map_drinks.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{
		"beer":   "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé"}

	names := make([]string, len(drinks))
	idx := 0
	for k := range drinks {
		names[idx] = k
		idx++
	}

	sort.Strings(names)

	for _, v := range names {
		fmt.Printf("%s: %s\n", v, drinks[v])
	}
}
