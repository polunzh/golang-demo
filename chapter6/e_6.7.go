package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
 s := "Hello, 世界"
 fmt.Println(strings.Map(IsAscii, s))
}

func IsAscii(c rune) rune {
	if c > unicode.MaxASCII {
		return '?'
	}

	return c
}