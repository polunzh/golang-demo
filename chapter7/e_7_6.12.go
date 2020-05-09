package main

import "fmt"

func main() {
	str := "zhangzhenqiang"
	fmt.Println(splitStr(str, 5))
}

func splitStr(s string, i int) (string, string) {
	return string(s[0:i]), string(s[i:])
}
