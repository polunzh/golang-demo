package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now();
  fmt.Println(t.Year())
  fmt.Println(t.Month())
  fmt.Println(t.Day())
  fmt.Println(t.Minute())
  fmt.Println(t.Second())
  week := 60 * 60 * 24 * 7 * 1e9
  fmt.Println(t.Add(time.Duration(week)))
}
