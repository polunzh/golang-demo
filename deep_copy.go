package main

import (
  "fmt"
)

type Robot struct {
  Name string
  Color string
  Model string
}

func main() {
  fmt.Println("深拷贝")
  robot1 := Robot{
    Name: "zhangzhenqiang" ,
    Color: "red",
    Model: "people",
  }

  robot2 := robot1
  fmt.Printf("robot1: %s, address: %p\n", robot1, &robot1)
  fmt.Printf("robot2: %s, address: %p\n", robot2, &robot2)
}
