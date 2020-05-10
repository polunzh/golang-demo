package main

import (
  "fmt"
)

func main() {
  fmt.Println();
  for row := 0; row < 10; row++ {
    for col := 0; col < 20; col++ {
      fmt.Print("*")
    }
    fmt.Println();
  }
}
