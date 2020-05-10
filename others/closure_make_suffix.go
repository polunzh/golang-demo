package main

import (
	"fmt"
	"strings"
)

func main() {
	addSuffix := MakeAddSuffix(".jpg")
	fmt.Println(addSuffix("zhang.jpg"))
	fmt.Println(addSuffix("zhang"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func (name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}

		return  name + suffix
	}
}