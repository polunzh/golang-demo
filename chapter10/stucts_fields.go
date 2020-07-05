package main

import "fmt"

type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main() {
	var ms *struct1 = new(struct1)
	ms.i1 = 10
	ms.f1 = 10.2
	ms.str = "zhangzhenqiang"

	fmt.Printf("the int is %d\n", ms.i1)
	fmt.Printf("the float is %.2f\n", ms.f1)
	fmt.Printf("the string is %s\n", ms.str)
	fmt.Printf("the string is %v\n", ms)
}
