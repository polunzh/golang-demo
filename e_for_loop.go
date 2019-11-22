package main

import "fmt"

func main() {
  fmt.Println("---for---")
  for i := 0; i < 5; i++ {
    fmt.Println(i) 
  }

  fmt.Println("---goto---")
  t := 0
  LOOP: fmt.Println(t)
  t += 1
  if t < 5 {
    goto LOOP 
  }
}
