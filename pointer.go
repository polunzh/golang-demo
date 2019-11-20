package main

import "fmt"

func main() {
  var i1 = 5
  var intP *int = &i1
  fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
  fmt.Println("intP:", *intP)

  s := "zhangzhenqiang"
  var p *string = &s
  *p = "zhangwanxia"
  fmt.Println("s:", s)
  fmt.Println("p:", p)
  fmt.Println("&p:", *p)

  var intPtr *int = nil
  *intPtr = 1 // runtime error
}
