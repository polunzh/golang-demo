package main

import "fmt"

func main() {
  fmt.Println(mul1(1, 2))
  a, b, c := mul2(1, 2)
  fmt.Println(a, b, c)
}

func mul1(x int, y int) (int, int, int) {
  return x + y, x * y, x - y
}

func mul2(x int, y int) (a int, b int, c int) {
  a, b, c = x + y, x * y, x - y
  return
}
