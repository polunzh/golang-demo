package main

import (
  "fmt"
  "math"
  "errors"
)

func main() {
  fmt.Println(MySqrt(-12))
  fmt.Println(MySqrt(12))
  fmt.Println(MySqrt2(12))
}

func MySqrt(val float64) (float64, error) {
  if val < 0 {
    return float64(math.NaN()), errors.New("Invalid value")
  }

  return math.Sqrt(val), nil
}

func MySqrt2(val float64) (res float64, err error) {
  if val < 0 {
    res = float64(math.NaN())
    err = errors.New("Invalid value")
    return
  }

  res = math.Sqrt(val)
  return
}
