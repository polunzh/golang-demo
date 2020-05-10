package main

import (
	"fmt"
	"strings"
)

func main() {
	var Days = map[int]string{1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
		7: "sunday"}

	var hasTuesday bool = false
	var hasHollyday bool = false
	for _, value := range Days {
		if value == strings.ToLower("Tuesday") {
			hasTuesday = true
		}

		if value == strings.ToLower("Hollyday") {
			hasHollyday = true
		}
	}

	fmt.Println(hasTuesday)
	fmt.Println(hasHollyday)
}
