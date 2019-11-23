package main

import (
  "fmt"
)

func main() {
  var i = 5;
  for i > 0 {
    fmt.Println(i)
    i--
  }


  // for-range
  fmt.Println("for range")
  var str = "zhangzhenqiang"
  for idx, val := range str {
    fmt.Printf("idx: %d, char: %c\n", idx, val) 
  }
  str2 := "This is 中国话"
  for idx, val := range str2 {
    // fmt.Printf("idx: %d, char: %c\n", idx, val) 
    fmt.Printf("%-2d      %d      %U '%c' % X\n", idx, val, val, val, []byte(string(val)))
  }
}

