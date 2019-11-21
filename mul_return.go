package main

import (
  "fmt"
  "strconv"
  "math"
)

func mySqrt(f float64) (v float64, ok bool) {
  if f < 0 { return }
  return math.Sqrt(f), true
}

func main() {
  var orig string = "ABC"
  // var newS string

  fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
  an, err := strconv.Atoi(orig)
  if err != nil {
    fmt.Printf("orig %s is not an integer\n", orig) 
    fmt.Println(err)
  }

  fmt.Print(an)

  t, ok := mySqrt(-1)
  fmt.Println(t, ok)
}
