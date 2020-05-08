package main

import "fmt"

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d: %c\n", i, c)
	}

	s = "zhangzhenqiang"
	fmt.Println(s[0:len(s)])
	c := []byte(s)
	c[0] = 'w'
	s = string(c)
	fmt.Println(s)
}
