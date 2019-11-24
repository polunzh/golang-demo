package main

import "fmt"

func main() {
  for i := 0; i < 10; i++ {
    fmt.Println(i, fibonacc(i) )
  }
}

func fibonacc(n int) int {
  if n <= 1 {
     return 1 
  }

  return fibonacc(n - 2) + fibonacc(n - 1)
}
