package main

import "fmt"

func main() {
  Greeting("zhang")
  Greeting("zhang", "zhen")
  Greeting("zhang", "bend", "dalin")
}

func Greeting(prefix string, who...string) {
  fmt.Println(prefix, who)
}
