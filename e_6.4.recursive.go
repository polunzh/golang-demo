package main

import "fmt"

func main() {
  result := 0
  idx := 0
  for i := 0; i <= 10; i++ {
    idx, result = fibonacci(i) 
    fmt.Printf("fibonacci(%d) is %d\n", idx, result)
  }
}

func fibonacci(n int) (idx int, res int) {
  if n == 0 {
    res = 1
    idx = 0
  } else if n == 1 {
    res = 1
    idx = 1
  } else {
    _, t2 := fibonacci(n - 2)
    t3, t4 := fibonacci(n - 1)
    idx = t3 + 1
    res = t2 + t4
  }

  return
}
