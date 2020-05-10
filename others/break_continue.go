package main

import "fmt"

func main() {
  var i = 0
  for {
    if i > 10 {
      break 
    }

    fmt.Println(i)
    i++
  }

  fmt.Println("continue");
  for i := 0; i < 10; i++ {
    if i % 2 == 0 {
      continue
    }

    fmt.Println(i) 
  }
}
