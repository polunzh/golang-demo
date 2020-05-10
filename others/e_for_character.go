package main

import "fmt"
import "strings"

func main() {
  for i := 1; i < 7; i++ {
    fmt.Println(strings.Repeat("G", i))
  }
}
