package main

import "fmt"

func main(){
  const Ln2 = 0.693
  const Ln3 = 1/Ln2

  const hardEight = (1 << 100) >> 97

  const Monday, Tuesday = 1, 2
  const (
    Wednesday = 3
    Thursday = 4
  )


  fmt.Println(Ln2)
  fmt.Println(Ln3)
  fmt.Println(hardEight)
  fmt.Println(Monday, Tuesday, Wednesday, Thursday)

  const (
    Unknown = 0
    Male = 1
    Female = 4
    D = iota
  )

  fmt.Println(Unknown, Male, Female, D)
}
