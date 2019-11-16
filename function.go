package main

import "fmt"

func Sum(a, b int) int {
  return a + b
}

func main() {
  var res = Sum(1, 2)
  fmt.Println(res);
}
