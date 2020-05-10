package main

import (
  "fmt"
  "math"
)

func main() {
  var n int16 = 32
  var m int32

  m = int32(n)
  fmt.Printf("32 bit is %d\n", m)
  fmt.Printf("16 bit is %d\n", n)

  fmt.Printf("min int32: %d", math.MinInt32)
  whole, fraction := math.Modf(10.100)
  fmt.Println(whole)
  fmt.Println(fraction)

  fmt.Println(1 ^ 1)
  fmt.Println(1 ^ 0)
  fmt.Println(0 ^ 1)
  fmt.Println(0 ^ 0)

  /*
  a, b := 10, 0
  c := a / b
  fmt.Println(c)
  */

  type Rope string
  var name Rope = "hello world"
  fmt.Println(name)
}
