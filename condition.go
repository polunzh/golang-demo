package main

import (
  "fmt"
  "runtime"
)

func main() {
  if runtime.GOOS == "windows" {
    fmt.Printf("%s to quit, os: %s", "Ctrl+Z, Enter", runtime.GOOS) 
  } else {
    fmt.Printf("%s to quit, os: %s\n", "Ctrl + C", runtime.GOOS) 
  }

  if name := "dog"; name == "cat" {
    fmt.Println("This is cat") 
  } else {
    fmt.Printf("This is %s\n", name) 
  }
}
