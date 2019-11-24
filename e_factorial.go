package main

import "fmt"

func main() {
  for i := 1; i < 10; i++ {
    fmt.Println(i, factorial(i))
  }
}

func factorial(n int) int {
  if n <= 1 {
    return 1 
  }

  return n * factorial(n - 1)
}
