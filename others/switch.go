package main

import "fmt"
import "time"

func calculate(n1 int32, n2 int32) (int32) {
  return n1 + n2
}

func main() {
	var i = 1

	switch i {
	case 1:
	case 2:
		fmt.Println("That is", i)
	}

	switch i {
	case 1:
		break
	case 2:
		fmt.Println("That is", i)
	}

	switch i {
	case 1:
		fallthrough
	case 2:
		fmt.Println("[fallthrough] That is", i)
	}

	num1 := 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

  res := calculate(1, 2) 
  switch {
    case res < 1: fmt.Println("Error 1") 
    case res < 2: fmt.Println("Error 2") 
    default: fmt.Println("Right")
  }

  t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
}
