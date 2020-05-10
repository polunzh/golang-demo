package main

import (
  "fmt"
  "time"
)

const LIM = 41

func main() {
  start := time.Now()
  for i := 0; i < LIM; i++ {
    fmt.Println(i, fibonacci(i) )
  }

	fmt.Printf("time costs: %d\n", time.Now().Sub(start))
}

func fibonacci(n int) int {
  if n <= 1 {
     return 1 
  }

  return fibonacci(n - 2) + fibonacci(n - 1)
}
