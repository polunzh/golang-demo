package main

import "fmt"

func main() {
  s := "张镇强"
  fmt.Println(s)
  fmt.Println(reverseStr(s))

  s = "zhang"
  fmt.Println(s)
  fmt.Println(reverseStr(s))
}

func reverseStr(str string) string {
  s := []rune(str)
  n, h := len(s), len(s) / 2

  for i := 0; i < h; i++ {
    s[i], s[n - i - 1] = s[n - i - 1], s[i]
  }

  return string(s)
}
