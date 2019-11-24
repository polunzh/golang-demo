package main

import "fmt"

func main() {
  c, d := Multi(1, 2, 3)
  c1, d1 := Multi2(1, 2, 3)
  fmt.Println(c, d)
  fmt.Println(c1, d1)

  var res int = 0;
  Multiply(3, 2, &res)
  fmt.Println(res)
}

func Multi(a int, b int, c int) (int, int) {
  return a + b, c
}

func Multi2(a int, b int, c int) (x1 int, x2 int) {
  x1 = a + b
  x2 = c

  return
}

func Multiply(a, b int, reply *int) {
  *reply = a + b
}
