package main

import "fmt"

func main() {
  arr := []byte{'z', 'z', 'a', 'z', 'z', 'a'}
  res :=  make([]byte, len(arr))

  tmp := byte(0)
  idx := 0
  for _, value := range(arr) {
    if(tmp != value) {
      res[idx] = value
      idx += 1
    }

    tmp = value
  }

  fmt.Printf("%s", res)
}
