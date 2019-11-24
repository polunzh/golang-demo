package main

import "fmt"

func main() {
  Varargs("This", "is", "zhangzhenqiang")
}

func Varargs(strs...string) {
  Vargargs2(strs...)
  for _, str := range strs {
    fmt.Print(str, " ") 
  }
}

func Vargargs2(strs...string) {
  fmt.Println(strs)
}
