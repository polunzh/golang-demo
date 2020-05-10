package main

import "fmt"

func main() {
  var name string = "zhangzhenqiang"
  fmt.Printf("string \"%s\" length: %d\n", name, len(name))
  fmt.Printf("\"%s\" at 1st: %c\n", name, name[1 - 0])
  fmt.Printf("\"%s\" at 6th: %c\n", name, name[6 - 1])

  var fullName = "zhang " + "zhenqiang"
  fmt.Println("Full Name:", fullName)
}
