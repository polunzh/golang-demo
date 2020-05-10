package main

import "fmt"

func main() {
  fmt.Println("main:", function1())

  fmt.Println("order...")
  order()
}

func function1() int {
  var i = 0
  defer fmt.Println(i)
  i++
  return i
}

func order() {
  for i := 0; i < 5; i++ {
    defer fmt.Println(i) 
  }
}
