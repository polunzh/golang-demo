package main

import (
  "fmt"
  "runtime"
  "os"
)

func main() {
  var a int;
  var c *int;
  var d = a;
  var e *int = c;
  fmt.Println(a, c)

  fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
  var goos string = runtime.GOOS
  fmt.Println("The operating system is:", goos)


  fmt.Println(&a, &d)
  fmt.Println(&c, &e)
  fmt.Println("log:", 23)

  name := "zhang"
  fmt.Println(name)

  log1 := 12
  log2 := 13
  fmt.Println("log1 =", log1, "log2 =", log2)
  log2, log1 = log1, log2
  fmt.Println("log1 =", log1, "log2 =", log2)

  path := os.Getenv("PATH")
  fmt.Println("Path is", path)
}
